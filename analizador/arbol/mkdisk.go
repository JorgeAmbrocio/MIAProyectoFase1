package arbol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// mkdisk estructura de la instrucción mkdisk
type mkdisk struct {
	size int
	path string
	//name          string
	unit          string
	fit           string
	multiplicador int
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *mkdisk) MatchParametros(lp []Parametro) {
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
			i.unit = strings.ToLower(p.Valor)
			break
		case "fit":
			i.fit = strings.ToLower(p.Valor)
		}
	}

	if i.unit == "" {
		i.unit = "m"
	}

	if i.fit == "" {
		i.fit = "wf"
	}

}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i mkdisk) Validar() bool {
	retorno := true

	if i.size < 0 || i.path == "" {
		retorno = false
	}

	return retorno
}

// CrearBinario ...
func (i mkdisk) CrearBinario() {
	file, err := os.Create(i.path)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	}

	// llenar el archivo de ceros
	if i.unit == "m" {
		i.multiplicador = 1024 * 1024
	} else {
		i.multiplicador = 1024
	}

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

	fmt.Println("Disco creado con èxito")
	fmt.Println("\t" + i.path)
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
