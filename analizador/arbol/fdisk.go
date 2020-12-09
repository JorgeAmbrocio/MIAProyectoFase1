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
	mbr           Mbr
	sizeBytes     int
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

	if i.fit == "" {
		i.fit = "wf"
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

			// calcular el tamaño en bytes de la nueva particiòn
			if i.unit == "m" {
				i.multiplicador = 1024 * 1024
			} else if i.unit == "k" {
				i.multiplicador = 1024
			}

			i.sizeBytes = i.size * i.multiplicador

			// ejecutar mètodo para crear la particiòn
			if i.tipo == "p" || i.tipo == "e" {
				i.crearParticionPE()
			} else {
				i.crearParticionL()
			}

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
func (i *fdisk) crearParticionPE() {
	// recuperar el mbr
	i.mbr = RecuperarMBR(i.path + i.name)
	auxMbr := i.mbr
	fmt.Println(i.mbr)

	// si la particiòn es de tipo Primaria o extendida
	// validar si mbr tiene particiones libres
	var idParticionLibre int = -1
	for i := 0; i <= 3; i++ {
		auxParticion := auxMbr.Partitions[i]
		if auxParticion.Status == 0 || auxParticion.Status == 2 {
			idParticionLibre = i
			break
		}
	}

	for index := 0; index <= 3 && i.tipo == "e"; index++ {
		auxParticion := auxMbr.Partitions[index]
		if auxParticion.Status == 1 && auxParticion.Type == 'e' {
			fmt.Println("Ya existe una particiòn extendidah.")
			return
		}
	}

	// ya tengo el id de la particiòn libre que puedo utilizar para crear
	if idParticionLibre != -1 {
		// sì hay una particiòn libre
		auxParticion := auxMbr.Partitions[idParticionLibre]
		fmt.Println(auxParticion)

		// buscar espacios libres
		var espaciosVacios = getEspaciosLibres(auxMbr)

		// revisar en espacios dependiendo del tipo de fit
		var inicioCorrecto int = -1
		var tamanoCorrecto int = 0
		for indice, objeto := range espaciosVacios.Inicios {
			tamano := espaciosVacios.Finales[indice] - objeto
			if i.sizeBytes <= tamano {
				// la particiòn nueva sì cabe en el espacio disponible
				if auxMbr.Fit == 'f' {
					// primer ajuste
					// guarda el inicio de èste espacio vacìo
					// y termina el ciclo
					inicioCorrecto = objeto
					tamanoCorrecto = tamano
					break
				}

				if auxMbr.Fit == 'w' {
					// peor ajuste
					// compara el espacio cactual
					// con el espacio ya guardado
					// se queda con el espacio màs grande
					if tamano > tamanoCorrecto {
						inicioCorrecto = objeto
						tamanoCorrecto = tamano
					}
				}

				if auxMbr.Fit == 'b' {
					// mejor ajuste
					// compara el espacio actual con el anterior
					// se queda con el espacio màs pequeño
					if tamano < tamanoCorrecto {
						inicioCorrecto = objeto
						tamanoCorrecto = tamano
					}
				}
			}
		}

		/*
			verifica si se encontrò el inicio correcto para
			crear la nueva particiòn
		*/
		if inicioCorrecto != -1 {
			// sì se encontrò el espacio adecuado
			// crear los datos de la nueva particiòn
			auxMbr.Partitions[idParticionLibre].Start = int64(inicioCorrecto)
			auxMbr.Partitions[idParticionLibre].Size = int64(i.sizeBytes)
			auxMbr.Partitions[idParticionLibre].Status = 1
			auxMbr.Partitions[idParticionLibre].Fit = i.fit[0]
			auxMbr.Partitions[idParticionLibre].Type = i.tipo[0]
			copy(auxMbr.Partitions[idParticionLibre].Name[:], i.name)

			// escribir mbr
			escribirMBR(i.path+i.name, auxMbr)

		} else {
			// no se encontrò espacio adecuado
			fmt.Println("No se ha encontrado espacio libre adecuado para crear èsta particiòn")
		}
	} else {
		fmt.Println("Se ha alcanzado el lìmite de particiones primarias y extendidah.")
	}
}

func (i *fdisk) crearParticionL() {

}
