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
	ruta string
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
		case "ruta":
			i.ruta = QuitarComillas(p.Valor)
			break
		}
	}
}

func (i *rep) Validar() bool {
	retorno := true

	if i.path == "" || i.name == "" || i.id == "" {
		retorno = false
	}

	if i.name == "file" && i.ruta == "" {
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

		} else if i.name == "disk" {
			// es un reporte de espacios
			// es un reporte mbr descripciòn
			var auxMbr = RecuperarMBR(particionMontada.path)
			var contenido = getReporteDSK(auxMbr, particionMontada.path)

			// crear el archivo
			file, err := os.Create(i.path + ".dot")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()

		} else if i.name == "tree" {
			//var superBloque = particionMontada.sp
			var contenido = getReporteTree(*particionMontada)

			// crear el archivo
			file, err := os.Create(i.path + ".dot")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()

		} else if i.name == "sb" {
			//var superBloque = particionMontada.sp
			var contenido = getReporteSuperBloque(*particionMontada)

			// crear el archivo
			file, err := os.Create(i.path + ".dot")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()
		} else if i.name == "bm_inode" {
			//var superBloque = particionMontada.sp
			// recuperar el bitmap inodo
			_, bitmap := recuperarBitMap(particionMontada.path, particionMontada.sp.BitMapInodeStart, int64(particionMontada.sp.InodesCount))
			var contenido = getReporteBitMap(bitmap)

			// crear el archivo
			file, err := os.Create(i.path + ".txt")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()
		} else if i.name == "bm_block" {
			//var superBloque = particionMontada.sp
			// recuperar el bitmap inodo
			_, bitmap := recuperarBitMap(particionMontada.path, particionMontada.sp.BitMapBlockStart, int64(particionMontada.sp.BlocksCount))
			var contenido = getReporteBitMap(bitmap)

			// crear el archivo
			file, err := os.Create(i.path + ".txt")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()
		} else if i.name == "inode" {
			//var superBloque = particionMontada.sp
			// recuperar el bitmap inodo
			_, bitmap := recuperarBitMap(particionMontada.path, particionMontada.sp.BitMapInodeStart, int64(particionMontada.sp.InodesCount))
			var contenido = getReporteInodos(bitmap, *particionMontada)

			// crear el archivo
			file, err := os.Create(i.path + ".dot")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()
		} else if i.name == "block" {
			//var superBloque = particionMontada.sp
			// recuperar el bitmap inodo
			var contenido = getReporteBloques(particionMontada)

			// crear el archivo
			file, err := os.Create(i.path + ".dot")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()
		} else if i.name == "file" {
			//var superBloque = particionMontada.sp
			// recuperar el bitmap inodo

			var contenido = getReporteFile(i.ruta, particionMontada)

			// crear el archivo
			file, err := os.Create(i.path + ".txt")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()
		} else if i.name == "journaling" {
			//var superBloque = particionMontada.sp
			// recuperar el bitmap inodo

			var contenido = getReporteJournal(particionMontada)

			// crear el archivo
			file, err := os.Create(i.path + ".dot")
			if err != nil {
				log.Fatal(err)
			}

			file.WriteString(contenido)
			file.Close()
		}

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
		} else if strings.Contains(i.path, ".svg") {
			comando := exec.Command("dot", i.path+".dot", "-Tsvg", "-o", i.path)
			if err := comando.Run(); err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("Se ha creado el reporte con èxito")
		fmt.Println("\t" + i.path)
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

func getReporteDSK(mbr Mbr, path string) string {
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
			percent := float32(float32(tamano) * 100 / float32(mbr.Size))
			retorno += "|{ Libre " + fmt.Sprintf("%.2f", percent) + "% }"
		} else {
			// es una particiòn
			percent := float32(float32(mbr.Partitions[particionCorrecta].Size) * 100 / float32(mbr.Size))
			if mbr.Partitions[particionCorrecta].Type == 'p' {
				retorno += "|{ Primaria " + BytesToString(mbr.Partitions[particionCorrecta].Name[:]) + " " + fmt.Sprintf("%.2f", percent) + "% }"
			} else {
				retorno += "|{ Extendida " + BytesToString(mbr.Partitions[particionCorrecta].Name[:]) + " " + fmt.Sprintf("%.2f", percent) + "% |"

				// es extendida
				// recorrer todos los ebr's
				posActual := mbr.Partitions[particionCorrecta].Start
				//contador := 1
				for {
					exito, auxEbr := RecuperarEBR(path, posActual)

					if !exito {
						// sì encontrò el ebr
						if auxEbr.Status == 1 && auxEbr.Size > 0 {
							percent := float32(float32(auxEbr.Size) * 100 / float32(mbr.Size))
							retorno += " " + BytesToString(auxEbr.Name[:]) + " "
							retorno += fmt.Sprintf("%.2f", percent) + "% "
						}
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
				retorno += "}"
			}
		}
	}

	retorno += "\"];"
	retorno += "}"

	return retorno
}

func getReporteTree(particionMontada ParticionMontada) string {
	var retorno string

	retorno += "digraph G {\n"
	retorno += "	concentrate=True;"
	retorno += "	rankdir=LR;"
	retorno += "	node[shape=record];"

	labels, arrow, _ := getRecursiveTree(0, particionMontada)

	retorno += labels + "\n"
	retorno += arrow + "\n"

	retorno += "}"

	return retorno
}

func getReporteSuperBloque(paricionMontada ParticionMontada) (retorno string) {

	retorno += "digraph g{\nrankdir = LR;\nnode[shape = record, width = .1, heigth = .1, width = 1.5];\n" +
		"inodo0[label = \"{NOMBRE | VALOR} |{s_inodes_count| " + strconv.Itoa(int(paricionMontada.sp.InodesCount)) +
		" }|{s_blocks_count| " + strconv.Itoa(int(paricionMontada.sp.BlocksCount)) +
		" }|{s_free_blocks_count| " + strconv.Itoa(int(paricionMontada.sp.FreeBlocksCount)) +
		" }|{s_free_inodes_count| " + strconv.Itoa(int(paricionMontada.sp.FreeInodesCount)) +
		" }|{s_mtime| " + BytesToString(paricionMontada.sp.MountedTime[:]) +
		" }|{s_umtime| " + BytesToString(paricionMontada.sp.UnMountedTime[:]) +
		" }|{s_mnt_count| " + strconv.Itoa(int(paricionMontada.sp.MountedCount)) +
		" }|{s_magic| " + strconv.Itoa(int(paricionMontada.sp.Magic)) +
		" }|{s_inode_size| " + strconv.Itoa(int(paricionMontada.sp.InodeSize)) +
		" }|{s_block_size| " + strconv.Itoa(int(paricionMontada.sp.BlockSize)) +
		" }|{s_first_ino| " + strconv.Itoa(int(paricionMontada.sp.FirstInode)) +
		" }|{s_first_blo| " + strconv.Itoa(int(paricionMontada.sp.FirstBlock)) +
		" }|{s_bm_inodde_start| " + strconv.Itoa(int(paricionMontada.sp.InodeStart)) +
		" }|{s_bm_block_start| " + strconv.Itoa(int(paricionMontada.sp.BlockStart)) + " }\"]" +
		"\n}"

	return retorno
}

func getReporteBitMap(bitmap []byte) (retorno string) {

	for indice, bit := range bitmap {
		retorno += strconv.Itoa(int(bit)) + " "
		if (indice+1)%20 == 0 {
			retorno += "\n"
		}
	}
	return retorno
}

func getReporteInodos(bitmap []byte, particion ParticionMontada) (retorno string) {
	retorno += "digraph g{\n\trankdir = LR;\n\tnode[shape = record, width = .1, heigth = .1, width = 1.5];\n\t"
	arrows := ""
	for indice, bit := range bitmap {
		// recuperar
		if bit == 1 {
			// inodo activo, recueprar el inodo
			_, inodo := recuperarInodo(particion.path, int64(particion.sp.InodeStart)+int64(particion.sp.InodeSize)*int64(indice))

			// validar que sea un inodo activo
			if inodo.UID == 0 {
				continue
			}
			// crear el reporte inodo
			retorno += getLabelInodoData(int32(indice), inodo)
			if arrows != "" {
				arrows += " nd_i" + strconv.Itoa(int(indice)) + ":t \n"
			}
			arrows += "nd_i" + strconv.Itoa(int(indice)) + ":t ->"
		}
	}

	if len(arrows) > 12 {
		arrows = arrows[:len(arrows)-12]
	}

	retorno += "\n\n"
	retorno += arrows
	retorno += "\n\n}"
	return retorno
}

func getRecursiveTree(indiceInodo int32, particionMontada ParticionMontada) (labels string, arrows string, idSiguiente string) {

	// obtener el inodo
	var sp = particionMontada.sp
	_, inodo := recuperarInodo(particionMontada.path, sp.InodeStart+int64(int64(sp.InodeSize)*int64(indiceInodo)))
	if inodo.UID == 0 { //  validar que sea un inodo habilitado
		return "", "", ""
	}
	nombreLabel := "nd_i" + strconv.Itoa(int(indiceInodo))
	labels = getLabelInodo(indiceInodo, inodo)

	if inodo.Type == 0 {
		// es inodo de carpetas
		for indice, bloque := range inodo.Block {
			if indice == 0 && bloque != -1 {
				// el bloque carpeta existe y debemos bucarlo
				// recueprar el bloque de carpetas
				_, bloqueCarpeta := recuperarBloqueCarpeta(particionMontada.path, sp.BlockStart+int64(bloque)*int64(sp.BlockSize))
				//fmt.Println(bloqueCarpeta)
				nombre, labelT := getLabelCarpeta(bloque, bloqueCarpeta)
				labels += labelT + "\n"
				arrows += nombreLabel + ":a" + strconv.Itoa(indice) + " -> " + nombre + ":t\n"

				for ii, apuntador := range bloqueCarpeta.Content[2:] {
					if apuntador.PointerInode != -1 {
						labelT, arrowt, sigt := getRecursiveTree(apuntador.PointerInode, particionMontada)
						labels += labelT + "\n"
						arrows += nombre + ":a" + strconv.Itoa(ii+2) + " -> " + sigt + ":t" + "\n"

						arrows += arrowt
					}
				}
			} else if indice != 0 && indice < 13 && bloque != -1 {
				// el bloque carpeta existe y debemos bucarlo
				// recueprar el bloque de carpetas
				_, bloqueCarpeta := recuperarBloqueCarpeta(particionMontada.path, sp.BlockStart+int64(bloque)*int64(sp.BlockSize))
				//fmt.Println(bloqueCarpeta)
				nombre, labelT := getLabelCarpeta(bloque, bloqueCarpeta)
				labels += labelT + "\n"
				arrows += nombreLabel + ":a" + strconv.Itoa(indice) + " -> " + nombre + ":t\n"

				for ii, apuntador := range bloqueCarpeta.Content {
					if apuntador.PointerInode != -1 {
						labelT, arrowt, sigt := getRecursiveTree(apuntador.PointerInode, particionMontada)
						labels += labelT + "\n"
						arrows += nombre + ":a" + strconv.Itoa(ii) + " -> " + sigt + ":t" + "\n"

						arrows += arrowt
					}
				}
			} else if indice == 13 && bloque != -1 {
				// llamar al bloque indirecto
				// por cada indirecto llamar a cada carpeta

			} else if indice == 14 && bloque != -1 {
				// por cada apuntador del indirecto, llamar al segundo bloque indirecto
				//por cada indirectsecundario llamar a carpetas llamar a recursivas
			}
		}
	} else {
		// es inodo de archivos
		for indice, bloque := range inodo.Block {
			if indice < 13 && bloque != -1 {
				// obtener el bloque archivo
				_, bloqueArchivo := recuperarBloqueArchivo(particionMontada.path, sp.BlockStart+int64(sp.BlockSize)*int64(bloque))
				nombre, labelT := getLabelArchivo(bloque, bloqueArchivo)

				labels += labelT
				arrows += nombreLabel + ":a" + strconv.Itoa(indice) + " -> " + nombre + ":p\n"
			} else if indice == 13 && bloque != -1 {
				// recueprar bloque indirecto
				_, bloqueIndirecto := recuperarBloqueIndirecto(&particionMontada, int64(bloque))

				nombreIndirecto, labelT := getLabelIndirecto(bloque, bloqueIndirecto)
				labels += labelT
				arrows += nombreLabel + ":a" + strconv.Itoa(indice) + " -> nd_bi" + strconv.Itoa(int(bloque)) + ":p\n"

				// recorer apuntadores indirecto
				for indice2, apuntador2 := range bloqueIndirecto.Apuntadores {
					if apuntador2 != -1 {
						// obtener el bloque archivo
						_, bloqueArchivo := recuperarBloqueArchivo(particionMontada.path, sp.BlockStart+int64(sp.BlockSize)*int64(apuntador2))
						nombre, labelT := getLabelArchivo(apuntador2, bloqueArchivo)

						labels += labelT
						arrows += nombreIndirecto + ":a" + strconv.Itoa(indice2) + " -> " + nombre + ":p\n"
					}
				}

			}
		}
	}

	return labels, arrows, nombreLabel
}

func getLabelIndirecto(indiceIndirecto int32, indirecto BloqueApuntadores) (string, string) {
	strIndice := strconv.Itoa(int(indiceIndirecto))
	retorno := "nd_bi" + strIndice +
		"[label=\"<t>indirecto: " + strIndice

	for indice, bloque := range indirecto.Apuntadores {
		strI := strconv.Itoa(indice)
		retorno += "|<a" + strI + ">a" + strI + ": " + strconv.Itoa(int(bloque))
	}

	retorno += "\"]\n"
	return ("nd_bi" + strIndice), retorno
}

func getLabelInodo(indiceInodo int32, inodo Inodo) string {
	strIndice := strconv.Itoa(int(indiceInodo))
	retorno := "nd_i" + strIndice +
		"[label=\"<t>inodo: " + strIndice +
		"|tipo: " + strconv.Itoa(int(inodo.Type)) +
		"|tamano: " + strconv.Itoa(int(inodo.Size))

	for indice, bloque := range inodo.Block {
		strI := strconv.Itoa(indice)
		retorno += "|<a" + strI + ">a" + strI + ": " + strconv.Itoa(int(bloque))
	}

	retorno += "\"]\n"
	return retorno
}

func getLabelInodoData(indiceInodo int32, inodo Inodo) string {
	strIndice := strconv.Itoa(int(indiceInodo))
	retorno := "nd_i" + strIndice +
		"[label=\"<t>inodo: " + strIndice +
		"|tipo: " + strconv.Itoa(int(inodo.Type)) +
		"|tamano: " + strconv.Itoa(int(inodo.Size)) +
		"|uid: " + strconv.Itoa(int(inodo.UID)) +
		"|gid: " + strconv.Itoa(int(inodo.GID)) +
		"|atime: " + BytesToString(inodo.Atmie[:]) +
		"|ctime: " + BytesToString(inodo.Ctime[:]) +
		"|mtime: " + BytesToString(inodo.Mtime[:]) +
		"|perms: " + strconv.Itoa(int(inodo.Perm[0])) + strconv.Itoa(int(inodo.Perm[1])) + strconv.Itoa(int(inodo.Perm[2]))

	for indice, bloque := range inodo.Block {
		strI := strconv.Itoa(indice)
		retorno += "|<a" + strI + ">block_" + strI + ": " + strconv.Itoa(int(bloque))
	}

	retorno += "\"]\n"
	return retorno
}

func getLabelCarpeta(indiceInodo int32, inodo BloqueCarpeta) (string, string) {
	strIndice := strconv.Itoa(int(indiceInodo))
	retorno := "nd_b" + strIndice +
		"[label=\"<t>carpeta: " + strIndice

	for indice, bloque := range inodo.Content {
		strI := strconv.Itoa(indice)
		retorno += "|<a" + strI + ">" + BytesToString(bloque.Name[:]) + strI + ": " + strconv.Itoa(int(bloque.PointerInode))
	}

	retorno += "\"]"

	apuntador := "nd_b" + strIndice
	return apuntador, retorno
}

func getLabelArchivo(indiceInodo int32, bloqueArchivo BloqueArchivo) (string, string) {
	strIndice := strconv.Itoa(int(indiceInodo))
	retorno := "nd_b" + strIndice +
		"[label=\"<t>archivo: " + strIndice

	retorno2 := BytesToString(bloqueArchivo.Content[:])
	retorno2 = strings.ReplaceAll(retorno2, "\n", "\\n")

	retorno += "|" + retorno2
	retorno += "\"]\n"

	apuntador := "nd_b" + strIndice
	return apuntador, retorno
}

func getReporteBloques(particion *ParticionMontada) (retorno string) {

	retorno += "digraph g{\n\trankdir = LR;\n\tnode[shape = record, width = .1, heigth = .1, width = 1.5];\n\t"

	// obtener los labels de todos los bloques activos en el àrbol de directorios
	contenido := getRecursiveTreeBlock(0, *particion)
	retorno += contenido

	// ordenar la posiciòn de las flechas
	_, bitmap := recuperarBitMap(particion.path, particion.sp.BitMapBlockStart, int64(particion.sp.BlocksCount))
	for indice, bit := range bitmap {
		// recorrer el bitmap
		// graficar todos los bits con valor 1
		if bit == 1 {
			// buscar bloque en el mapa
			retorno += "nd_b" + strconv.Itoa(indice) + " -> "
		}
	}

	retorno = retorno[:len(retorno)-4]

	retorno += "\n\n}"
	return retorno
}

func getRecursiveTreeBlock(indiceInodo int32, particionMontada ParticionMontada) (labels string) {

	// obtener el inodo
	var sp = particionMontada.sp
	_, inodo := recuperarInodo(particionMontada.path, sp.InodeStart+int64(int64(sp.InodeSize)*int64(indiceInodo)))
	if inodo.UID == 0 { //  validar que sea un inodo habilitado
		return ""
	}
	if inodo.Type == 0 {
		// es inodo de carpetas
		for indice, bloque := range inodo.Block {
			if indice == 0 && bloque != -1 {
				// el bloque carpeta existe y debemos bucarlo
				// recueprar el bloque de carpetas
				_, bloqueCarpeta := recuperarBloqueCarpeta(particionMontada.path, sp.BlockStart+int64(bloque)*int64(sp.BlockSize))
				_, labelTmp := getLabelCarpeta(bloque, bloqueCarpeta)
				labels += labelTmp + "\n"

				for _, apuntador := range bloqueCarpeta.Content[2:] {
					if apuntador.PointerInode != -1 {
						labels2 := getRecursiveTreeBlock(apuntador.PointerInode, particionMontada)
						labels += labels2
					}
				}
			} else if indice != 0 && indice < 13 && bloque != -1 { // indice > 0; indice < 13
				// el bloque carpeta existe y debemos bucarlo
				// recueprar el bloque de carpetas
				_, bloqueCarpeta := recuperarBloqueCarpeta(particionMontada.path, sp.BlockStart+int64(bloque)*int64(sp.BlockSize))
				_, labelTmp := getLabelCarpeta(bloque, bloqueCarpeta)
				labels += labelTmp + "\n"

				for _, apuntador := range bloqueCarpeta.Content {
					if apuntador.PointerInode != -1 {
						labels2 := getRecursiveTreeBlock(apuntador.PointerInode, particionMontada)
						labels += labels2
					}
				}
			} else if indice == 13 && bloque != -1 {
				// llamar al bloque indirecto
				// por cada indirecto llamar a cada carpeta
			} else if indice == 14 && bloque != -1 {
				// por cada apuntador del indirecto, llamar al segundo bloque indirecto
				//por cada indirectsecundario llamar a carpetas llamar a recursivas
			}
		}
	} else {
		// es inodo de archivos
		for indice, bloque := range inodo.Block {
			if indice < 13 && bloque != -1 {
				// obtener el bloque archivo
				_, bloqueArchivo := recuperarBloqueArchivo(particionMontada.path, particionMontada.sp.BlockStart+int64(particionMontada.sp.BlockSize)*int64(bloque))
				_, labelBloqueArchivo := getLabelArchivo(bloque, bloqueArchivo)
				labels += labelBloqueArchivo
			}
		}
	}

	return labels
}

func getReporteFile(ruta string, particion *ParticionMontada) (retorno string) {

	if UsuarioActualLogueado.UID == 0 {
		fmt.Println("\nDebes estar logueado para crear èste reporte, (no se pueden leer permisos)")
		return
	}

	// recuperar el primer inodo
	_, inodo := recuperarInodo(particion.path, particion.sp.InodeStart)
	//var indiceInodo = 0
	// recorrer todas las carpetas
	pathSplit := strings.Split(ruta, "/")
	for indice, carpeta := range pathSplit {
		if indice != 0 && indice != (len(pathSplit) /*-1*/) {
			// recuperar carpeta del inodo
			_, _, iInodoCarpetaSiguiente := getCarpetaFromInodo(carpeta, inodo, *UsuarioActualLogueado.particion)
			if iInodoCarpetaSiguiente != -1 {
				// la carpeta existe
				_, inodo = recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart+int64(UsuarioActualLogueado.particion.sp.InodeSize)*int64(iInodoCarpetaSiguiente))
			} else {
				// la carpeta no existe
				fmt.Println("\tNo se ha encontrado el directorio " + carpeta)
			}
		}
	}
	// estamos en el inodo del archivo
	contenidoArchivo := getContenidoArchivo(inodo, *particion)

	// ya tenemos el contenido
	retorno = pathSplit[len(pathSplit)-1]
	retorno += "\n" + contenidoArchivo
	return retorno
}

