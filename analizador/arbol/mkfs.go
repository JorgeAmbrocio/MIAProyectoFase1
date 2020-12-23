package arbol

import (
	"fmt"
	"strings"
)

type mkfs struct {
	Id         string
	Type       string
	EsRecovery bool
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct fdisk
func (i *mkfs) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "id":
			i.Id = p.Valor
			break
		case "type":
			i.Type = p.Valor
			break
		}
	}
}

func (i *mkfs) Validar() bool {
	retorno := true

	if i.Id == "" {
		retorno = false
	}
	return retorno
}
func Emkfs(p []Parametro) {
	fmt.Println("Ejecutando mkfs")
	i := mkfs{}
	i.MatchParametros(p)
	if i.Validar() {
		i.CrearSistemaDeArchivos()
		//i.CrarArchivoUsuarios()
	}
}

func (i *mkfs) CrearSistemaDeArchivos() {
	// recuperar la particiòn montada
	if exito, particionMontada := RecuperarParticionMontada(i.Id); exito {
		// sì se encontrò la particiòn
		// ejecutar el formato
		if !i.EsRecovery {
			if i.Type == "full" {
				WriteCeros(particionMontada.path, particionMontada.Start, particionMontada.Start+particionMontada.Size)
			}
			// escribir el super boque
			particionMontada.sp = crearSuperBloque(*particionMontada)
			particionMontada.sp.FirstBlock = 2
			particionMontada.sp.FirstInode = 2
			particionMontada.sp.FreeInodesCount -= 2
			particionMontada.sp.FreeBlocksCount -= 2
			escribirSuperBloque(particionMontada.path, particionMontada.sp, particionMontada.Start)

		}

		// crear primer inodo
		inodo := Inodo{
			UID:   1,
			GID:   1,
			Type:  0,
			Size:  0,
			Ctime: getFechaByte(),
			Atmie: getFechaByte(),
			Mtime: getFechaByte(),
			Perm:  [3]int8{7, 7, 5},
		}
		inodo.iniciarPunteros()
		inodo.Block[0] = 0

		//crear primer bloque para primer inodo
		bloqueCarpeta := BloqueCarpeta{}
		bloqueCarpeta.iniciarPunteros()
		bloqueCarpeta.indicarPadreyActual(0, 0)

		// crear el archivo
		var auxNombre [12]byte
		copy(auxNombre[:], "users.txt")
		bloqueCarpeta.Content[2].Name = auxNombre
		bloqueCarpeta.Content[2].PointerInode = 1

		// inodo del archivo
		inodoArchivo := Inodo{
			UID:   1,
			GID:   1,
			Type:  1,
			Size:  512,
			Ctime: getFechaByte(),
			Atmie: getFechaByte(),
			Mtime: getFechaByte(),
			Perm:  [3]int8{7, 7, 5},
		}
		inodoArchivo.iniciarPunteros()
		inodoArchivo.Block[0] = 1 // indica que dirige al bloque de archivos 1

		// bloque de archivo 1
		var auxCont [64]byte
		copy(auxCont[:], "1,G,root\n1,U,root,root,123\n")
		bloqueArchivo := BloqueArchivo{
			Content: auxCont,
		}

		//iniciar bitmaps
		bitmapInodos := []byte{1, 1}
		for i := 1; i < int(particionMontada.sp.InodesCount-2); i++ {
			bitmapInodos = append(bitmapInodos, 0)
		}
		bitmapBloques := []byte{1, 1}
		for i := 1; i < int(particionMontada.sp.BlocksCount-2); i++ {
			bitmapBloques = append(bitmapBloques, 0)
		}

		// escribir el super bloque
		escribirSuperBloque(particionMontada.path, particionMontada.sp, particionMontada.Start)

		//escribir los bitmap
		escribirBitMap(particionMontada.path, bitmapInodos, particionMontada.sp.BitMapInodeStart)
		escribirBitMap(particionMontada.path, bitmapBloques, particionMontada.sp.BitMapBlockStart)

		// escribir los bloques
		escribirInodo(particionMontada.path, inodo, particionMontada.sp.InodeStart+0*int64(particionMontada.sp.InodeSize))
		escribirInodo(particionMontada.path, inodoArchivo, particionMontada.sp.InodeStart+1*int64(particionMontada.sp.InodeSize))

		escribirBloqueCarpeta(particionMontada.path, bloqueCarpeta, particionMontada.sp.BlockStart+0*int64(particionMontada.sp.BlockSize))
		escribirBloqueArchivo(particionMontada.path, bloqueArchivo, particionMontada.sp.BlockStart+1*int64(particionMontada.sp.BlockSize))

		fmt.Println("Formato inicial del disco con èxito")
	} else {
		fmt.Println("No se ha encontrado la particiòn " + i.Id)
	}
}
