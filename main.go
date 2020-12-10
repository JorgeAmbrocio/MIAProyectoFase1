package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hola mundo")

	//cadena := "mkdisk -size->5 -path->/home/disco/archivo.dsk \n"
	//cadena += "exec -path->/ruta/ \n"
	//cadena += "pause\n"
	//cadena += "rmdisk -path->/home/folder/algo/archivo.dsk \n"
	//cadena := "fdisk -path->/home/folder/algo/archivo.dsk -name->particion2 -size->10 -type->e -unit->k -fit->ff \n"
	//cadena := "fdisk -path->/home/folder/algo/archivo.dsk -name->particion2 -add->-6 -unit->k \n"
	//"/home/folder/algo/archivo.particion
	//cadena := "rep -name->mbr   -path->/home/algo/reporte.jpg -id->vda1 \n"
	//cadena := "mount -name->particion2   -path->/home/folder/algo/archivo.dsk \n"
	//cadena += "mount -name->particion1   -path->/home/folder/algo/archivo.dsk \n"
	//cadena += "rep -path->/home/folder/algo/reporte.jpg -name->disk -id->vda1 \n"
	cadena := "exec -path->/home/script.arch \n"

	reader := bufio.NewReader(strings.NewReader(cadena))
	yyParse(newLexer(reader))

	fmt.Println("Terminó el análisis")
	fmt.Println(lInstruccion)
	ast := lAST[len(lAST)-1]
	ast.EjecutarAST()
}

//os.Exit(yyParse(newLexer(bufio.NewReader(os.Stdin))))
//os.Exit(yyParse(newLexer(reader)))
//; lInstruccion = append(lInstruccion, oInstruccion)

/*
var ooInstruccion arbol.Instruccion
	ooInstruccion = arbol.Instruccion{Tipo: "mkdisk"}
	ooInstruccion.Parametros = append(ooInstruccion.Parametros, arbol.Parametro{Tipo: "path", Valor: "/"})
	ooInstruccion.Parametros = append(ooInstruccion.Parametros, arbol.Parametro{Tipo: "size", Valor: "/algo"})
	ooInstruccion.Ejecutar()

	ooInstruccion = arbol.Instruccion{Tipo: "mkdisk"}
	ooInstruccion.Parametros = append(ooInstruccion.Parametros, arbol.Parametro{Tipo: "path", Valor: "/2"})
	ooInstruccion.Parametros = append(ooInstruccion.Parametros, arbol.Parametro{Tipo: "size", Valor: "/algo2"})
	ooInstruccion.Ejecutar()

	mkdisk -size->6 -unit->m -path->/home/folder/algo/ -name->disco1.dsk
mkdisk -size->7 -unit->m -path->/home/folder/algo/ -name->disco2.dsk

pause

rmdisk -path->/home/folder/algo/

*/
