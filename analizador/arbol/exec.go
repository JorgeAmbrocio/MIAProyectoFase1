package arbol

import (
	"fmt"
	"strings"
)

type execs struct {
	path string
}

func (i *execs) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "path":
			i.path = QuitarComillas(p.Valor)
			break
		}
	}

}

func (i *execs) Validar() bool {
	retorno := true

	if i.path == "" {
		retorno = false
	}

	return retorno
}

func Eexec(p []Parametro) {
	fmt.Println("Se està ejecutando el excect")
	i := execs{}
	i.MatchParametros(p)
	if i.Validar() {
		i.ejecutar()
	}
}

func (i *execs) ejecutar() {

}
