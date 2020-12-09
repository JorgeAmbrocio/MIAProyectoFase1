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
	switch i.Tipo {
	case "pause":
		fmt.Println("Se ha ejecutado -> ", i.Tipo, "\n\t->", i.Parametros)
		fmt.Println("Presiona enter para continuar")
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		break
	case "exec":
		fmt.Println("Se ha ejecutado -> ", i.Tipo, "\n\t->", i.Parametros)
		break
	case "mkdisk":
		fmt.Println("Se ha ejecutado -> ", i.Tipo, "\n\t->", i.Parametros)
		Emkdisk(i.Parametros)
		break
	case "rmdisk":
		fmt.Println("Se ha ejecutado -> ", i.Tipo, "\n\t->", i.Parametros)
		Ermdisk(i.Parametros)
		break
	case "fdisk":
		fmt.Println("Se ha ejecutado -> ", i.Tipo, "\n\t->", i.Parametros)
		Efdisk(i.Parametros)
		break
	case "moun":
		fmt.Println("Se ha ejecutado -> ", i.Tipo, "\n\t->", i.Parametros)
		break
	case "unmount":
		fmt.Println("Se ha ejecutado -> ", i.Tipo, "\n\t->", i.Parametros)
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
