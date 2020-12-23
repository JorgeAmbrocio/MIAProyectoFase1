package arbol

import (
	"fmt"
	"strconv"
	"strings"
)

// mkdisk estructura de la instrucción mkdisk
type chgrp struct {
	usr   string
	pwd   string
	grupo string
}

// MatchParametros adjudica los parámetros en lista
// a los atributos del struct mkdisk
func (i *chgrp) MatchParametros(lp []Parametro) {
	for _, p := range lp {
		switch strings.ToLower(p.Tipo) {
		case "usr":
			i.usr = QuitarComillas(p.Valor)
			break
		case "grp":
			i.grupo = QuitarComillas(p.Valor)
		}
	}
}

// Validar : indica si el objeto cuenta con los parámetros suficientes para ejecutarse
func (i chgrp) Validar() bool {
	retorno := true

	if i.usr == "" || i.grupo == "" {
		retorno = false
	}

	return retorno
}

func Echgrp(p []Parametro) {
	i := chgrp{}
	i.MatchParametros(p)
	if i.Validar() {
		i.cambiarGrupo()
	}
}

func (i *chgrp) cambiarGrupo() {
	fmt.Println("CHGRP")
	if UsuarioActualLogueado.UID == 1 {
		particion := UsuarioActualLogueado.particion
		mapaUsuario, mapaGrupo := getUsuarioYGrupo(*particion)
		// verificar que el nuevo grupo exista
		if idGrupo := mapaGrupo[i.grupo]; idGrupo != 0 {
			// el grupo sì existe
			// crear nuevo string con todos los grupos y usuarios
			// recorrer todos lod grupos
			var contenidoGrupos string = ""
			var contenidoUsuarios string = ""
			for indice, grupo := range mapaGrupo {
				contenidoGrupos += strconv.Itoa(int(grupo)) + ",G," + indice + "\n"
			}

			for _, usuario := range mapaUsuario {
				if usuario.nombre == i.usr {
					contenidoUsuarios += strconv.Itoa(int(usuario.UID)) + ",U," + i.grupo + "," + usuario.nombre + "," + usuario.contrasena + "\n"
				} else {
					contenidoUsuarios += strconv.Itoa(int(usuario.UID)) + ",U," + usuario.grupo + "," + usuario.nombre + "," + usuario.contrasena + "\n"
				}

			}
			contenidoGrupos += contenidoUsuarios

			escribirContenidoArchivo(contenidoGrupos, 1, UsuarioActualLogueado.particion)
			guardarJournal(10, 0, i.usr, i.grupo, [3]int8{}, UsuarioActualLogueado.particion)
			fmt.Println("\tSe ha ejecutado el cambio de grupo con èxito")
		} else {
			fmt.Println("\tEl grupo que requieres no existe " + i.grupo)
		}
	} else {
		fmt.Println("\tNecesitar ser el usuario root para ejecutar èste comando")
	}
}
