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
			ii.FirstBlock = int32(indice)
			break
		}
	}
}

func (ii *SuperBlock) getNextIndiceBloque(particion ParticionMontada) (indice int32) {
	indice = ii.FirstBlock
	// encontrar nuevo bloque vacìo
	_, bitmap := recuperarBitMap(particion.path, particion.sp.BitMapBlockStart, int64(particion.sp.BlocksCount))
	bitmap[ii.FirstBlock] = 1 // indicar que èste bit serà utilizado
	ii.encontrarSiguienteBloqueLibre(bitmap)
	escribirBitMap(particion.path, bitmap, particion.sp.BitMapBlockStart)
	return indice
}

func (ii *SuperBlock) getNextIndiceInodo(particion ParticionMontada) (indice int32) {
	indice = ii.FirstInode
	// encontrar nuevo bloque vacìo
	_, bitmap := recuperarBitMap(particion.path, particion.sp.BitMapInodeStart, int64(particion.sp.InodesCount))
	bitmap[ii.FirstInode] = 1 // indicar que èste bit serà utilizado
	ii.encontrarSiguienteInodoLibre(bitmap)
	escribirBitMap(particion.path, bitmap, particion.sp.BitMapInodeStart)
	return indice
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
	if !tienePermiso(inodo, 'r') {
		fmt.Println("No tienes permisos")
		return ""
	}
	// recorrer los apuntadores del inodo
	for indice, apuntador := range inodo.Block {
		switch {
		case indice <= 12 && apuntador != -1:
			// apuntadores directos
			// OBTENER EL BLOQUE ARCHIVO
			_, bloqueArchivo := recuperarBloqueArchivo(particion.path, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))
			contenido2 := BytesToString(bloqueArchivo.Content[:])
			contenido += contenido2
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

func escribirContenidoArchivo(contenido string, indiceInodo int64, particion *ParticionMontada) {

	// recueprar el inodo
	_, inodo := recuperarInodo(particion.path, particion.sp.InodeStart+int64(particion.sp.InodeSize)*indiceInodo)

	if !tienePermiso(inodo, 'w') {
		fmt.Println("No tienes permisos")
		return
	}

	// recorrer todos los apuntadores para empezar a escribir
	indiceCaracterEscribir := 0
	for indice, apuntador := range inodo.Block {
		// verifica que aùn hay caracteres por escribir
		if indiceCaracterEscribir < len(contenido)-1 {
			switch {
			case indice <= 12:

				if apuntador != -1 {
					// apuntadores directos con contenido
					// reescribr el bloque
					// OBTENER EL BLOQUE ARCHIVO
					_, bloqueArchivo := recuperarBloqueArchivo(particion.path, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))
					bloqueArchivo.Content = [64]byte{} // limpiar el contenido anterior

					var cantidadCaracteres = len(contenido) //- indiceCaracterEscribir
					if cantidadCaracteres > indiceCaracterEscribir+63 {
						cantidadCaracteres = indiceCaracterEscribir + 63
					}
					copy(bloqueArchivo.Content[:], contenido[indiceCaracterEscribir:cantidadCaracteres])
					indiceCaracterEscribir += 63
					escribirBloqueArchivo(particion.path, bloqueArchivo, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))

				} else {
					// apuntadores directos vacìos
					// buscar un nuevo bloque en bitmap para añadir el contenido extra
					apuntador2 := particion.sp.FirstBlock
					inodo.Block[indice] = apuntador2
					// encontrar nuevo bloque vacìo
					_, bitmap := recuperarBitMap(particion.path, particion.sp.BitMapBlockStart, int64(particion.sp.BlocksCount))
					bitmap[apuntador2] = 1 // indicar que èste bit serà utilizado
					particion.sp.encontrarSiguienteBloqueLibre(bitmap)
					escribirBitMap(particion.path, bitmap, particion.sp.BitMapBlockStart)

					bloqueArchivo := BloqueArchivo{}        // crear nuevo bloque
					var cantidadCaracteres = len(contenido) //- indiceCaracterEscribir
					if cantidadCaracteres > indiceCaracterEscribir+63 {
						cantidadCaracteres = indiceCaracterEscribir + 63
					}
					copy(bloqueArchivo.Content[:], contenido[indiceCaracterEscribir:cantidadCaracteres])
					indiceCaracterEscribir += 63
					escribirBloqueArchivo(particion.path, bloqueArchivo, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador2))

				}
				break
			case indice == 13:
				// apuntador indirecto
				break
			case indice == 14:
				// apuntador indirecto doble
				break
			}
		} else {
			// ya no hay caracteres por escribir
			// indicar a todos los bloques como -1
			switch {
			case indice <= 12 && indice != -1:
				// apuntadores directos con contenido
				inodo.Block[indice] = -1
				break
			case indice == 13:
				// apuntador indirecto
				break
			case indice == 14:
				// apuntador indirecto doble
				break
			}
		}
	}

	// escribir los nuevos datos
	escribirInodo(particion.path, inodo, particion.sp.InodeStart+int64(particion.sp.InodeSize)*indiceInodo)
}

