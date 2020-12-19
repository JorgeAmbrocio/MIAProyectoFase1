package arbol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*MBR*/
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

/*EBR*/
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
		fmt.Println(err)
		return true, ebr
	}

	return false, ebr
}

/*SUPER BLOQUE*/
/*
	NÙMERO DE ESTRUCTURAS
	tamaño_particion = sizeof(superblock) + n + n*Sizeof(Journaling) + 3 * n + n *sizeof(inodos) + 3 * n * Sizeof(block)
	numero_estructuras = floor(n)
*/
func escribirSuperBloque(path string, sp SuperBlock, seek int64) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	// escribir la estructura
	file.Seek(seek, 0)
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, &sp)
	WriteNextBytes(file, binario.Bytes())
}

func recuperarSuperBloque(path string, seek int64) (bool, SuperBlock) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// preparar estructura mbr
	sp := SuperBlock{}
	//ebr.Size = -1

	file.Seek(seek, 0)
	var tamano = binary.Size(sp)
	datos := ReadNextBytes(file, tamano)
	buffer := bytes.NewBuffer(datos)
	err = binary.Read(buffer, binary.BigEndian, &sp)
	if err != nil {
		fmt.Println(err)
		return false, sp
	}

	return true, sp
}

func crearSuperBloque(particion ParticionMontada) SuperBlock {
	var n int = int(particion.Size) - binary.Size(SuperBlock{})
	n = n / (binary.Size(Journaling{}) + 4 + binary.Size(Inodo{}) + 3*binary.Size(BloqueCarpeta{}))

	pointerBmi := int(particion.Start) + binary.Size(SuperBlock{}) + n*binary.Size(Journaling{})
	pointerBmb := pointerBmi + n
	pointerInodoStart := pointerBmb + 3*n
	pointerBlockStart := pointerInodoStart + n*binary.Size(Inodo{})
	return SuperBlock{
		Type:            3,
		InodesCount:     int32(n),
		FreeInodesCount: int32(n),
		InodeSize:       int32(binary.Size(Inodo{})),
		InodeStart:      int64(pointerInodoStart),
		FirstInode:      0,

		BlocksCount:     int32(n * 3),
		FreeBlocksCount: int32(n * 3),
		BlockSize:       int32(binary.Size(BloqueCarpeta{})),
		BlockStart:      int64(pointerBlockStart),
		FirstBlock:      0,

		MountedTime:  getFechaByte(),
		MountedCount: 1,
		Magic:        0xEF53,

		BitMapInodeStart: int64(pointerBmi),
		BitMapBlockStart: int64(pointerBmb),
	}
}

/*JOURNALING*/

func escribirJournal(path string, jn Journaling, seek int64) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	// escribir la estructura
	file.Seek(seek, 0)
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, &jn)
	WriteNextBytes(file, binario.Bytes())
}

func recuperarJournal(path string, seek int64) (bool, Journaling) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// preparar estructura mbr
	sp := Journaling{}
	//ebr.Size = -1

	file.Seek(seek, 0)
	var tamano = binary.Size(sp)
	datos := ReadNextBytes(file, tamano)
	buffer := bytes.NewBuffer(datos)
	err = binary.Read(buffer, binary.BigEndian, &sp)
	if err != nil {
		fmt.Println(err)
		return false, sp
	}

	return true, sp
}

/*BITMAP*/
// INODOS
func escribirBitMap(path string, bitmap []byte, seek int64) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	// escribir la estructura
	file.Seek(seek, 0)
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, &bitmap)
	WriteNextBytes(file, binario.Bytes())
}

func recuperarBitMap(path string, seek int64, size int64) (bln bool, bitmap []byte) {
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	file.Seek(seek, 0)
	bitmap = ReadNextBytes(file, int(size))

	return false, bitmap
}

func recuperarInodo(path string, seek int64) (bool, Inodo) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// preparar estructura mbr
	sp := Inodo{}
	//ebr.Size = -1

	file.Seek(seek, 0)
	var tamano = binary.Size(sp)
	datos := ReadNextBytes(file, tamano)
	buffer := bytes.NewBuffer(datos)
	err = binary.Read(buffer, binary.BigEndian, &sp)
	if err != nil {
		fmt.Println(err)
		return false, sp
	}

	return true, sp
}

func crearInodo() Inodo {
	inodo := Inodo{}
	inodo.iniciarPunteros()
	return inodo
}

func escribirInodo(path string, inodo Inodo, seek int64) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	// escribir la estructura
	file.Seek(seek, 0)
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, &inodo)
	WriteNextBytes(file, binario.Bytes())
}

func (ii *Inodo) iniciarPunteros() {
	for i := 0; i < 15; i++ {
		ii.Block[i] = -1
	}
}

