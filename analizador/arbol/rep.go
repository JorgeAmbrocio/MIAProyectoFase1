package arbol

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type rep struct {
	name string
	path string
	id   string
}

func (i *rep) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "path":
			i.path = QuitarComillas(p.Valor)
			spliteado := strings.Split(i.path, "/")
			var directorio string = ""
			for _, str := range spliteado[1 : len(spliteado)-1] {
				directorio += "/" + str
			}
			CrearTodasCarpetas(directorio)
			break
		case "name":
			i.name = p.Valor
			break
		case "id":
			i.id = p.Valor
			break
		}
	}
}

func (i *rep) Validar() bool {
	retorno := true

	if i.path == "" || i.name == "" || i.id == "" {
		retorno = false
	}

	return retorno
}

func Erep(p []Parametro) {
	i := rep{}
	i.MatchParametros(p)
	if i.Validar() {
		i.crearReporte()
	}
}

func (i *rep) crearReporte() {

	// encontrar el id a buscar
	if existe, particionMontada := RecuperarParticionMontada(i.id); existe {
		if i.name == "mbr" {
			// es un reporte mbr descripciòn
			var auxMbr = RecuperarMBR(particionMontada.path)
			var contenido = getReporteMBR(auxMbr, particionMontada.path)

			// crear el archivo
			file, err := os.Create(i.path + ".dot")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()

			// compilar el archivo creado
			comando := exec.Command("dot", i.path+".dot", "-Tjpg", "-o", i.path)
			if err := comando.Run(); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Se ha creado el reporte con èxito")
			fmt.Println("\t" + i.path)
		} else {
			// es un reporte de espacios
			// es un reporte mbr descripciòn
			var auxMbr = RecuperarMBR(particionMontada.path)
			var contenido = getReporteDSK(auxMbr)

			// crear el archivo
			file, err := os.Create(i.path + ".dot")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()

			// compilar el archivo creado
			if strings.Contains(i.path, ".jpg") {
				comando := exec.Command("dot", i.path+".dot", "-Tjpg", "-o", i.path)
				if err := comando.Run(); err != nil {
					fmt.Println(err)
				}
			} else if strings.Contains(i.path, ".png") {
				comando := exec.Command("dot", i.path+".dot", "-Tpng", "-o", i.path)
				if err := comando.Run(); err != nil {
					fmt.Println(err)
				}
			} else if strings.Contains(i.path, ".pdf") {
				comando := exec.Command("dot", i.path+".dot", "-Tpdf", "-o", i.path)
				if err := comando.Run(); err != nil {
					fmt.Println(err)
				}
			}

			fmt.Println("Se ha creado el reporte con èxito")
			fmt.Println("\t" + i.path)
		}
	}
}

