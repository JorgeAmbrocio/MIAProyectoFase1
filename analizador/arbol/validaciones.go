package arbol

import (
	"log"
	"os"
	"strings"
)

// QuitarComillas quita las comillas iniciales y finales de un string
func QuitarComillas(str string) string {
	return strings.ReplaceAll(str, "\"", "")
}

/*CrearTodasCarpetas todas las carpteas pertinentes
 */
func CrearTodasCarpetas(ruta string) {
	if !ExisteCarpeta(ruta) {
		err := os.MkdirAll(ruta, 0770)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// ExisteCarpeta indica si una carpeta existe
func ExisteCarpeta(ruta string) bool {
	retorno := true
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		retorno = false
	}
	return retorno
}