func (ii *SuperBlock) encontrarSiguienteInodoLibre(bitmap []byte) {
	for indice, bit := range bitmap {
		// recorrer todos los bit referentes a inodo
		// recuperar el bit
		if bit == 0 {
			// inodo libre
			ii.FirstInode = int32(indice)
			break
		}
	}
}
func (ii *SuperBlock) encontrarSiguienteBloqueLibre(bitmap []byte) {
	for indice, bit := range bitmap {
		// recorrer todos los bit referentes a inodo
		// recuperar el bit
		if bit == 0 {
			// inodo libre
			ii.FirstInode = int32(indice)
			break
		}
	}
}

func addInodo(path string, inodo Inodo, sp *SuperBlock) {
	// encontrar la primer posiciòn del bitmap vacìa
	if file, err := os.OpenFile(path, os.O_RDWR, 0777); err == nil {
		// archivo abierto
		// iniciar en seeek de datos
		file.Seek(sp.BitMapInodeStart, 0)
		bitmap := ReadNextBytes(file, int(sp.InodesCount))

		//auxInodoUsado := sp.FirstInode // guardar el index del inodo utilizado
		bitmap[sp.FirstInode] = 1 // guarar en el bitmap la posiciòn del inodo utilizado

		sp.encontrarSiguienteInodoLibre(bitmap) // encontrar el siguiente bit libre en el bitmap de inodos
		escribirBitMap(path, bitmap, sp.BitMapInodeStart)

		// crear el primer bloque para el primer apuntador
		if inodo.Type == 0 {
			// crear bloque de carpetas
			//bloque := crearBloqueCarpeta()
			file.Seek(sp.BitMapBlockStart, 0)
			bitmapBlock := ReadNextBytes(file, int(sp.BlocksCount))
			bitmapBlock[sp.BitMapBlockStart] = 1

			sp.encontrarSiguienteBloqueLibre(bitmapBlock)

			for i := 0; i < 15; i++ {
				if inodo.Block[i] == -1 {
					// encontramos un apuntador libre

					// es un apuntador directo

					// es un apuntador indirecto

					// es un apuntador indirecto doble

				}
			}

		} else {
			// crar bloque de archivos

		}
		file.Close()
	} else {
		fmt.Println(err)
	}

}

// bloques carpeta
// crear bloque carpeta
func crearBloqueCarpeta() BloqueCarpeta {
	b := BloqueCarpeta{}
	b.iniciarPunteros()
	return b
}

func escribirBloqueCarpeta(path string, bc BloqueCarpeta, seek int64) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	// escribir la estructura
	file.Seek(seek, 0)
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, &bc)
	WriteNextBytes(file, binario.Bytes())
}

func recuperarBloqueCarpeta(path string, seek int64) (bool, BloqueCarpeta) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// preparar estructura mbr
	sp := BloqueCarpeta{}
	//ebr.Size = -1

	file.Seek(seek, 0)
	var tamano = binary.Size(sp)
	datos := ReadNextBytes(file, tamano)
	buffer := bytes.NewBuffer(datos)
	err = binary.Read(buffer, binary.BigEndian, &sp)
	if err != nil {
		fmt.Println(err)
		return false, sp
	}

	return true, sp
}

func (ii *BloqueCarpeta) iniciarPunteros() {
	for i := 0; i < 4; i++ {
		ii.Content[i].PointerInode = -1
	}
}
func (ii *BloqueCarpeta) indicarPadreyActual(padre int32, actual int32) {
	// el actual es el primer apuntador
	ii.Content[0].Name[0] = '.'
	ii.Content[0].PointerInode = actual
	// el padre es el segundo apuntador
	ii.Content[1].Name[0] = '.'
	ii.Content[1].Name[1] = '.'
	ii.Content[1].PointerInode = padre
}

func recuperarBloqueArchivo(path string, seek int64) (bool, BloqueArchivo) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// preparar estructura mbr
	sp := BloqueArchivo{}
	//ebr.Size = -1

	file.Seek(seek, 0)
	var tamano = binary.Size(sp)
	datos := ReadNextBytes(file, tamano)
	buffer := bytes.NewBuffer(datos)
	err = binary.Read(buffer, binary.BigEndian, &sp)
	if err != nil {
		fmt.Println(err)
		return false, sp
	}

	return true, sp
}

// escribir bloque apuntadores
// crear
func crearBloqueApuntadores() {

}

// escribir bloque archivos
func crearBloqueArchivos() {

}

func escribirBloqueArchivo(path string, bc BloqueArchivo, seek int64) {
	// recuperar mbr
	// abrir archivo
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	// escribir la estructura
	file.Seek(seek, 0)
	var binario bytes.Buffer
	binary.Write(&binario, binary.BigEndian, &bc)
	WriteNextBytes(file, binario.Bytes())
}