func getReporteMBR(mbr Mbr, path string) string {
	var retorno = ""

	retorno += "digraph test {\n"
	retorno += "	graph [ratio=fill];\n"
	retorno += "	node [label=\"\\N\", fontsize=15, shape=plaintext];\n"
	retorno += "	graph [bb=\"0,0,352,154\"];\n"

	/*formato dot para mbr*/
	retorno += "	arset [label=<\n"
	retorno += "		<table>\n"
	retorno += "		<tr><td>Atributo</td><td>Valor</td></tr>\n"
	// contenido de los valores para la tabla
	retorno += "		<tr><td>mbr_fecha_creacion</td><td>" + BytesToString(mbr.Date[:]) + "</td></tr>\n"
	retorno += "		<tr><td>mbr_disk_signature</td><td>" + strconv.Itoa(int(mbr.Signature)) + "</td></tr>\n"
	retorno += "		<tr><td>mbr_size</td><td>" + strconv.Itoa(int(mbr.Size)) + "</td></tr>\n"

	for j := 0; j < 4; j++ {
		retorno += "		<tr><td>part_status_" + strconv.Itoa(j+1) + "</td><td>" + strconv.Itoa(int(mbr.Partitions[j].Status)) + "</td></tr>\n"
		retorno += "		<tr><td>part_type_" + strconv.Itoa(j+1) + "</td><td>" + ByteToString(mbr.Partitions[j].Type) + "</td></tr>\n"
		retorno += "		<tr><td>part_fit_" + strconv.Itoa(j+1) + "</td><td>" + ByteToString(mbr.Partitions[j].Fit) + "</td></tr>\n"
		retorno += "		<tr><td>part_start_" + strconv.Itoa(j+1) + "</td><td>" + strconv.Itoa(int(mbr.Partitions[j].Start)) + "</td></tr>\n"
		retorno += "		<tr><td>part_size_" + strconv.Itoa(j+1) + "</td><td>" + strconv.Itoa(int(mbr.Partitions[j].Size)) + "</td></tr>\n"
		retorno += "		<tr><td>part_name_" + strconv.Itoa(j+1) + "</td><td>" + BytesToString(mbr.Partitions[j].Name[:]) + "</td></tr>\n"
	}

	retorno += "		</table>\n"
	retorno += "	>, ];\n"

	/*formato dot para ebr*/
	for j := 0; j < 4; j++ {
		if mbr.Partitions[j].Type == 'e' {
			// es extendida
			// recorrer todos los ebr's
			posActual := mbr.Partitions[j].Start
			contador := 1
			for {
				exito, auxEbr := RecuperarEBR(path, posActual)

				if !exito {
					// sì encontrò el ebr
					retorno += "	" + BytesToString(auxEbr.Name[:]) + " [label=<\n"
					retorno += "		<table>\n"
					retorno += "		<tr><td>Atributo</td><td>Valor</td></tr>\n"

					// crear el contenido para el ebr
					retorno += "		<tr><td>part_status_" + strconv.Itoa(contador+1) + "</td><td>" + strconv.Itoa(int(auxEbr.Status)) + "</td></tr>\n"
					retorno += "		<tr><td>part_type_" + strconv.Itoa(contador+1) + "</td><td>L</td></tr>\n"
					retorno += "		<tr><td>part_fit_" + strconv.Itoa(contador+1) + "</td><td>" + ByteToString(auxEbr.Fit) + "</td></tr>\n"
					retorno += "		<tr><td>part_start_" + strconv.Itoa(contador+1) + "</td><td>" + strconv.Itoa(int(auxEbr.Start)) + "</td></tr>\n"
					retorno += "		<tr><td>part_size_" + strconv.Itoa(contador+1) + "</td><td>" + strconv.Itoa(int(auxEbr.Size)) + "</td></tr>\n"
					retorno += "		<tr><td>part_name_" + strconv.Itoa(contador+1) + "</td><td>" + BytesToString(auxEbr.Name[:]) + "</td></tr>\n"

					retorno += "		</table>\n"
					retorno += "	>, ];\n"

					// indicar la posiciòn siguiente
					if auxEbr.Next == -1 {
						break
					} else {
						posActual = auxEbr.Next
					}
				} else {
					break
				}
			}
			break
		}
	}

	retorno += "}\n"
	return retorno
}

func getReporteDSK(mbr Mbr) string {
	var retorno = ""

	retorno += "digraph G {\n"
	retorno += "	concentrate=True;"
	retorno += "	rankdir=TB;"
	retorno += "	node [shape=record];"
	//retorno += "	hoja [label=\"mbr\n|{Extendida: 70% |}|{ Primaria 25% } | { Libre 5% }\"];"
	retorno += "	hojasa [label=\"mbr\n"

	// obtener espacios vacìos
	//espaciosVacios := getEspaciosLibres(mbr)
	espaciosTotales := getEspaciosLibres(mbr)
	for _, particion := range mbr.Partitions { // obtener espacios llenos
		espaciosTotales.Inicios = append(espaciosTotales.Inicios, int(particion.Start))
		espaciosTotales.Finales = append(espaciosTotales.Finales, int(particion.Start+particion.Size))
	}

	//  ordenar todos los espacios
	sort.Ints(espaciosTotales.Inicios)
	sort.Ints(espaciosTotales.Finales)

	// encontrar tamaños
	// recorrer todos los espacios
	for indice, espacio := range espaciosTotales.Inicios {
		// omitir los ceros
		if espacio == 0 {
			continue
		}

		// identificar si el espacio es libre o particiòn
		var particionCorrecta int = -1
		for in, part := range mbr.Partitions {
			if part.Start == int64(espacio) && part.Status == 1 {
				particionCorrecta = in
				break
			}
		}

		if particionCorrecta == -1 {
			// es un espacio libre
			tamano := espaciosTotales.Finales[indice] - espacio
			retorno += "|{ Libre " + strconv.Itoa(int(int(tamano)*100/int(mbr.Size))) + "% }"
		} else {
			// es una particiòn
			if mbr.Partitions[particionCorrecta].Type == 'p' {
				retorno += "|{ Primaria " + BytesToString(mbr.Partitions[particionCorrecta].Name[:]) + " " + strconv.Itoa(int(int(mbr.Partitions[particionCorrecta].Size)*100/int(mbr.Size))) + "% }"
			} else {
				retorno += "|{ Extendida " + BytesToString(mbr.Partitions[particionCorrecta].Name[:]) + " " + strconv.Itoa(int(int(mbr.Partitions[particionCorrecta].Size)*100/int(mbr.Size))) + "% |}"
			}
		}
	}

	retorno += "\"];"
	retorno += "}"

	return retorno
}
