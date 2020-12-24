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
			if strings.Contains(i.path, ".txt") {
				//i.path += ".txt"
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
		var indiceInodo = int64(0)
		// recorrer todas las carpetas
		pathSplit := strings.Split(i.path, "/")
		for indice, carpeta := range pathSplit {
			if indice != 0 && indice != (len(pathSplit)-1) {
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
					if i.p == "p" {
						if !tienePermiso(inodo, 'w') {
							fmt.Println("\tNo tienes permisos de escritura en el directorio " + carpeta)
							return
						}
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

			if indice != (len(pathSplit)) {
				// recuperar carpeta del inodo
				_, _, iInodoCarpetaSiguiente := getCarpetaFromInodo(carpeta, inodo, *UsuarioActualLogueado.particion)
				if iInodoCarpetaSiguiente != -1 {
					// la carpeta existe
					//_, inodo2: = recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart+int64(UsuarioActualLogueado.particion.sp.InodeSize)*int64(iInodoCarpetaSiguiente))

					fmt.Println("El archivo ya existe, serà reemplazado")
					delete := rem{}
					delete.path = i.path
					delete.eliminar()
				}
			}
		}
		// llegamos al item del archivo
		// crear archivo
		crearArchivoEnInodo(int(indiceInodo), inodo, UsuarioActualLogueado.particion, i.size, pathSplit[len(pathSplit)-1])
		// guardar el journal y terminar
		guardarJournal(3, int32(i.size), i.path, i.p, [3]int8{}, UsuarioActualLogueado.particion)
		fmt.Println("\tArchivo creado con èxito ")
	} else {
		fmt.Println("\tNecesitas estar logueado para crear archivos")
	}
}
