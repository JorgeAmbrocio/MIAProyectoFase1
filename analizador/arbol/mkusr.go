package arbol

import (
	"fmt"
	"strconv"
	"strings"
)

// mkdisk estructura de la instrucción mkdisk
type mkusr struct {
	usr   string
	pwd   string
	grupo string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *mkusr) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "usr":
			i.usr = QuitarComillas(p.Valor)
			break
		case "pwd":
			i.pwd = (p.Valor)
			break
		case "grp":
			i.grupo = QuitarComillas(p.Valor)
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i mkusr) Validar() bool {
	retorno := true

	if i.usr == "" || i.pwd == "" || i.grupo == "" {
		retorno = false
	}

	return retorno
}

func Emkusr(p []Parametro) {
	i := mkusr{}
	i.MatchParametros(p)
	if i.Validar() {
		i.crearUsuario()
	}
}

func (i *mkusr) crearUsuario() {
	fmt.Println("Creando usuario")
	if UsuarioActualLogueado.GUID == 1 && UsuarioActualLogueado.UID == 1 {
		// sì es el usuario root
		// crear el grupo de usuarios

		// validar si el grupo ya existe
		mapaUsuario, mapaGrupo := getUsuarioYGrupo(*UsuarioActualLogueado.particion)
		if mapaUsuario[i.usr].UID == 0 {
			// verificar si el grupo existe
			if mapaGrupo[i.grupo] != 0 {
				// usuario no existe, se debe crear
				// obtener el contenido del archivo usuarios
				contenido := getContenidoArchivoUsuarios(*UsuarioActualLogueado.particion)
				// encontrar el id màs grande
				nuevoId := 0
				for _, maxId := range mapaUsuario {
					if int32(nuevoId) <= maxId.UID {
						nuevoId = int(maxId.UID) + 1
					}
				}
				// añadir el nuevo contenido
				contenido += strconv.Itoa(nuevoId) + ",U," + i.grupo + "," + i.usr + "," + i.pwd + "\n"
				escribirContenidoArchivo(contenido, 1, UsuarioActualLogueado.particion)

				// guardar el journal y terminar
				guardarJournal(2, 0, i.usr+","+i.pwd+","+i.grupo, "", [3]int8{}, UsuarioActualLogueado.particion)
				fmt.Println("\tUsuario creado con èxito.")
			} else {
				fmt.Println("El grupo no existe, no se podrà crear el usuario")
			}
		} else {
			fmt.Println("\tEl usuario \"" + i.usr + "\" ya existe.")
		}

	} else {
		fmt.Println("\tÈsta acciòn requiere de permisos nivel administrador.")
	}
}
