package arbol

import (
	"fmt"
	"strconv"
)

type fdisk struct {
	size          int
	path          string
	name          string
	unit          string
	tipo          string
	fit           string
	delete        string
	add           int
	multiplicador int
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct fdisk
func (i *fdisk) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch p.Tipo {
		case "size":
			if valor, err := strconv.Atoi(p.Valor); err != nil {
				fmt.Println("ERROR FDISK: size no es válido ", p.Valor)
			} else {
				i.size = valor
			}
			break
		case "path":
			i.path = QuitarComillas(p.Valor)
			break
		case "name":
			i.name = p.Valor
			break
		case "unit":
			i.unit = p.Valor
			break
		case "type":
			i.tipo = p.Valor
			break
		case "fit":
			i.fit = p.Valor
			break
		case "delete":
			i.delete = p.Valor
			break
		case "add":
			if valor, err := strconv.Atoi(p.Valor); err != nil {
				fmt.Println("ERROR FDISK: add no es válido ", p.Valor)
			} else {
				i.add = valor
			}
			break
		}
	}

	if i.unit == "" {
		i.unit = "m"
		i.multiplicador = 1024
	} else {
		i.multiplicador = 1
	}
}

func (i *fdisk) Validar() bool {
	retorno := true

	if i.path == "" || i.name == "" {
		retorno = false
	}

	return retorno
}

func Efdisk(p []Parametro) {
	fmt.Println("spy el fdisk y me estoy ejecutando")

	i := fdisk{}
	i.MatchParametros(p)
	if i.Validar() {
		// identificar què tipo de ejecuciòn se debe realizar

		// crear una particiòn
		if i.size != 0 && i.tipo != "" {
			// se està creando una particiòn
			if i.fit != "" {
				i.fit = "wf"
			}
			if i.unit != "" {
				i.unit = "k"
			}

			//fmt.Println("Se ha creado la particiòn")
			i.crearParticion()
			return
		}

		if i.add != 0 && false {
			// editando el tamaño de una particiòn
			if i.unit != "" {
				i.unit = "k"
			}
			return
		}

		if i.delete != "" && false {
			// eliminando una particiòn
			return
		}

	}

}

/*acciones de fdisk*/
func (i *fdisk) crearParticion() {
	// recuperar el mbr
	mbr := RecuperarMBR(i.path + i.name)

	fmt.Println(mbr)

}