/*USUARIOS*/
func getUsuarioYGrupo(particion ParticionMontada) (map[string]UsuarioArchivo, map[string]int) {
	_, inodo := recuperarInodo(particion.path, particion.sp.InodeStart+1*int64(particion.sp.InodeSize))
	contenidoArchivo := getContenidoArchivo(inodo, particion)

	mapaUsuario := make(map[string]UsuarioArchivo)
	mapaGrupo := make(map[string]int)

	filas := strings.Split(contenidoArchivo, "\n")
	for _, fila := range filas {
		atributos := strings.Split(fila, ",")
		if len(atributos) == 5 {
			// usuario
			if uid, err := strconv.Atoi(atributos[0]); err == nil {
				usr := UsuarioArchivo{UID: int32(uid), grupo: atributos[2], nombre: atributos[3], contrasena: atributos[4]}
				mapaUsuario[usr.nombre] = usr
			}
		} else {
			// grupo
			if gid, err := strconv.Atoi(atributos[0]); err == nil {
				mapaGrupo[atributos[2]] = gid
			}
		}
	}
	return mapaUsuario, mapaGrupo
}

func getContenidoArchivoUsuarios(particion ParticionMontada) string {
	_, inodo := recuperarInodo(particion.path, particion.sp.InodeStart+1*int64(particion.sp.InodeSize))

	contenidoArchivo := getContenidoArchivo(inodo, particion)

	return contenidoArchivo
}

func getCarpetaFromInodo(nombreBuscar string, inodo Inodo, particion ParticionMontada) (int32, int, int32) {

	if !tienePermiso(inodo, 'r') {
		fmt.Println("No tienes permisos")
		return -1, -1, -1
	}

	// recorrer todos los apuntadores
	for i, apuntador := range inodo.Block {
		switch {
		case i <= 12 && apuntador != -1:
			// apuntadores directos conº contenido
			// obtener el bloque archivo
			_, bloqueCarpeta := recuperarBloqueCarpeta(particion.path, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))

			// recorrer los apuntadores de la carpeta
			for j, contenidoCarpeta := range bloqueCarpeta.Content {
				var auxNombreCarpeta [12]byte
				copy(auxNombreCarpeta[:], nombreBuscar)
				if auxNombreCarpeta == contenidoCarpeta.Name {
					// sì es el nombre de la carpeta
					// retornar el apuntador y la carpeta
					// indice apuntador del inodo,
					return apuntador, j, contenidoCarpeta.PointerInode
				}
			}
			break
		case i == 13:
			// apuntador indirecto

			break
		case i == 14:
			// apuntador indirecto doble
			break
		}
	}

	return -1, -1, -1
}

