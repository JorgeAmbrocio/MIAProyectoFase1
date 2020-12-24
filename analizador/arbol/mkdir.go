package arbol

import (
	"fmt"
	"strings"
)

type mkdir struct {
	path string
	p    string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *mkdir) MatchParametros(lp []Parametro) {
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

	if i.p == "" {
		i.p = " "
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i mkdir) Validar() bool {
	retorno := true
	if i.path == "" {
		retorno = false
	}
	return retorno
}

func Emkdir(p []Parametro) {
	i := mkdir{}
	i.MatchParametros(p)
	if i.Validar() {
		i.crearCarpeta()
	}
}

func (i *mkdir) crearCarpeta() {
	// recuperar el primer inodo
	u := UsuarioActualLogueado
	if UsuarioActualLogueado.UID == 0 {
		fmt.Println("\tDebes estar logueado para crear directorios")
		return
	}
	_, inodo := recuperarInodo(u.particion.path, u.particion.sp.InodeStart)
	var indiceInodo = int64(0)
	// recorrer todas las carpetas
	pathSplit := strings.Split(i.path, "/")
	for indice, carpeta := range pathSplit {
		if indice != 0 && indice != (len(pathSplit)) {
			// recuperar carpeta del inodo
			_, _, iInodoCarpetaSiguiente := getCarpetaFromInodo(carpeta, inodo, *UsuarioActualLogueado.particion)
			if iInodoCarpetaSiguiente != -1 {
				// la carpeta existe
				_, inodo = recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart+int64(UsuarioActualLogueado.particion.sp.InodeSize)*int64(iInodoCarpetaSiguiente))
				indiceInodo = int64(iInodoCarpetaSiguiente)
				if !tienePermiso(inodo, 'r') {
					fmt.Println("\tNo tienes permisos de lectura en el directorio " + carpeta)
					return
				}
			} else {
				// la carpeta no existe
				if i.p == "p" || indice == len(pathSplit)-1 {
					// crear carpeta de manera forzada
					if !tienePermiso(inodo, 'w') {
						fmt.Println("\tNo tienes permisos de escritura en el directorio " + carpeta)
						return
					}
					indiceCarpetaNueva := crearCarpetaEnInodo(int64(indiceInodo), inodo, UsuarioActualLogueado.particion, carpeta)
					if indiceCarpetaNueva != -1 {
						_, inodo = recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart+int64(UsuarioActualLogueado.particion.sp.InodeSize)*int64(indiceCarpetaNueva))
						fmt.Println("\tCarpeta creada con èxito " + carpeta)
					}
				} else {
					// mostrar error porque la ruta no es vàlida
					fmt.Println("\tNo se ha encontrado el directorio " + carpeta)
					return
				}
			}
		}
	}

	// guardar el journal y terminar
	guardarJournal(4, int32(byte(i.p[0])), i.path, "", [3]int8{}, UsuarioActualLogueado.particion)

}