func getReporteJournal(particion *ParticionMontada) (retorno string) {

	retorno += "digraph g{\n\trankdir = LR;\n\tnode[shape = record, width = .1, heigth = .1, width = 1.5];\n\t"

	arrows := ""

	// RECORRER TODOS LOS BLQOUES DE TIPO JOURNAL
	for indice := int32(0); indice < particion.sp.JournalCount; indice++ {
		// recuperar bloque journal
		_, journal := recuperarJournal(particion.path, indice, *particion)

		retorno += "nd_j" + strconv.Itoa(int(indice)) +
			"[label=\"<t>Journal : " + strconv.Itoa(int(indice)) +
			"|Acciòn : " + strconv.Itoa(int(journal.TipoOperacion)) +
			"|Tipo: " + strconv.Itoa(int(journal.Tipo)) +
			"|nombre: " + BytesToString(journal.Nombre[:]) +
			"|contenido: " + BytesToString(journal.Content[:]) +
			"|fecha: " + BytesToString(journal.Fecha[:]) +
			"|propietario: " + strconv.Itoa(int(journal.Propietario[0])) + strconv.Itoa(int(journal.Propietario[0])) +
			"|permisos: " + BytesToString([]byte{6, 6, 4})
		retorno += "\"]\n"
		arrows += "nd_j" + strconv.Itoa(int(indice)) + " -> "
	}

	if len(arrows) > 4 {
		arrows = arrows[:len(arrows)-4]
	}

	retorno += "\n\n"
	retorno += arrows
	retorno += "\n\n}"
	return retorno
}
