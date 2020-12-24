package main

import (
	"bufio"
	"fmt"
	"os"
	"proyectos/MIAProyectoFase1/analizador/arbol"
	"regexp"
	"strings"
)

func main() {

	for {
		fmt.Println("Escribe un comando")
		fmt.Print("-> ")

		text := bufio.NewReader(os.Stdin)
		texto, _ := text.ReadString('\n')
		//texto := "exec -path->/home/test3.arch \n"
		//yyParse(newLexer(bufio.NewReader(strings.NewReader(texto))))
		analizador(texto)

		//ast := lAST[len(lAST)-1]
		//ast.EjecutarAST()
		//break
	}
}

func Excec(cadena string) {
	if file, err := os.OpenFile(cadena, os.O_RDWR, 775); err == nil {
		buffer := bufio.NewReader(file)
		texto, _ := buffer.ReadString(0)
		analizador(texto)
	} else {
		fmt.Println(err)
	}

}

func analizador(contenido string) {
	/*
		var oParametro arbol.Parametro
		var oInstruccion arbol.Instruccion
		var lInstruccion []arbol.Instruccion
		var lAST []arbol.AST
		var auxPath string
	*/

	// quitar comentarios
	regular1 := regexp.MustCompile(`#[^\n]+`)
	contenido = regular1.ReplaceAllString(contenido, "")
	// saltos de linea doble
	regular2 := regexp.MustCompile(`\n(\n)*`)
	contenido = regular2.ReplaceAllString(contenido, "\n")
	// cambiar las asignaciones
	regular3 := regexp.MustCompile(`->`)
	contenido = regular3.ReplaceAllString(contenido, "=")
	// splitear contenido por fila
	lineas := strings.Split(contenido, "\n")

	var arbol arbol.AST = arbol.AST{}
	for _, linea := range lineas {
		// splitear por guiones para separar los parametros
		atributos := strings.Split(linea, "-")
		instruccion := SepararParametros(atributos)

		if instruccion.Tipo == "exec" {
			cadena := instruccion.Parametros[0].Valor
			Excec(cadena)
			continue
		}

		arbol.Instrucciones = append(arbol.Instrucciones, instruccion)
	}

	arbol.EjecutarAST()
}

func SepararParametros(parametros []string) (i arbol.Instruccion) {

	for indice, parametro := range parametros {
		parametro = strings.TrimSpace(parametro)
		if indice == 0 {
			// es el tipo de instrucciòn
			i.Tipo = strings.ToLower(parametro)
		} else {
			// es un paràmetro
			sparametro := strings.Split(parametro, "=")
			p := arbol.Parametro{}
			p.Tipo = strings.ToLower(strings.TrimSpace(sparametro[0]))
			if len(sparametro) == 2 {
				p.Valor = strings.TrimSpace(sparametro[1])
			} else {
				p.Valor = p.Tipo
			}
			i.Parametros = append(i.Parametros, p)
		}
	}
	return i
}
