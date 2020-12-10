package arbol

import (
	"fmt"
	"log"
	"os"
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
	addBytes      int
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
		i.unit = "k"
	}

	if i.fit == "" {
		i.fit = "wf"
	}

	if i.tipo == "" {
		i.tipo = "p"
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
	i := fdisk{}
	i.MatchParametros(p)
	if i.Validar() {
		// identificar què tipo de ejecuciòn se debe realizar

		// crear una particiòn
		if i.size != 0 {
			// se està creando una particiòn
			if i.fit != "" {
				i.fit = "ff"
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
			i.addBytes = i.add * i.multiplicador

			// ejecutar mètodo para crear la particiòn
			if i.tipo == "p" || i.tipo == "e" {
				i.crearParticionPE()
			} else {
				i.crearParticionL()
			}

			return
		}

		if i.add != 0 {
			// editando el tamaño de una particiòn
			if i.unit == "" {
				i.unit = "k"
			}

			// calcular el tamaño en bytes de la nueva particiòn
			if i.unit == "m" {
				i.multiplicador = 1024 * 1024
			} else if i.unit == "k" {
				i.multiplicador = 1024
			}

			i.addBytes = i.add * i.multiplicador

			i.addParticion()
			return
		}

		if i.delete != "" {
			// eliminando una particiòn
			i.eliminarParticion()
			return
		}

	}

}

/*acciones de fdisk*/
func (i *fdisk) crearParticionPE() {
	// recuperar el mbr
	i.mbr = RecuperarMBR(i.path)
	auxMbr := i.mbr

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

			if i.tipo[0] == 'e' {
				auxEbr := Ebr{0, 'f', 0, 0, -1, [16]byte{}}
				escribirEBR(i.path, auxEbr, int64(inicioCorrecto))
			}

			// escribir mbr
			escribirMBR(i.path, auxMbr)
			fmt.Println("Particiòn creada con èxito")

		} else {
			// no se encontrò espacio adecuado
			fmt.Println("No se ha encontrado espacio libre adecuado para crear èsta particiòn")
		}
	} else {
		fmt.Println("Se ha alcanzado el lìmite de particiones primarias y extendidah.")
	}
}

func (i *fdisk) crearParticionL() {
	// recuperar el mbr
	i.mbr = RecuperarMBR(i.path)
	auxMbr := i.mbr
	//fmt.Println(i.mbr)

	// encontrar la particiòn con el tipo Extendido
	var nombreByte [16]byte
	copy(nombreByte[:], i.name)
	for _, particion := range auxMbr.Partitions {
		if particion.Type == 'e' {
			// sì existe la particiòn exentendida

		}
	}
}

func (i *fdisk) eliminarParticion() {
	// recuperar el mbr
	i.mbr = RecuperarMBR(i.path)
	auxMbr := i.mbr
	//fmt.Println(i.mbr)

	// encontrar la particiòn con el nombre a eliminar
	var nombreByte [16]byte
	copy(nombreByte[:], i.name)
	for indice := 0; indice <= 3; indice++ {
		auxParticion := auxMbr.Partitions[indice]
		if auxParticion.Status == 1 && auxParticion.Name == nombreByte {
			// modificar mbr para dejar estado eliminado
			auxMbr.Partitions[indice].Status = 2

			if i.delete == "full" {
				// llenar de ceros el espacio de la particiòn
				file, err := os.OpenFile(i.path, os.O_RDWR, 0777)
				if err != nil {
					log.Fatal(err)
				}
				file.Seek(auxParticion.Start, 0)
				for j := 0; j < int(auxParticion.Size); j++ {
					file.Write([]byte{0})
				}
				file.Close()
			}
			escribirMBR(i.path, auxMbr)
			fmt.Println("Particiòn eliminada con èxito")
			break
		}
	}
}

func (i *fdisk) addParticion() {
	// recuperar el mbr
	i.mbr = RecuperarMBR(i.path)
	auxMbr := i.mbr
	//fmt.Println(i.mbr)

	// encontrar la particiòn con el nombre a eliminar
	var nombreByte [16]byte
	copy(nombreByte[:], i.name)
	for indice := 0; indice <= 3; indice++ {
		auxParticion := auxMbr.Partitions[indice]
		if auxParticion.Status == 1 && auxParticion.Name == nombreByte {
			// modificar mbr

			// si es positivo
			if i.add > 0 {
				espaciosVacios := getEspaciosLibres(auxMbr)
				finParticion := auxParticion.Start + auxParticion.Size

				for indice2, inicio := range espaciosVacios.Inicios {
					if finParticion == int64(inicio) {
						// espacio contiguo
						if i.addBytes <= espaciosVacios.Finales[indice2]-inicio {
							// se puede añadir espacio
							auxMbr.Partitions[indice].Size += int64(i.addBytes)

							// guardar mbr
							escribirMBR(i.path, auxMbr)
							fmt.Println("Particiòn editada con èxito")
						}
					}
				}
			} else {
				// si add es negativo
				if int64(i.addBytes) < auxParticion.Size {
					// reducir la particiòn
					auxMbr.Partitions[indice].Size += int64(i.addBytes)

					// guardar mbr
					escribirMBR(i.path, auxMbr)
					fmt.Println("Particiòn editada con èxito")
				}
			}

			break
		}
	}
}
