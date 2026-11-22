package main

import "fmt"

type Entorno struct {
	ambitos []map[string]Dato
}

// 1. Crea el entorno inicializando el ámbito global
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
