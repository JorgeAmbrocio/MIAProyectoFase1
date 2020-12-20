package arbol

import (
	"fmt"
	"strconv"
	"strings"
)

// mkdisk estructura de la instrucción mkdisk
type mkgrp struct {
	name string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *mkgrp) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "name":
			i.name = QuitarComillas(p.Valor)
			break
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i mkgrp) Validar() bool {
	retorno := true

	if i.name == "" {
		retorno = false
	}

	return retorno
}

func Emkgrp(p []Parametro) {
	i := mkgrp{}
	i.MatchParametros(p)
	if i.Validar() {
		i.crearGrupo()
	}
}

func (i mkgrp) crearGrupo() {
	fmt.Println("Creando grupo")
	if UsuarioActualLogueado.GUID == 1 && UsuarioActualLogueado.UID == 1 {
		// sì es el usuario root
		// crear el grupo de usuarios

		// validar si el grupo ya existe
		_, mapaGrupo := getUsuarioYGrupo(*UsuarioActualLogueado.particion)
		if mapaGrupo[i.name] == 0 {
			// grupo no existe, se debe crear
			// obtener el contenido del archivo usuarios
			contenido := getContenidoArchivoUsuarios(*UsuarioActualLogueado.particion)
			// encontrar el id màs grande
			nuevoId := 0
			for _, maxId := range mapaGrupo {
				if nuevoId <= maxId {
					nuevoId = maxId + 1
				}
			}

			// añadir el nuevo contenido
			contenido += strconv.Itoa(nuevoId) + ",G," + i.name + "\n"
			escribirContenidoArchivo(contenido, 1, UsuarioActualLogueado.particion)
			fmt.Println("\tGrupo creado con èxito.")
		} else {
			fmt.Println("\tEl grupo \"" + i.name + "\" ya existe.")
		}

	} else {
		fmt.Println("\tÈsta acciòn requiere de permisos nivel administrador.")
	}
}
