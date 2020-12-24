package arbol

import (
	"fmt"
	"strings"
)

type mv struct {
	path string
	dest string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *mv) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "path":
			i.path = QuitarComillas(p.Valor)
			if i.path[len(i.path)-1] == '/' {
				i.path = i.path[:len(i.path)-1]
			}
			break
		case "dest":
			i.dest = QuitarComillas(p.Valor)
			if i.path[len(i.path)-1] == '/' {
				i.dest = i.path[:len(i.path)-1]
			}
			break
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i mv) Validar() bool {
	retorno := true
	if i.path == "" || i.dest == "" {
		retorno = false
	}
	return retorno
}

func Emv(p []Parametro) {
	i := mv{}
	i.MatchParametros(p)
	if i.Validar() {
		i.mover()
	}
}

func (i *mv) mover() {
	fmt.Println("MV")

	if UsuarioActualLogueado.GUID == 0 {
		fmt.Println("\tDebes estar logueado para ejecutar èste comando")
		return
	}
	/*
		// verificar si existe la ruta vieja
		inodo, pointerInodo, pointerCarpeta, indiceCarpeta := encontrarArchivoConPermiso(i.path, *UsuarioActualLogueado.particion)
		if pointerInodo == -1 {
			fmt.Println("\tRuta actual no existe o se tiene permisos de lectura/escritura")
			return
		}

		// verificar si existe la ruta nueva
		inodo2, pointerInodo2, pointerCarpeta2, indiceCarpeta2 := encontrarArchivoConPermiso(i.dest, *UsuarioActualLogueado.particion)
		if pointerInodo == -1 {
			fmt.Println("\tRuta destino no existe o se tiene permisos de lectura/escritura")
			return
		}

		fmt.Println(inodo, inodo2, pointerInodo, pointerInodo2, pointerCarpeta, pointerCarpeta2, indiceCarpeta, indiceCarpeta2)
		// intercambiar los ìndices de los inodos
		// eliminar del puntero anterior
		//recuperar carpeta
		_, carpeta := recuperarBloqueCarpeta(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(indiceCarpeta))
		_, carpeta2 := recuperarBloqueCarpeta(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(indiceCarpeta2))
		// crear en el nuevo puntero
		auxCarpeta := carpeta

		// encontrarun puntero libre para añadir la nueva carpeta
	*/

}