func crearArchivoEnInodo(indiceInodo int, inodo Inodo, particion *ParticionMontada, cantidadCaracteres int, nombre string) {
	// buscar un apuntador libre para crear el archivo

	if !tienePermiso(inodo, 'w') {
		fmt.Println("No tienes permisos")
		return
	}

	for indiceParaCarpeta, apuntador := range inodo.Block {
		if apuntador != -1 {
			// ya existe un bloque de carpeta
			// recorrer todos los apuntadores de carpeta
			_, bloqueCarpeta := recuperarBloqueCarpeta(particion.path, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))

			// recorrer los apuntadores de la carpeta
			for indiceCarpeta, contenidoCarpeta := range bloqueCarpeta.Content {
				if contenidoCarpeta.PointerInode == -1 {
					// carpeta libre, podemos crear aquì el archivo

					// encontrar inodo libre
					indiceInodoLibre := particion.sp.getNextIndiceInodo(*particion)

					// preparar el contenido de la carpeta
					var auxNombreCarpeta [12]byte
					copy(auxNombreCarpeta[:], nombre)
					// asignar valroes a la carpeta
					bloqueCarpeta.Content[indiceCarpeta].Name = auxNombreCarpeta
					bloqueCarpeta.Content[indiceCarpeta].PointerInode = indiceInodoLibre

					// crear el inodo
					inodoNuevo := crearInodo()
					inodoNuevo.iniciarPunteros()
					inodoNuevo.Type = 1 // indicar que es inodo de archivos
					inodoNuevo.Ctime = getFechaByte()
					inodoNuevo.GID = UsuarioActualLogueado.GUID
					inodoNuevo.UID = UsuarioActualLogueado.UID
					inodoNuevo.Perm = [3]int8{7, 7, 7}
					inodoNuevo.Size = int32(cantidadCaracteres)

					// escribir las estructuras utilizadas
					escribirBloqueCarpeta(particion.path, bloqueCarpeta, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))
					escribirInodo(particion.path, inodoNuevo, particion.sp.InodeStart+int64(particion.sp.InodeSize)*int64(indiceInodoLibre))

					// ya se ha creado el inodo
					// escribir el archivo
					contenido := ""
					for i := 0; i < cantidadCaracteres; i++ {
						contenido += "1"
					}

					escribirContenidoArchivo(contenido, int64(indiceInodoLibre), particion)
					//fmt.Println("Archivo creado con èxito :D, el programador està muy feliz")
					return
				}
			}
		} else {
			// no existe un bloque carpeta
			// crear un bloque nuevo de carpeta
			apuntadorCarpeta, bloqueCarpeta := crearCarpetaEnIndiceInodo(int64(indiceInodo), inodo, int32(indiceParaCarpeta), particion)
			// recorrer los apuntadores de la carpeta
			for indiceCarpeta, contenidoCarpeta := range bloqueCarpeta.Content {
				if contenidoCarpeta.PointerInode == -1 {
					// carpeta libre, podemos crear aquì el archivo

					// encontrar inodo libre
					indiceInodoLibre := particion.sp.getNextIndiceInodo(*particion)

					// preparar el contenido de la carpeta
					var auxNombreCarpeta [12]byte
					copy(auxNombreCarpeta[:], nombre)
					// asignar valroes a la carpeta
					bloqueCarpeta.Content[indiceCarpeta].Name = auxNombreCarpeta
					bloqueCarpeta.Content[indiceCarpeta].PointerInode = indiceInodoLibre

					// crear el inodo
					inodoNuevo := crearInodo()
					inodoNuevo.iniciarPunteros()
					inodoNuevo.Type = 1 // indicar que es inodo de archivos
					inodoNuevo.Ctime = getFechaByte()
					inodoNuevo.GID = UsuarioActualLogueado.GUID
					inodoNuevo.UID = UsuarioActualLogueado.UID
					inodoNuevo.Perm = [3]int8{7, 7, 7}
					inodoNuevo.Size = int32(cantidadCaracteres)

					// escribir las estructuras utilizadas
					escribirBloqueCarpeta(particion.path, bloqueCarpeta, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntadorCarpeta))
					escribirInodo(particion.path, inodoNuevo, particion.sp.InodeStart+int64(particion.sp.InodeSize)*int64(indiceInodoLibre))

					// ya se ha creado el inodo
					// escribir el archivo
					contenido := ""
					for i := 0; i < cantidadCaracteres; i++ {
						contenido += "1"
					}

					escribirContenidoArchivo(contenido, int64(indiceInodoLibre), particion)
					fmt.Println("Archivo creado con èxito :D, el programador està muy feliz")
					return
				}
			}
		}
	}
}

