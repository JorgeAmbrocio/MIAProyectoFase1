package arbol

import (
	"fmt"
	"strings"
)

type rem struct {
	path string
	p    string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *rem) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "path":
			i.path = QuitarComillas(p.Valor)
			if i.path[len(i.path)-1] == '/' {
				i.path = i.path[:len(i.path)-1]
			}
			break
		case "p":
			i.p = strings.ToLower(p.Valor)
			break
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i rem) Validar() bool {
	retorno := true
	if i.path == "" {
		retorno = false
	}

	if i.path == "/users.txt" {
		retorno = false
		fmt.Println("No puedes eliminar el archivo de usuarios")
	}
	return retorno
}

func Erem(p []Parametro) {
	i := rem{}
	i.MatchParametros(p)
	if i.Validar() {
		i.eliminar()
	}
}

func (i *rem) eliminar() {
	fmt.Println("Eliminando")
	if UsuarioActualLogueado.UID != 0 {
		// recuperar el primer inodo
		_, inodo := recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart)
		var apuntadorCarpetaContenedora int32 = 0
		var indiceEnCarpeta int = 0
		// recorrer todas las carpetas
		pathSplit := strings.Split(i.path, "/")
		for indice, carpeta := range pathSplit {
			if indice != 0 && indice != (len(pathSplit)) {
				// recuperar carpeta del inodo
				apuntadorCarpeta, indiceApuntadorCarpeta, iInodoCarpetaSiguiente := getCarpetaFromInodo(carpeta, inodo, *UsuarioActualLogueado.particion)
				apuntadorCarpetaContenedora = apuntadorCarpeta
				indiceEnCarpeta = indiceApuntadorCarpeta
				if iInodoCarpetaSiguiente != -1 {
					// la carpeta existe
					_, inodo = recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart+int64(UsuarioActualLogueado.particion.sp.InodeSize)*int64(iInodoCarpetaSiguiente))
					if !tienePermiso(inodo, 'r') {
						fmt.Println("\tNo tienes permisos para eliminar èste archivo")
						return
					}
				} else {
					// la carpeta no existe
					// mostrar error porque la ruta no es vàlida
					fmt.Println("\tNo se ha encontrado el directorio " + carpeta)
					return
				}
			}
		}
		// llegamos al item del archivo
		// eliminar los inodos
		if !tienePermiso(inodo, 'w') {
			fmt.Println("\tNo tienes permisos para eliminar èste archivo")
			return
		}
		_, bloqueCarpeta := recuperarBloqueCarpeta(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(apuntadorCarpetaContenedora))
		remove(apuntadorCarpetaContenedora, bloqueCarpeta, int32(indiceEnCarpeta), UsuarioActualLogueado.particion)
		// guardar el journal y terminar
		guardarJournal(5, 0, i.path, "", [3]int8{}, UsuarioActualLogueado.particion)
		fmt.Println("\tSe ha removido con èxito")
	} else {
		fmt.Println("Necesitas estar logueado para poder utilizar èste comando.")
	}
}
