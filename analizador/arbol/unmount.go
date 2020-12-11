package arbol

import (
	"fmt"
	"strings"
)

type unmount struct {
	id string
}

func (i *unmount) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "id":
			i.id = p.Valor
			break
		}
	}
}

func (i *unmount) Validar() bool {
	retorno := true
	if i.id == "" {
		retorno = false
	}
	return retorno
}

func Eunmount(p []Parametro) {
	i := unmount{}
	i.MatchParametros(p)
	if i.Validar() {
		// ejecutar la funciòn
		//pp := &particionesMontadas
		//fmt.Println(pp)
		for indice, particion := range particionesMontadas {
			if particion.id == i.id {
				if len(particionesMontadas) == 1 {
					particionesMontadas = []Montada{}
				} else {
					particionesMontadas = append(particionesMontadas[:indice], particionesMontadas[indice+1:]...)
				}
				fmt.Println("Desmontado con èxito " + i.id + " " + BytesToString(particion.particion.Name[:]))
				return
			}
		}

		fmt.Println("No se ha encontrado la particiòn montada " + i.id)
	}
}
