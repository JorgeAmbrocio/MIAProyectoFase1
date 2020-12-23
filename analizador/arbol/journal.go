package arbol

import (
	"encoding/binary"
	"fmt"
	"strings"
)

/*
type Journaling struct {
	TipoOperacion int16 // 1-12
	Tipo          int32
	Nombre        [160]byte
	Content1      [160]byte
	//Content     []byte
	Fecha       [20]byte
	Propietario [2]byte // uid, gid
	Permisos    [3]int8 // propietario, grupos, otros
}
*/
type recovery struct {
	id string
}

var activoJournal bool = true

func (i *recovery) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "id":
			i.id = p.Valor
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i *recovery) Validar() bool {
	retorno := true

	if i.id == "" {
		retorno = false
	}

	return retorno
}

func Erecovery(p []Parametro) {
	i := recovery{}
	i.MatchParametros(p)
	if i.Validar() {
		i.recuperarSistema()
	}
}

func (i *recovery) recuperarSistema() {
	// guardar el usuario actual logueado
	fmt.Println("Iniciando la recuperaciòn del sistema")
	var auxUsuarioActivo = UsuarioActualLogueado
	activoJournal = false
	// recuperar la particiòn montada
	if exito, particion := RecuperarParticionMontada(i.id); exito {
		// recuperar el inicio del sistema de archivos
		fs := mkfs{}
		fs.Id = i.id
		fs.EsRecovery = true
		fs.CrearSistemaDeArchivos()

		// recorrer todos los journal
		for indice := int32(0); indice < particion.sp.JournalCount; indice++ {
			// recuperar el journal actual
			_, journal := recuperarJournal(particion.path, indice, *particion)

			// indicar permisos
			UsuarioActualLogueado = Usuario{
				UID:       journal.Propietario[0],
				GUID:      journal.Propietario[1],
				particion: particion,
			}

			// identificar el tipo de acciòn
			switch journal.TipoOperacion {
			case 1:
				accion := mkgrp{}
				accion.name = BytesToString(journal.Nombre[:])
				accion.crearGrupo()
				break
			case 2:
				// separar el nombre del usuario
				var spliteado = strings.Split(BytesToString(journal.Nombre[:]), ",")

				// crear la acciòn
				accion := mkusr{}
				accion.usr = spliteado[0]
				accion.pwd = spliteado[1]
				accion.grupo = spliteado[2]

				accion.crearUsuario()
				break
			case 3:
				// mkfile

				// crear la acciòn
				accion := mkfile{}
				accion.path = BytesToString(journal.Nombre[:])
				accion.size = int(journal.Tipo)

				accion.p = ByteToString(byte(journal.Tipo))
				accion.crearArchivo()
				break
			case 4:
				// mkdir
				accion := mkdir{}
				accion.path = BytesToString(journal.Nombre[:])
				accion.p = ByteToString(byte(journal.Tipo))
				accion.crearCarpeta()
				break
			case 5:
				// rem
				accion := rem{}
				accion.path = BytesToString(journal.Nombre[:])
				accion.eliminar()
				break
			case 6:
				// ren
				accion := ren{}
				accion.path = BytesToString(journal.Nombre[:])
				accion.name = BytesToString(journal.Content[:])
				accion.rename()
				break
			case 7:
				break
			case 8:
				break
			case 9:
				break
			case 10:
				break
			default:

			}
		}
	} else {
		fmt.Println("\tNo se ha encontrado la particiòn con id " + i.id)
	}

	UsuarioActualLogueado = auxUsuarioActivo
	activoJournal = true
}

func guardarJournal(
	TipoOperacion int16,
	tipo int32,
	Nombre string,
	Content1 string,
	//Fecha [20]byte,
	//Propietario [2]int32,
	Permisos [3]int8, particion *ParticionMontada) {

	// verifica si està activa la secciòn de escritura journal
	// cuardo se està ejecutando la recuperaciòn del sistema
	// no se debe seguir escribiendo bloques de journal
	if !activoJournal {
		return // evita escribir bloques journal cuando se realiza la recuperaciòn del sistema
	}

	// preparar variables auxiliares
	var auxNombre [160]byte
	copy(auxNombre[:], Nombre)

	var auxContenido [160]byte
	copy(auxContenido[:], Content1)

	journal := Journaling{
		TipoOperacion: TipoOperacion,
		Tipo:          tipo,
		Nombre:        auxNombre,
		Content:       auxContenido,
		Propietario:   [2]int32{UsuarioActualLogueado.UID, UsuarioActualLogueado.GUID},
		Permisos:      Permisos,
		Fecha:         getFechaByte(),
	}

	// escribir journal
	escribirJournal(particion.path,
		journal,
		particion.Start+int64(binary.Size(SuperBlock{}))+int64(binary.Size(Journaling{}))*int64(particion.sp.JournalCount))
	particion.sp.JournalCount++
}
