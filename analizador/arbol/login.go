package arbol

import (
	"fmt"
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
	fmt.Println("Iniciando login")
	// recuperar la particiòn montada
	if exito, particion := RecuperarParticionMontada(i.id); exito {
		// recupear inodo del archivo
		// i know que el inodo de usuarios.txt es el inodo "1"
		mapaUsuario, mapaGrupo := getUsuarioYGrupo(*particion)
		// obtener usuario
		if usr := mapaUsuario[i.usr]; usr.UID != 0 && usr.contrasena == i.pwd {
			// sì encontramos el usuario
			// validar el grupo
			if grp := mapaGrupo[usr.grupo]; grp != 0 {
				// sì existe el grupo
				// cargar el usuario
				UsuarioActualLogueado = Usuario{GUID: int32(grp), UID: int32(usr.UID), particion: particion}
				fmt.Println("\tYa se logueó joven")
			} else {
				fmt.Println("\tTu grupo no existe")
			}
		} else {
			fmt.Println("\tLo siento, bruh, no existes")
		}
	} else {
		fmt.Println("\tNo se ha encontrado la particiòn con id " + i.id)
	}
}