func crearCarpetaEnIndiceInodo(indiceInodo int64, inodo Inodo, indiceParaCarpeta int32, particion *ParticionMontada) (int32, BloqueCarpeta) {

	if !tienePermiso(inodo, 'w') {
		fmt.Println("Lo siento, bruh, tu mamà no te dio permiso")
		blq := BloqueCarpeta{}
		blq.iniciarPunteros()
		return -1, blq
	}

	// obtener el bit en el que se puede crear la carpeta
	indiceCarpeta := particion.sp.getNextIndiceBloque(*particion)

	// apuntar el inodo al ìndice de carpeta futuro
	inodo.Block[indiceParaCarpeta] = indiceCarpeta

	// crear el objeto bloque carpeta
	bloqueCarpeta := BloqueCarpeta{}
	bloqueCarpeta.iniciarPunteros()
	//bloqueCarpeta.indicarPadreyActual(0, indiceCarpeta)

	// escribir el bloque con punteros vacìos
	escribirBloqueCarpeta(particion.path, bloqueCarpeta, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(indiceCarpeta))
	escribirInodo(particion.path, inodo, particion.sp.InodeStart+int64(particion.sp.InodeSize)*indiceInodo)

	return indiceCarpeta, bloqueCarpeta
}

func crearCarpetaEnInodo(indiceInodo int64, inodo Inodo, particion *ParticionMontada, nombreCarpeta string) (retorno int32) {

	if !tienePermiso(inodo, 'w') {
		fmt.Println("No tienes permisos")
		return -1
	}

	for indiceParaCarpeta, apuntador := range inodo.Block {
		if apuntador != -1 {
			// ya existe un bloque de carpeta
			// recorrer todos los apuntadores de carpeta
			_, bloqueCarpeta := recuperarBloqueCarpeta(particion.path, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))

			// recorrer los apuntadores de la carpeta
			for indiceCarpeta, contenidoCarpeta := range bloqueCarpeta.Content {
				if contenidoCarpeta.PointerInode == -1 {
					// carpeta libre, podemos crear aquì la carpeta nueva

					// encontrar inodo libre
					indiceInodoLibre := particion.sp.getNextIndiceInodo(*particion)
					retorno = indiceInodoLibre
					// preparar el contenido de la carpeta
					var auxNombreCarpeta [12]byte
					copy(auxNombreCarpeta[:], nombreCarpeta)
					// asignar valroes a la carpeta
					bloqueCarpeta.Content[indiceCarpeta].Name = auxNombreCarpeta
					bloqueCarpeta.Content[indiceCarpeta].PointerInode = indiceInodoLibre

					// crear el inodo
					inodoNuevo := crearInodo()
					inodoNuevo.iniciarPunteros()
					inodoNuevo.Type = 0 // indicar que es inodo de archivos
					inodoNuevo.Ctime = getFechaByte()
					inodoNuevo.GID = UsuarioActualLogueado.GUID
					inodoNuevo.UID = UsuarioActualLogueado.UID
					inodoNuevo.Perm = [3]int8{7, 7, 7}
					inodoNuevo.Size = int32(0)

					// apuntar al primer bloque de carpeta
					// obtener el bit en el que se puede crear la carpeta
					indiceCarpeta2 := particion.sp.getNextIndiceBloque(*particion)

					// apuntar el inodo al ìndice de carpeta futuro
					inodoNuevo.Block[0] = indiceCarpeta2

					// crear el objeto bloque carpeta
					bloqueCarpeta2 := BloqueCarpeta{}
					bloqueCarpeta2.iniciarPunteros()
					bloqueCarpeta2.indicarPadreyActual(int32(indiceInodo), indiceInodoLibre)

					// escribir las estructuras utilizadas
					escribirBloqueCarpeta(particion.path, bloqueCarpeta2, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(indiceCarpeta2))
					escribirBloqueCarpeta(particion.path, bloqueCarpeta, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))
					escribirInodo(particion.path, inodoNuevo, particion.sp.InodeStart+int64(particion.sp.InodeSize)*int64(indiceInodoLibre))

					//escribirContenidoArchivo(contenido, int64(indiceInodoLibre), particion)
					fmt.Println("Carpeta creada con èxito :D, el programador està muy feliz")
					return retorno
				}
			}
		} else {
			// no existe un bloque carpeta
			// crear un bloque nuevo de carpeta
			apuntadorCarpeta, bloqueCarpeta := crearCarpetaEnIndiceInodo(int64(indiceInodo), inodo, int32(indiceParaCarpeta), particion)
			// recorrer los apuntadores de la carpeta
			for indiceCarpeta, contenidoCarpeta := range bloqueCarpeta.Content {
				if contenidoCarpeta.PointerInode == -1 {
					// carpeta libre, podemos crear aquì el archivo

					// encontrar inodo libre
					indiceInodoLibre := particion.sp.getNextIndiceInodo(*particion)

					// preparar el contenido de la carpeta
					var auxNombreCarpeta [12]byte
					copy(auxNombreCarpeta[:], nombreCarpeta)
					// asignar valroes a la carpeta
					bloqueCarpeta.Content[indiceCarpeta].Name = auxNombreCarpeta
					bloqueCarpeta.Content[indiceCarpeta].PointerInode = indiceInodoLibre

					// crear el inodo
					inodoNuevo := crearInodo()
					inodoNuevo.iniciarPunteros()
					inodoNuevo.Type = 0 // indicar que es inodo de archivos
					inodoNuevo.Ctime = getFechaByte()
					inodoNuevo.GID = UsuarioActualLogueado.GUID
					inodoNuevo.UID = UsuarioActualLogueado.UID
					inodoNuevo.Perm = [3]int8{7, 7, 7}
					inodoNuevo.Size = int32(0)

					// apuntar al primer bloque de carpeta
					// obtener el bit en el que se puede crear la carpeta
					indiceCarpeta2 := particion.sp.getNextIndiceBloque(*particion)

					// apuntar el inodo al ìndice de carpeta futuro
					inodoNuevo.Block[0] = indiceCarpeta2

					// crear el objeto bloque carpeta
					bloqueCarpeta2 := BloqueCarpeta{}
					bloqueCarpeta2.iniciarPunteros()
					bloqueCarpeta2.indicarPadreyActual(int32(indiceInodo), indiceInodoLibre)

					// escribir las estructuras utilizadas
					escribirBloqueCarpeta(particion.path, bloqueCarpeta2, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(indiceCarpeta2))
					escribirBloqueCarpeta(particion.path, bloqueCarpeta, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntadorCarpeta))
					escribirInodo(particion.path, inodoNuevo, particion.sp.InodeStart+int64(particion.sp.InodeSize)*int64(indiceInodoLibre))

					fmt.Println("Carpeta creada con èxito :D, el programador està muy feliz")
					return indiceInodoLibre
				}
			}
		}
	}

	return -1
}

