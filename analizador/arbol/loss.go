package arbol

import (
	"fmt"
	"strings"
)

type loss struct {
	id string
}

func (i *loss) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "id":
			i.id = p.Valor
			break
		}
	}
}

func (i *loss) Validar() bool {
	retorno := true

	if i.id == "" {
		retorno = false
	}

	return retorno
}

func Eloss(p []Parametro) {
	// escribir ceros para formatear el disco
	i := loss{}
	i.MatchParametros(p)
	if i.Validar() {
		fmt.Println("Loss")
		if exito, particion := RecuperarParticionMontada(i.id); exito {

			WriteCeros(particion.path, particion.sp.BitMapInodeStart, particion.Start+particion.Size)
			fmt.Println("\tEl sistema ha fallado con èxito")
		} else {
			fmt.Println("\tNo se ha encontrado la particiòn")
		}
	}
}
