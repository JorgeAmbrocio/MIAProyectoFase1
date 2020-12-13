package main

import (
	"bufio"
	"fmt"
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
		texto := "exec -path->/home/ajustes.arch \n"
		yyParse(newLexer(bufio.NewReader(strings.NewReader(texto))))

		ast := lAST[len(lAST)-1]
		ast.EjecutarAST()
		break
	}
}

func prueba() {
	var a int = 10
	var b int = 3

	//var c float32 = float32(a) / float32(b)
	var algo string
	algo = fmt.Sprintf("%.2f", float32(a)/float32(b))
	fmt.Println(algo)

}