func tienePermiso(inodo Inodo, accion byte) (retorno bool) {
	// acciòn; R -> read; W -> write; X -> ejecutar; P -> reName, chmod;
	retorno = false

	// el usuario logueado es el root ?
	if UsuarioActualLogueado.UID == 1 {
		return true
	}

	// obtener tipo usuario
	var tipoUsuario int = 2
	switch {
	case inodo.UID == UsuarioActualLogueado.UID:
		tipoUsuario = 0
		break
	case inodo.GID == UsuarioActualLogueado.GUID:
		tipoUsuario = 1
		break
	default:
		tipoUsuario = 2
	}

	// verificar el permiso
	switch inodo.Perm[tipoUsuario] {
	case 7: // 111
		retorno = true
		break
	case 6: // 110
		retorno = accion == 'r' || accion == 'w'
		break
	case 5: // 101
		retorno = accion == 'r' || accion == 'x'
		break
	case 4: // 100
		retorno = accion == 'r'
		break
	case 3: // 011
		retorno = accion == 'w' || accion == 'x'
		break
	case 2: // 010
		retorno = accion == 'w'
		break
	case 1: // 001
		retorno = accion == 'x'
		break
	case 0: // 000
		retorno = false
		break
	}

	return retorno
}

func removerDeCarpeta(indiceCarpeta int32, carpeta BloqueCarpeta, indiceInodo int32, inodo Inodo, particion *ParticionMontada) {
	fmt.Println("no sirve")
}

