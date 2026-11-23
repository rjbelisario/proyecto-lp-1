package main

import "fmt"

type Entorno struct {
	ambitos []map[string]Dato
}

//Crea el entorno inicializando el ámbito global
func NuevoEntorno() *Entorno {
	return &Entorno{
		ambitos: []map[string]Dato{
			make(map[string]Dato), // Ámbito global
		},
	}
}

//Guarda o actualiza una variable en el ámbito más reciente (arriba de la pila)
func Definir(e *Entorno, nombre string, valor Dato) {
	actual := len(e.ambitos) - 1
	e.ambitos[actual][nombre] = valor
}

func Buscar(e *Entorno, nombre string) (Dato, bool) {
	for i := len(e.ambitos) - 1; i >= 0; i-- {
		if valor, existe := e.ambitos[i][nombre]; existe {
			return valor, true
		}
	}
	return Dato{}, false
}

func VariablesActivas(e *Entorno) map[string]Dato {
	variables := make(map[string]Dato)
	for _, ambito := range e.ambitos {
		for nombre, valor := range ambito {
			variables[nombre] = valor
		}
	}
	return variables
}

//Abre un nuevo bloque local (comando: begin)
func AbrirAmbito(e *Entorno) {
	e.ambitos = append(e.ambitos, make(map[string]Dato))
}

//Cierra el bloque local actual (comando: end)
func CerrarAmbito(e *Entorno) error {
	if len(e.ambitos) <= 1 {
		return fmt.Errorf("error: no hay ningún bloque 'begin' abierto para cerrar")
	}
	e.ambitos = e.ambitos[:len(e.ambitos)-1]
	return nil
}
