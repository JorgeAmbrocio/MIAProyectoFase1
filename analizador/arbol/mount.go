package arbol

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
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

type ParticionMontada struct {
	id     string
	letra  byte
	numero int
	path   string
	Status byte
	Type   byte
	Fit    byte
	Start  int64
	Size   int64
	End    int64
	Next   int64
	Name   [16]byte
	sp     SuperBlock
}

//var particionesMontadas []Montada
var particionesMontadas []ParticionMontada

func (i *mount) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
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

		if i.path == particion.path && auxNameParticion == particion.Name {
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
		if auxNameParticion == particionn.Name && particionn.Status == 1 {
			// es la particiòn que buscamos
			/*montada := Montada{
				letra:     auxLetra,
				numero:    auxNumero,
				path:      i.path,
				particion: particionn,
			}*/

			pmontada := ParticionMontada{
				letra:  auxLetra,
				numero: auxNumero,
				path:   i.path,
				Status: particionn.Status,
				Type:   particionn.Type,
				Fit:    particionn.Fit,
				Start:  particionn.Start,
				Size:   particionn.Size,
				Name:   particionn.Name,
			}

			//montada.id = "vd" + string(montada.letra) + strconv.Itoa(montada.numero)
			pmontada.id = "vd" + string(pmontada.letra) + strconv.Itoa(pmontada.numero)

			//particionesMontadas = append(particionesMontadas, montada)
			particionesMontadas = append(particionesMontadas, pmontada)

			fmt.Println("Particiòn montada con èxito :D")
			fmt.Println("\t" + pmontada.id)
			fmt.Println("\t" + i.path)
			fmt.Println("\t" + i.name)
			//part := particionesMontadas
			//fmt.Println(part)
			return
		}
	}

	// si llegamos hasta acà significa que no encontramos la particiòn primaria o extendida
	fmt.Println("No se ha encontrado la particiòn primaria o extendida, se buscarà entre las particiones lògicas")
	i.montarParticionL(auxLetra, auxNumero)
}

func (i *mount) montarParticionL(auxLetra byte, auxNumero int) {
	// recuperar el mbr
	auxMbr := RecuperarMBR(i.path)

	// encontrar la particiòn con el nombre a eliminar
	var nombreByte [16]byte
	copy(nombreByte[:], i.name)
	for indice := 0; indice <= 3; indice++ {
		auxParticion := auxMbr.Partitions[indice]
		// verificar que la particiòn sea de tipo extendida y que se encuentre activa
		if auxParticion.Status == 1 && auxParticion.Type == 'e' {
			// la particiòn sì es la adecuada

			posActual := auxParticion.Start
			//ebrAtnerior := Ebr{}
			for {

				exito, auxEbr := RecuperarEBR(i.path, posActual)

				if !exito {
					// sì encontrò el ebr
					// verificar que el ebr sea el correcto para editar
					var auxNombre [16]byte
					copy(auxNombre[:], i.name)
					if auxEbr.Status == 1 && auxEbr.Size > 0 && auxNombre == auxEbr.Name {
						// el ebr a montar es el actual
						// es la particiòn que buscamos
						pmontada := ParticionMontada{
							letra:  auxLetra,
							numero: auxNumero,
							path:   i.path,
							Status: auxEbr.Status,
							Next:   auxEbr.Next,
							Fit:    auxEbr.Fit,
							Start:  auxEbr.Start + int64(binary.Size(Ebr{})),
							Size:   auxEbr.Size - int64(binary.Size(Ebr{})),
							Name:   auxEbr.Name,
						}

						pmontada.id = "vd" + string(pmontada.letra) + strconv.Itoa(pmontada.numero)

						particionesMontadas = append(particionesMontadas, pmontada)

						fmt.Println("Particiòn montada con èxito :D")
						fmt.Println("\t" + pmontada.id)
						fmt.Println("\t" + i.path)
						fmt.Println("\t" + i.name)
					}
					// indicar la posiciòn siguiente
					if auxEbr.Next == -1 {
						break
					} else {
						//ebrAtnerior = auxEbr
						posActual = auxEbr.Next
					}
				} else {
					break
				}
			}
			break
		}
	}
}

func RecuperarParticionMontada(id string) (bool, *ParticionMontada) {

	for indice, particion := range particionesMontadas {
		if particion.id == id {
			return true, &particionesMontadas[indice]
		}
	}

	return false, &ParticionMontada{}
}