func remove(iCarpeta int32, carpeta BloqueCarpeta, iEliminar int32, particion *ParticionMontada) {

	if removeRecursivo(carpeta.Content[iEliminar].PointerInode, particion) {
		// limpiar el bloque carpeta
		carpeta.Content[iEliminar].Name = [12]byte{}
		carpeta.Content[iEliminar].PointerInode = -1

		// escribir la carpeta modificada
		escribirBloqueCarpeta(particion.path, carpeta, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(iCarpeta))
	}

}

func removeRecursivo(punteroInodoEliminar int32, particion *ParticionMontada) (retorno bool) {
	retorno = false
	_, inodo := recuperarInodo(particion.path, particion.sp.InodeStart+int64(particion.sp.InodeSize)*int64(punteroInodoEliminar))

	// verificar permisos para eliminar
	if !tienePermiso(inodo, 'w') {
		return false
	}

	if punteroInodoEliminar == 0 {
		fmt.Println("No se puede eliminar la raìz")
		return false
	}

	// identificar tipo de inodo
	//_, bitMapInodo := recuperarBitMap(particion.path, particion.sp.BitMapInodeStart, int64(particion.sp.InodesCount))
	//_, bitMapBlock := recuperarBitMap(particion.path, particion.sp.BitMapBlockStart, int64(particion.sp.BlocksCount))
	switch inodo.Type {
	case 0:
		// es un inodo de carpetas
		// recorrre bloqeu carpeta
		var auxSiEliminar bool = true
		for indice, apuntador := range inodo.Block {
			switch {
			case indice < 13 && apuntador != -1:
				// recuperar blqoue carpeta
				_, bloqueCarpeta := recuperarBloqueCarpeta(particion.path, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))

				var contadorEliminados int = 0
				for indice2, apuntadorCarpeta := range bloqueCarpeta.Content {
					if apuntadorCarpeta.PointerInode != -1 {
						// recorrer todos los punteros del bloque carpeta
						if apuntadorCarpeta.Name[0] != '.' { // no es un apuntador a sì mismo ni padre
							if removeRecursivo(apuntadorCarpeta.PointerInode, particion) {
								// editar y guardar el bitmap
								_, bitMapInodo := recuperarBitMap(particion.path, particion.sp.BitMapInodeStart, int64(particion.sp.InodesCount))
								bitMapInodo[apuntadorCarpeta.PointerInode] = 0
								// escribir bitmap de bloques
								escribirBitMap(particion.path, bitMapInodo, particion.sp.BitMapInodeStart)

								// limpiar el bloque carpeta
								bloqueCarpeta.Content[indice2].PointerInode = -1
								bloqueCarpeta.Content[indice2].Name = [12]byte{}
								particion.sp.FreeBlocksCount++

								contadorEliminados++
							} else {
								fmt.Println("No se pudo eliminar el archivo ")
								auxSiEliminar = false
							}
						} else {
							contadorEliminados++
						}
					} else {
						contadorEliminados++
					}
				}
				// escribir blqque carpeta editado
				escribirBloqueCarpeta(particion.path, bloqueCarpeta, particion.sp.BlockStart+int64(particion.sp.BlockSize)*int64(apuntador))

				if contadorEliminados == 4 {
					// se eliminaron los 4 punteros
					// eliminar el bloque carpeta

					// eliminar el puntero del inodo hacia el bloque carpeta
					inodo.Block[indice] = -1

					// editar y guardar el bitmap
					_, bitMapBlock := recuperarBitMap(particion.path, particion.sp.BitMapBlockStart, int64(particion.sp.BlocksCount))

					bitMapBlock[apuntador] = 0
					escribirBitMap(particion.path, bitMapBlock, particion.sp.BitMapBlockStart)
				}
				break
			case indice == 13 && apuntador != -1:

				break
			case indice == 14 && apuntador != -1:

				break
			}
		}

		_, bitMapInodo := recuperarBitMap(particion.path, particion.sp.BitMapInodeStart, int64(particion.sp.InodesCount))
		_, bitMapBlock := recuperarBitMap(particion.path, particion.sp.BitMapBlockStart, int64(particion.sp.BlocksCount))

		// escribir el inodo de nuevo porque lo modificamos
		escribirInodo(particion.path, inodo, particion.sp.InodeStart+int64(particion.sp.InodeSize)*int64(punteroInodoEliminar))

		if auxSiEliminar {
			// eliminar el inodo actual
			particion.sp.FreeInodesCount++
			bitMapInodo[punteroInodoEliminar] = 0

			// limpiar en el bitmap de bloques el primer bloque porque nunca se eliminan los punteros Apadre y Actual
			//bitMapBlock[inodo.Block[0]] = 0
		}

		// escribir bitmap de bloques
		escribirBitMap(particion.path, bitMapBlock, particion.sp.BitMapBlockStart)
		escribirBitMap(particion.path, bitMapInodo, particion.sp.BitMapInodeStart)
		retorno = true
		break
	case 1:
		// es un inodo de archivo

		_, bitMapInodo := recuperarBitMap(particion.path, particion.sp.BitMapInodeStart, int64(particion.sp.InodesCount))
		_, bitMapBlock := recuperarBitMap(particion.path, particion.sp.BitMapBlockStart, int64(particion.sp.BlocksCount))
		for indice, apuntador := range inodo.Block {
			switch {
			case indice < 13 && apuntador != -1:
				// recuperar blqoue archivo
				bitMapBlock[apuntador] = 0
				particion.sp.FreeBlocksCount++
				break
			case indice == 13 && apuntador != -1:

				break
			case indice == 14 && apuntador != -1:

				break
			}
		}
		// eliminar el inodo actual
		bitMapInodo[punteroInodoEliminar] = 0

		escribirBitMap(particion.path, bitMapBlock, particion.sp.BitMapBlockStart)
		escribirBitMap(particion.path, bitMapInodo, particion.sp.BitMapInodeStart)
		retorno = true
		break
	}

	return retorno
}

