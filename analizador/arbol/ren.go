package arbol

import (
	"fmt"
	"strings"
)

type ren struct {
	path string
	name string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *ren) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "path":
			i.path = QuitarComillas(p.Valor)
			if i.path[len(i.path)-1] == '/' {
				i.path = i.path[:len(i.path)-1]
			}
			break
		case "name":
			i.name = p.Valor
			break
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i ren) Validar() bool {
	retorno := true

	if i.path == "" || i.name == "" {
		retorno = false
	}

	return retorno
}

// Emkdisk es la ejecución del mkdisk
func Eren(p []Parametro) {
	/*
		crear objeto
		incluir parámetros
		validar si cuenta con lo suficiente para ejecutarse
	*/
	i := ren{}
	i.MatchParametros(p)
	if i.Validar() {
		i.rename()
	}
}

func (i *ren) rename() {
	fmt.Println("REN")
	if UsuarioActualLogueado.UID == 0 {
		fmt.Println("\tDebes estar logueado para renombrar")
		return
	}
	// verifica que sì es un archivo
	inodo, pointerInodo, pointerCarpeta, indiceCarpeta := encontrarArchivo(i.path, *UsuarioActualLogueado.particion)

	if pointerInodo == -1 {
		fmt.Println("\tNo se ha encontrado la ruta indicada")
		return
	} else if tienePermiso(inodo, 'w') {
		// recuperamos el bloque carpeta
		_, bloqueCarpeta := recuperarBloqueCarpeta(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(pointerCarpeta))
		// indicar el nuevo nombre
		var auxName [12]byte
		copy(auxName[:], i.name)
		bloqueCarpeta.Content[indiceCarpeta].Name = [12]byte{}
		bloqueCarpeta.Content[indiceCarpeta].Name = auxName
		// escribir nuevamente los bloques
		escribirBloqueCarpeta(UsuarioActualLogueado.particion.path, bloqueCarpeta, UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(pointerCarpeta))
		// guardar el journal y terminar
		guardarJournal(6, int32(0), i.path, i.name, [3]int8{}, UsuarioActualLogueado.particion)
		fmt.Println("\tTermiando con èxito")
	}
}
