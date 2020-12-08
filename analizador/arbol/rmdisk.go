package arbol

import (
	"log"
	"os"
)

type rmdisk struct {
	path string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *rmdisk) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch p.Tipo {
		case "path":
			i.path = QuitarComillas(p.Valor)
			break
		}
	}

}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i rmdisk) Validar() bool {
	retorno := true

	if i.path == "" {
		retorno = false
	}

	return retorno
}

// Ermdisk es la ejecución del rmdisk
func Ermdisk(p []Parametro) {

	i := rmdisk{}
	i.MatchParametros(p)
	if i.Validar() {
		// eliminar archivo
		err := os.Remove(i.path)
		if err != nil {
			log.Fatal(err)
		}
	}

}
