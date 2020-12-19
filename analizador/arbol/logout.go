package arbol

import (
	"fmt"
)

func Elogout(p []Parametro) {
	// desloguear
	if UsuarioActualLogueado.UID != 0 {
		// sì està logueado
		UsuarioActualLogueado = Usuario{}
		fmt.Println("Se ha desloguado con èxito")
	} else {
		// no estaba logueado
		fmt.Println("No estaba logueado")
	}
}
