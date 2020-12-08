package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hola mundo")

	cadena := "mkdisk -size->4 -path->/home/folder/algo/ -name->archivo.dsk \n"
	//cadena += "exec -path->/ruta/ \n"
	//cadena += "pause\n"
	//cadena += "rmdisk -path->/home/folder/algo/archivo.dsk \n"
	//cadena := "fdisk -path->/home/folder/algo/ -name->archivo.dsk -size->1 -type->p -unit->k \n"
	//"/home/folder/algo/archivo.dsk"

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


*/
