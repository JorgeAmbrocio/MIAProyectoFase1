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

	// verificar si existe la ruta vieja
	_, pointerInodo, pointerCarpeta, indiceCarpeta := encontrarArchivoConPermiso(i.path, *UsuarioActualLogueado.particion)
	if pointerInodo == -1 {
		fmt.Println("\tRuta actual no existe o se tiene permisos de lectura/escritura")
		return
	}

	// verificar si existe la ruta nueva
	inodo2, _, pointerCarpeta2, indiceCarpeta2 := encontrarArchivoConPermiso(i.dest, *UsuarioActualLogueado.particion)
	if pointerInodo == -1 {
		fmt.Println("\tRuta destino no existe o se tiene permisos de lectura/escritura")
		return
	}

	fmt.Println(inodo2, pointerInodo, pointerCarpeta, pointerCarpeta2, indiceCarpeta, indiceCarpeta2)
	// intercambiar los ìndices de los inodos
	// eliminar del puntero anterior
	//recuperar carpeta
	_, carpeta := recuperarBloqueCarpeta(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(pointerCarpeta))
	//_, carpeta2 := recuperarBloqueCarpeta(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(pointerCarpeta2))
	// crear en el nuevo puntero
	auxCarpeta := carpeta

	// eliminar punteros anteriores
	carpeta.Content[indiceCarpeta].Name = [12]byte{}
	carpeta.Content[indiceCarpeta].PointerInode = -1

	// buscar puntero vacìo para insertar el mov
	for indice, apuntador := range inodo2.Block {
		// recorrer todos los punteros del inodo
		switch {
		case indice < 13 && apuntador != -1:
			// recuperar el bloque de carpeta
			_, bloqueCarpeta := recuperarBloqueCarpeta(
				UsuarioActualLogueado.particion.path,
				UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(apuntador))
			// recorrer los apuntadores de la carpeta
			for indice2, apuntador2 := range bloqueCarpeta.Content {
				// buscar el aapuntador libre
				if apuntador2.PointerInode == -1 {
					// añadir en èsta parte el indice para la carpeta a mover
					bloqueCarpeta.Content[indice2].Name = auxCarpeta.Content[indiceCarpeta].Name
					bloqueCarpeta.Content[indice2].PointerInode = auxCarpeta.Content[indiceCarpeta].PointerInode

					// escribir el bloque carpeta nuevo
					escribirBloqueCarpeta(
						UsuarioActualLogueado.particion.path,
						bloqueCarpeta,
						UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(apuntador),
					)

					// escribir bloque carpeta antiguo
					escribirBloqueCarpeta(
						UsuarioActualLogueado.particion.path,
						carpeta,
						UsuarioActualLogueado.particion.sp.BlockStart+int64(UsuarioActualLogueado.particion.sp.BlockSize)*int64(pointerCarpeta),
					)
					// guardar el journal y terminar
					guardarJournal(8, 0, i.path, i.dest, [3]int8{}, UsuarioActualLogueado.particion)

					fmt.Println("\tCarpeta movida con èxito")
					return
				}
			}
			break
		case indice == 13 && apuntador != -1:
			break
		case indice == 14 && apuntador != -1:
			break
		}
	}

	// encontrarun puntero libre para añadir la nueva carpeta
	fmt.Println("\tError al mover la ruta, no se encontrò espacio para dicha acciòn")
}