func encontrarArchivo(path string, particion ParticionMontada) (inodo Inodo, pointerInodo int32, pointerCarpeta int32, indiceCarpeta int32) {
	// recuperar el primer inodo
	_, inodo = recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart)
	// recorrer todas las carpetas
	pathSplit := strings.Split(path, "/")
	for indice, carpeta := range pathSplit {
		if indice != 0 && indice != (len(pathSplit)) {
			// recuperar carpeta del inodo
			apuntadorCarpeta, indiceApuntadorCarpeta, iInodoCarpetaSiguiente := getCarpetaFromInodo(carpeta, inodo, *UsuarioActualLogueado.particion)
			pointerCarpeta = apuntadorCarpeta
			indiceCarpeta = int32(indiceApuntadorCarpeta)
			pointerInodo = iInodoCarpetaSiguiente
			if iInodoCarpetaSiguiente != -1 {
				// la carpeta existe
				_, inodo = recuperarInodo(UsuarioActualLogueado.particion.path, UsuarioActualLogueado.particion.sp.InodeStart+int64(UsuarioActualLogueado.particion.sp.InodeSize)*int64(iInodoCarpetaSiguiente))
				/*if !tienePermiso(inodo, 'r') {
					fmt.Println("No tienes permiso para leer èsta carpeta " + carpeta)
					return Inodo{}, -1, -1, -1
				}*/
			} else {
				// la carpeta no existe
				// mostrar error porque la ruta no es vàlida
				fmt.Println("\tNo se ha encontrado el directorio " + carpeta)
				return Inodo{}, -1, -1, -1
			}
		}
	}
	return inodo, pointerInodo, pointerCarpeta, indiceCarpeta
}
