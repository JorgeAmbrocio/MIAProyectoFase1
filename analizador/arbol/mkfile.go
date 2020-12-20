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
			spliteado := strings.Split(i.path, "/")
			var directorio string = ""
			for _, str := range spliteado[1 : len(spliteado)-1] {
				directorio += "/" + str
			}
			CrearTodasCarpetas(directorio)
			break
		case "unit":
			i.p = strings.ToLower(p.Valor)
			break
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i mkfile) Validar() bool {
	retorno := true

	if i.size < 0 || i.path == "" {
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
		fmt.Println("vamo a crear el archivo")
	}
}