/*leer y escribir en archivos binarios*/
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

func WriteCeros(path string, inicio int64, final int64) {
	// abrir archivo
	if file, err := os.OpenFile(path, os.O_RDWR, 0777); err == nil {
		// sì se pudo abrir el archivo
		file.Seek(inicio, 0)
		for i := inicio; i < final; i++ {
			// recorrer todas las posiciones
			file.Write([]byte{0})
		}
		file.Close()
	} else {
		// no se pudo abrir el archivo
		fmt.Println(err)
		fmt.Println("No se pudieron escribir los ceros en el archivo, " + path)
	}
}

/*fecha en formato requerido*/
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
	return fechab
}

/*problemas con nulos en array de bytes*/
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

type EspaciosL struct {
	Inicios []int
	Finales []int
	ebrs    []Ebr
}
type EspacioL struct {
	Inicio int
	Final  int
	ebr    Ebr
}

func getEspaciosLibres(mbr Mbr) Espacios {
	var espaciosVacios Espacios
	var iniciaLibre int = binary.Size(mbr) // inicia el espacio libre
	// luego de la escritura del mbr

	// recorrer las particiones activas
	var espaciosLlenos Espacios
	for i := 0; i <= 3; i++ {
		if mbr.Partitions[i].Status == 1 {
			espaciosLlenos.Inicios = append(espaciosLlenos.Inicios, int(mbr.Partitions[i].Start))
			espaciosLlenos.Finales = append(espaciosLlenos.Finales, int(mbr.Partitions[i].Start+mbr.Partitions[i].Size))

		}
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

func fusionarEspaciosVacios(espacios EspaciosL) EspaciosL {
	var auxEspacios EspaciosL
	var espacioActual EspacioL
	for i := 0; i < len(espacios.Inicios)-1; i++ {
		if espacios.Finales[i] == espacios.Inicios[i+1] {
			if espacioActual.Inicio == 0 { // es la primera vez que se juntan espacios
				espacioActual.Inicio = espacios.Inicios[i]
				espacioActual.ebr = espacios.ebrs[i]
			}
			espacioActual.Final = espacios.Finales[i]

		} else {
			if espacioActual.Inicio != 0 {
				// guardar espacios finales actual
				auxEspacios.Inicios = append(auxEspacios.Inicios, int(espacioActual.Inicio))
				auxEspacios.Finales = append(auxEspacios.Finales, int(espacioActual.Final))
				auxEspacios.ebrs = append(auxEspacios.ebrs, espacioActual.ebr)

				espacioActual.Inicio = 0
				espacioActual.Final = 0
				espacioActual.ebr = Ebr{}
			} else {
				// guardar el espacio en la posiciòn I
				auxEspacios.Inicios = append(auxEspacios.Inicios, int(espacios.Inicios[i]))
				auxEspacios.Finales = append(auxEspacios.Finales, int(espacios.Finales[i]))
				auxEspacios.ebrs = append(auxEspacios.ebrs, espacios.ebrs[i])

				espacioActual.Inicio = 0
				espacioActual.Final = 0
				espacioActual.ebr = Ebr{}
			}
		}
	}

	if espacioActual.Inicio != 0 {
		// guardar espacios finales actual
		auxEspacios.Inicios = append(auxEspacios.Inicios, int(espacioActual.Inicio))
		auxEspacios.Finales = append(auxEspacios.Finales, int(espacioActual.Final))
		auxEspacios.ebrs = append(auxEspacios.ebrs, espacioActual.ebr)
	} else {
		// guardar el espacio en la posiciòn I
		i := len(espacios.Inicios) - 1
		auxEspacios.Inicios = append(auxEspacios.Inicios, int(espacios.Inicios[i]))
		auxEspacios.Finales = append(auxEspacios.Finales, int(espacios.Finales[i]))
		auxEspacios.ebrs = append(auxEspacios.ebrs, espacios.ebrs[i])
	}
	return auxEspacios
}

/*CONTROL DE ARCHIVOS*/
func getContenidoArchivo(inodo Inodo, particion ParticionMontada) (contenido string) {
	// recorrer los apuntadores del inodo
	for indice, apuntador := range inodo.Block {
		switch {
		case indice <= 12 && apuntador != -1:
			// apuntadores directos
			// OBTENER EL BLOQUE ARCHIVO
			_, bloqueArchivo := recuperarBloqueArchivo(particion.path, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))
			contenido2 := BytesToString(bloqueArchivo.Content[:])
			contenido2 = strings.ReplaceAll(contenido2, "\n", "\\n")
			break
		case indice == 13 && apuntador != -1:
			// apuntador indirecto
			break
		case indice == 14 && apuntador != -1:
			// apuntador indirecto doble
			break
		}
	}
	return contenido
}
