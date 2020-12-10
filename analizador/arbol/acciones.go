package arbol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func RecuperarMBR(path string) Mbr {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// preparar estructura mbr
	mbr := Mbr{}
	mbr.Size = -1

	var tamano = binary.Size(mbr)
	datos := ReadNextBytes(file, tamano)
	buffer := bytes.NewBuffer(datos)
	err = binary.Read(buffer, binary.BigEndian, &mbr)
	if err != nil {
		log.Fatal(err)
	}

	return mbr
}

func escribirMBR(path string, mbr Mbr) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// escribir la estructura
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, &mbr)
	WriteNextBytes(file, binario.Bytes())
}

func escribirEBR(path string, ebr Ebr, seek int64) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// escribir la estructura
	file.Seek(seek, 0)
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, &ebr)
	WriteNextBytes(file, binario.Bytes())
}

func RecuperarEBR(path string, seek int64) (bool, Ebr) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// preparar estructura mbr
	ebr := Ebr{}
	ebr.Size = -1

	file.Seek(seek, 0)
	var tamano = binary.Size(ebr)
	datos := ReadNextBytes(file, tamano)
	buffer := bytes.NewBuffer(datos)
	err = binary.Read(buffer, binary.BigEndian, &ebr)
	if err != nil {
		log.Fatal(err)
	}

	return false, ebr
}

func ReadNextBytes(file *os.File, size int) []byte {
	bytes := make([]byte, size)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func WriteNextBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func getFechaByte() [20]byte {
	var fechab [20]byte
	var fecha = time.Now()

	fechastr := strconv.Itoa(int(fecha.Year()))
	fechastr += "/"
	fechastr += strconv.Itoa(int(fecha.Month()))
	fechastr += "/"
	fechastr += strconv.Itoa(int(fecha.Day()))
	fechastr += "-"
	fechastr += strconv.Itoa(int(fecha.Hour()))
	fechastr += ":"
	fechastr += strconv.Itoa(int(fecha.Minute()))
	fechastr += ":"
	fechastr += strconv.Itoa(int(fecha.Second()))

	copy(fechab[:], fechastr)
	fmt.Println(fechastr)
	return fechab
}

func BytesToString(b []byte) string {
	cadena := ""
	for i := 0; i < len(b); i++ {
		if b[i] != 0 {
			cadena += string(b[i])
		}
	}
	return cadena
}

func ByteToString(b byte) string {
	cadena := ""
	if b != 0 {
		cadena += string(b)
	}
	return cadena
}

/*buscar espacios libres en un disco*/
type Espacios struct {
	Inicios []int
	Finales []int
}

func getEspaciosLibres(mbr Mbr) Espacios {
	var espaciosVacios Espacios
	var iniciaLibre int = binary.Size(mbr) // inicia el espacio libre
	// luego de la escritura del mbr

	// recorrer las particiones activas
	var espaciosLlenos Espacios
	for i := 0; i <= 3; i++ {
		espaciosLlenos.Inicios = append(espaciosLlenos.Inicios, int(mbr.Partitions[i].Start))
		espaciosLlenos.Finales = append(espaciosLlenos.Finales, int(mbr.Partitions[i].Start+mbr.Partitions[i].Size))
	}

	// ordenar espacios llenos de menor a mayor
	sort.Ints(espaciosLlenos.Inicios)
	sort.Ints(espaciosLlenos.Finales)

	// encontrar espacios vaciòs
	var posicionActual int = int(iniciaLibre)
	for indice, objeto := range espaciosLlenos.Inicios {
		if objeto > posicionActual {
			// hay un espacio libre
			espaciosVacios.Inicios = append(espaciosVacios.Inicios, int(posicionActual))
			espaciosVacios.Finales = append(espaciosVacios.Finales, int(objeto-1))
			posicionActual = espaciosLlenos.Finales[indice]
		} else if objeto == posicionActual {
			posicionActual = espaciosLlenos.Finales[indice]
		}
	}

	if posicionActual < int(mbr.Size) {
		espaciosVacios.Inicios = append(espaciosVacios.Inicios, int(posicionActual))
		espaciosVacios.Finales = append(espaciosVacios.Finales, int(mbr.Size))
	}

	return espaciosVacios
}

func getEspaciosLlenos(mbr Mbr) {
	var espacios Espacios

	for _, paricion := range mbr.Partitions {
		espacios.Inicios = append(espacios.Inicios, int(paricion.Start))
		espacios.Finales = append(espacios.Finales, int(paricion.Start+paricion.Size))
	}
}
