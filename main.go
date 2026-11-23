package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	entorno := NuevoEntorno()
	entrada := bufio.NewScanner(os.Stdin)
	info, err := os.Stdin.Stat()
	interactivo := err == nil && info.Mode()&os.ModeCharDevice != 0

	for {
		if interactivo {
			fmt.Print("> ")
		}
		if !entrada.Scan() {
			break
		}
		linea := strings.TrimSpace(entrada.Text())
		switch linea {
		case "":
			continue
		case "exit":
			return
		case "begin":
			AbrirAmbito(entorno)
		case "end":
			if err := CerrarAmbito(entorno); err != nil {
				fmt.Println("Error:", err)
			}
		case "vars":
			mostrarVariables(entorno)
		default:
			resultado, err := Evaluar(linea, &Pila{}, entorno)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(resultado)
			}
		}
	}
	if err := entrada.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error al leer la entrada:", err)
		os.Exit(1)
	}
}

func mostrarVariables(entorno *Entorno) {
	variables := VariablesActivas(entorno)
	nombres := make([]string, 0, len(variables))
	for nombre := range variables {
		nombres = append(nombres, nombre)
	}
	sort.Strings(nombres)
	for _, nombre := range nombres {
		fmt.Printf("%s = %s\n", nombre, variables[nombre])
	}
}
