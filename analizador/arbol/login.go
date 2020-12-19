package arbol

import (
	"fmt"
	"strconv"
	"strings"
)

type login struct {
	usr string
	pwd string
	id  string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *login) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "usr":
			i.usr = p.Valor
			break
		case "pwd":
			i.pwd = p.Valor
			break
		case "id":
			i.id = p.Valor
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i login) Validar() bool {
	retorno := true

	if i.usr == "" || i.pwd == "" || i.id == "" {
		retorno = false
	}

	return retorno
}

func Elogin(p []Parametro) {
	i := login{}
	i.MatchParametros(p)
	if i.Validar() {
		i.ingresar()
	}
}

func (i *login) ingresar() {
	// recuperar la particiòn montada
	if exito, particion := RecuperarParticionMontada(i.id); exito {
		// recupear inodo del archivo
		// i know que el inodo de usuarios.txt es el inodo "1"
		_, inodo := recuperarInodo(particion.path, particion.sp.InodeStart+1*int64(particion.sp.InodeSize))
		contenidoArchivo := getContenidoArchivo(inodo, *particion)

		mapaUsuario := make(map[string]UsuarioArchivo)
		mapaGrupo := make(map[string]int)

		filas := strings.Split(contenidoArchivo, "\n")
		for _, fila := range filas {
			atributos := strings.Split(fila, ",")
			if len(atributos) == 5 {
				// usuario

				if uid, err := strconv.Atoi(atributos[0]); err == nil {
					usr := UsuarioArchivo{UID: int32(uid), grupo: atributos[2], nombre: atributos[3], contrasena: atributos[4]}
					mapaUsuario[usr.nombre] = usr
				}
			} else {
				// grupo
				if gid, err := strconv.Atoi(atributos[0]); err == nil {
					mapaGrupo[atributos[2]] = gid
				}
			}
		}
		// obtener usuario
		if usr := mapaUsuario[i.usr]; usr.UID != 0 && usr.contrasena == i.pwd {
			// sì encontramos el usuario
			// validar el grupo
			if grp := mapaGrupo[usr.grupo]; grp != 0 {
				// sì existe el grupo
				// cargar el usuario
				UsuarioActualLogueado = Usuario{GUID: int32(grp), UID: int32(usr.UID)}
				fmt.Println("Ya se logueó joven")
			} else {
				fmt.Println("Tu grupo no existe")
			}
		} else {
			fmt.Println("Lo siento, bruh, no existes")
		}
	} else {
		fmt.Println("No se ha encontrado la particiòn con id " + i.id)
	}
}
