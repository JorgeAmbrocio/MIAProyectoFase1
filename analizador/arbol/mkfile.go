package arbol

import (
	"fmt"
	"strconv"
	"strings"
)

type mkfile struct {
	path string
	p    string
	size int
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *mkfile) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "size":
			if valor, err := strconv.Atoi(p.Valor); err != nil {
				fmt.Println("ERROR MKDISK: size no es válido ", p.Valor)
			} else {
				i.size = valor
			}
			break
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
func (i mkfile) Validar() bool {
	retorno := true
	if i.size <= 0 || i.path == "" {
		retorno = false
	}
	return retorno
}

func Emkfile(p []Parametro) {
	/*
		crear objeto
		incluir parámetros
		validar si cuenta con lo suficiente para ejecutarse
	*/
	i := mkfile{}
	i.MatchParametros(p)
	if i.Validar() {
		//i.CrearBinario()
		i.crearArchivo()
	}
}

func (i *mkfile) crearArchivo() {
	fmt.Println("Creando archivo " + i.path)

	if UsuarioActualLogueado.UID != 0 {
		// recuperar el primer inodo
		_, inodo := recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart)
		var indiceInodo = 0
		// recorrer todas las carpetas
		pathSplit := strings.Split(i.path, "/")
		for indice, carpeta := range pathSplit {
			if indice != 0 && indice != (len(pathSplit)-1) {
				// recuperar carpeta del inodo
				_, _, iInodoCarpetaSiguiente := getCarpetaFromInodo(carpeta, inodo, *UsuarioActualLogueado.particion)
				if iInodoCarpetaSiguiente != -1 {
					// la carpeta existe
					_, inodo = recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart+int64(UsuarioActualLogueado.particion.sp.InodeSize)*int64(iInodoCarpetaSiguiente))
				} else {
					// la carpeta no existe
					if i.p == "p" {
						// crear carpeta de manera forzada
						indiceCarpetaNueva := crearCarpetaEnInodo(int64(indiceInodo), inodo, UsuarioActualLogueado.particion, carpeta)
						if indiceCarpetaNueva != -1 {
							_, inodo = recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart+int64(UsuarioActualLogueado.particion.sp.InodeSize)*int64(indiceCarpetaNueva))
						}
					} else {
						// mostrar error porque la ruta no es vàlida
						fmt.Println("\tNo se ha encontrado el directorio " + carpeta)
						return
					}
				}
			}
		}
		// llegamos al item del archivo
		// crear archivo
		crearArchivoEnInodo(indiceInodo, inodo, UsuarioActualLogueado.particion, i.size, pathSplit[len(pathSplit)-1])
	} else {
		fmt.Println("\tNecesitas estar logueado para crear archivos")
	}
}
