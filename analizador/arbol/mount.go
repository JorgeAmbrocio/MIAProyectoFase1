package arbol

import (
	"fmt"
	"strconv"
)

type mount struct {
	name string
	path string
}

type Montada struct {
	id        string
	letra     byte
	numero    int
	path      string
	particion Partition
}

var particionesMontadas []Montada

func (i *mount) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch p.Tipo {
		case "path":
			i.path = QuitarComillas(p.Valor)
			break
		case "name":
			i.name = p.Valor
			break
		}
	}

}

func (i *mount) Validar() bool {
	retorno := true

	if i.path == "" || i.name == "" {
		retorno = false
	}

	return retorno
}

func Emount(p []Parametro) {
	fmt.Println("Soy el mount y me estoy ejecutando")

	i := mount{}
	i.MatchParametros(p)
	if i.Validar() {
		i.montarParticion()
	}
}

func (i *mount) montarParticion() {
	// letra y numero
	var auxLetra byte = 'a'
	var auxNumero int = 1

	// validar si la particiòn ya ha sido montada
	for _, particion := range particionesMontadas {
		var auxNameParticion [16]byte
		copy(auxNameParticion[:], i.name)

		if i.path == particion.path {
			auxNumero++
			
		} else {
			if auxLetra <= particion.letra {
				auxLetra = particion.letra + 1
			}
		}

		if i.path == particion.path && auxNameParticion == particion.particion.Name {
			// ya fue montada
			fmt.Println("La particiòn ya se encuentra montada :D")
			return
		}
	}

	// si llega hasta èste punto, significa que la particiòn no se ha montado
	// cargar particiòn
	// obtener el mbr
	auxMbr := RecuperarMBR(i.path)

	// encontrar la particiòn con el nombre
	var auxNameParticion [16]byte
	copy(auxNameParticion[:], i.name)

	for _, particionn := range auxMbr.Partitions {
		if auxNameParticion == particionn.Name {
			// es la particiòn que buscamos
			montada := Montada{
				letra:     auxLetra,
				numero:    auxNumero,
				path:      i.path,
				particion: particionn,
			}

			montada.id = "vd" + string(montada.letra) + strconv.Itoa(montada.numero)

			particionesMontadas = append(particionesMontadas, montada)
			part := particionesMontadas
			fmt.Println(part)
			break
		}
	}

}

func RecuperarParticionMontada(id string) (bool, Montada) {

	for _, particion := range particionesMontadas {
		if particion.id == id {
			return true, particion
		}
	}

	return false, Montada{}
}
