package arbol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

// mkdisk estructura de la instrucción mkdisk
type mkdisk struct {
	size          int
	path          string
	name          string
	unit          string
	fit           string
	multiplicador int
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *mkdisk) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch p.Tipo {
		case "size":
			if valor, err := strconv.Atoi(p.Valor); err != nil {
				fmt.Println("ERROR MKDISK: size no es válido ", p.Valor)
			} else {
				i.size = valor
			}
			break
		case "path":
			i.path = QuitarComillas(p.Valor)
			CrearTodasCarpetas(i.path)
			break
		case "name":
			i.name = p.Valor
			break
		case "unit":
			i.unit = p.Valor
			break
		case "fit":
			i.unit = p.Valor
		}
	}

	if i.unit == "" {
		i.unit = "m"
		i.multiplicador = 1024 * 1024
	} else {
		i.multiplicador = 1024
	}

	if i.fit == "" {
		i.fit = "wf"
	}

}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i mkdisk) Validar() bool {
	retorno := true

	if i.size < 0 || i.path == "" || i.name == "" {
		retorno = false
	}

	return retorno
}

// CrearBinario ...
func (i mkdisk) CrearBinario() {
	file, err := os.Create(i.path + i.name)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	// llenar el archivo de ceros
	var auxTamano [1024]byte // 1 kilobyte

	for contador := 0; contador < (i.multiplicador/1024)*i.size; contador++ {
		var binario bytes.Buffer
		binary.Write(&binario, binary.BigEndian, &auxTamano)
		WriteNextBytes(file, binario.Bytes())
	}

	// regresar putero a la posiciòn inicial
	file.Seek(0, 0)

	// crear estructura
	mbr := Mbr{
		Size:      int32(i.size * i.multiplicador),
		Signature: int32(rand.Intn(1000)),
		Date:      getFechaByte(),
		Fit:       i.fit[0],
	}

	// escribir la estructura
	dirMbr := &mbr
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, dirMbr)
	WriteNextBytes(file, binario.Bytes())

	file.Close()
}

// Emkdisk es la ejecución del mkdisk
func Emkdisk(p []Parametro) {
	/*
		crear objeto
		incluir parámetros
		validar si cuenta con lo suficiente para ejecutarse
	*/
	i := mkdisk{}
	i.MatchParametros(p)
	if i.Validar() {
		i.CrearBinario()
	}
}
