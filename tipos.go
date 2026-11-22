package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// TipoDato represents a data type in the system.
type TipoDato int

const (
	TipoInt TipoDato = iota
	TipoFloat
	TipoBool
	TipoString
	TipoIdentificador
)

// Dato es la e
type Dato struct {
	Tipo     TipoDato
	Entero   int
	Decimal  float64
	Booleano bool
	Texto    string //sirve para los strings y las variables
}

func ConverToken(token string) (Dato, error) {
	//evaluar si es bool
	if token == "true" || token == "false" {
		val, _ := strconv.ParseBool(token)
		return Dato{Tipo: TipoBool, Booleano: val}, nil
	}

	//evaluar si es int
	if val, err := strconv.Atoi(token); err == nil {
		return Dato{Tipo: TipoInt, Entero: val}, nil
	}

	//evaluar si es float
	if val, err := strconv.ParseFloat(token, 64); err == nil {
		return Dato{Tipo: TipoFloat, Decimal: val}, nil
	}

	//evaluar si es string (entre "")
	if len(token) >= 2 && strings.HasPrefix(token, "\"") && strings.HasSuffix(token, "\"") {
		limpio := strings.Trim(token, "\"")
		return Dato{Tipo: TipoString, Texto: limpio}, nil
	}

	//evaluar si es nombre de variable (identificador)
	if esIdentificadorValido(token) {
		return Dato{Tipo: TipoIdentificador, Texto: token}, nil
	}

	//si nada es, devuelve error
	return Dato{}, fmt.Errorf("token invalido: %s", token)
}

func esIdentificadorValido(token string) bool {
	if token == "" {
		return false
	}

	for i, r := range token {
		if i == 0 {
			// Primer carácter: solo letra o guion bajo
			if !unicode.IsLetter(r) && r != '_' {
				return false
			}
		} else {
			// Resto: letra, guion bajo o dígito
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
				return false
			}
		}
	}

	return true
}
