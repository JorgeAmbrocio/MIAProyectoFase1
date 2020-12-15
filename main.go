package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"proyectos/MIAProyectoFase1/analizador/arbol"
	"strconv"
	"strings"
)

func main() {

	//prueba()
	//return

	for {
		fmt.Println("Escribe un comando")
		fmt.Print("-> ")

		//text := bufio.NewReader(os.Stdin)
		//texto, _ := text.ReadString('\n')
		texto := "exec -path->/home/fase2.arch \n"
		yyParse(newLexer(bufio.NewReader(strings.NewReader(texto))))

		ast := lAST[len(lAST)-1]
		ast.EjecutarAST()
		break
	}
}

func prueba() {
	fmt.Println(strconv.Itoa(binary.Size(arbol.BloqueCarpeta{})))
	fmt.Println(strconv.Itoa(binary.Size(arbol.BloqueApuntadores{})))
	fmt.Println(strconv.Itoa(binary.Size(arbol.BloqueArchivo{})))

}
