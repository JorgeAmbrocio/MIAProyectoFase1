package arbol

import (
	"fmt"
	"strings"
)

// mkdisk estructura de la instrucción mkdisk
type cat struct {
	paths []string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *cat) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		p.Tipo = p.Tipo[:len(p.Tipo)-1]
		switch strings.ToLower(p.Tipo) {
		case "file":
			i.paths = append(i.paths, QuitarComillas(p.Valor))
			break
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i cat) Validar() bool {
	retorno := true

	if len(i.paths) == 0 {
		retorno = false
	}

	return retorno
}

// Emkdisk es la ejecución del mkdisk
func Ecat(p []Parametro) {
	/*
		crear objeto
		incluir parámetros
		validar si cuenta con lo suficiente para ejecutarse
	*/
	i := cat{}
	i.MatchParametros(p)
	if i.Validar() {
		fmt.Println("CAT")
		if UsuarioActualLogueado.UID == 0 {
			fmt.Println("\tDebes estar logueado para crear ejecutar CAT")
			return
		}
		var contenidoFinal string = ""
		for _, path := range i.paths {
			// recorrer todas las rutas
			if strings.Contains(path, ".txt") {
				// verifica que sì es un archivo
				inodo, pointerInodo, _, _ := encontrarArchivo(path, *UsuarioActualLogueado.particion)

				if pointerInodo == -1 {
					continue // si no encontrò el inodo, ir a la siguiente iteracciòn
				}

				contenidoArchivo := getContenidoArchivo(inodo, *UsuarioActualLogueado.particion)
				contenidoFinal += "Archivo " + path + "\n\t"
				contenidoFinal += contenidoArchivo + "\n\n"
			}
		}
		fmt.Println(contenidoFinal)
		fmt.Println("\tTermiando con èxito")
	}
}
