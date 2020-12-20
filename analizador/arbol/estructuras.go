package arbol

import (
	"bufio"
	"fmt"
	"os"
)

// ESTRUCTURAS PARA EL FUNCIONAMIENTO DEL AST

// Instruccion unidad contenedora de datos a ejecutar
type Instruccion struct {
	Tipo       string
	Parametros []Parametro
}

// Ejecutar identifica el tipo de instrucción
// luego la manda a ejecutar según sea el caso
func (i Instruccion) Ejecutar() {
	//fmt.Println("Se ha ejecutado -> ", i.Tipo, "\n\t->", i.Parametros)

	switch i.Tipo {
	case "pause":
		fmt.Println("Presiona enter para continuar")
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		break
	case "exec":
		Eexec(i.Parametros)
		break
	case "mkdisk":
		Emkdisk(i.Parametros)
		break
	case "rmdisk":
		Ermdisk(i.Parametros)
		break
	case "fdisk":
		Efdisk(i.Parametros)
		break
	case "mount":
		Emount(i.Parametros)
		break
	case "unmount":
		Eunmount(i.Parametros)
		break
	case "rep":
		Erep(i.Parametros)
		break
	case "mkfs":
		Emkfs(i.Parametros)
		break
	case "login":
		Elogin(i.Parametros)
		break
	case "logout":
		Elogout(i.Parametros)
		break
	case "mkgrp":
		Emkgrp(i.Parametros)
		break
	case "mkusr":
		Emkusr(i.Parametros)
		break
	default:
		fmt.Println("No se reconoce la instrucción -> ", i.Tipo)
	}
}

// Parametro unidad que indica atributos de una intruccion
type Parametro struct {
	Tipo  string
	Valor string
}

// AST segmento de código para determinar la estructura de instruciones y sus funciones
type AST struct {
	Instrucciones []Instruccion
}

// EjecutarAST todas las instrucciones del árbol
func (ast AST) EjecutarAST() {
	fmt.Println("\n\n\tEjecutando AST")
	for _, instruccion := range ast.Instrucciones {
		instruccion.Ejecutar()
	}
}

// Mbr ...
type Mbr struct {
	Size       int32
	Date       [20]byte
	Signature  int32
	Partitions [4]Partition
	Fit        byte
}

// Partition ...
type Partition struct {
	Status byte // 0 vacìa, 1 ocupada, 2 eliminada
	Type   byte // P primaria, E extendida, L lógica
	Fit    byte // W peor ajuste, B mejor ajuste, P primer ajuste
	Start  int64
	Size   int64 // en tamaño bytes
	Name   [16]byte
}

// Ebr ...
type Ebr struct {
	Status byte // 0 vacìa, 1 ocupada, 2 eliminada
	Fit    byte // W peor ajuste, B mejor ajuste, P primer ajuste
	Start  int64
	Size   int64 // en tamaño bytes
	Next   int64
	Name   [16]byte
}

/*ESTRUCTURAS FASE DOS*/
type SuperBlock struct {
	Type             byte
	InodesCount      int32
	BlocksCount      int32
	FreeBlocksCount  int32
	FreeInodesCount  int32
	MountedTime      [20]byte
	UnMountedTime    [20]byte
	MountedCount     int32
	Magic            int32
	InodeSize        int32
	BlockSize        int32
	FirstInode       int32
	FirstBlock       int32
	BitMapInodeStart int64
	BitMapBlockStart int64
	InodeStart       int64
	BlockStart       int64
}

type Inodo struct {
	UID   int32
	GID   int32
	Size  int32
	Atmie [20]byte
	Ctime [20]byte
	Mtime [20]byte
	Block [15]int32
	Type  byte
	Perm  [3]int8 // bit 1: usuario, bit 2: grupo, bit 3 otros
}

type BloqueCarpeta struct {
	Content [4]Contenido
}

type Contenido struct {
	Name         [12]byte
	PointerInode int32
}

type BloqueArchivo struct {
	Content [64]byte
}

type BloqueApuntadores struct {
	Apuntadores [16]int32
}

type Journaling struct {
	TipoOperacion [10]byte
	Tipo          byte
	Nombre        [12]byte
	Contenid      [64]byte
	Fecha         [20]byte
	Propietario   [10]byte
	Permisos      int64
}

/*USUARIO*/
type Usuario struct {
	GUID      int32
	UID       int32
	particion *ParticionMontada
}

type UsuarioArchivo struct {
	UID        int32
	grupo      string
	nombre     string
	contrasena string
}

var UsuarioActualLogueado Usuario
