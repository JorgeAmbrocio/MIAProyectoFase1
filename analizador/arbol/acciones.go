package arbol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
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

func ReadNextBytes(file *os.File, size int) []byte {
	bytes := make([]byte, size)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
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
