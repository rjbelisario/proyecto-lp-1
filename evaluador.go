package main

import (
	"fmt"
	"strings"
)

func Evaluar(linea string, pila *Pila, entorno *Entorno) (Dato, error) {
	for _, texto := range strings.Fields(linea) {
		token, err := leerToken(texto)
		if err != nil {
			return Dato{}, err
		}
		dato := token.dato
		if token.operador != "" {
			dato, err = aplicarOperador(token.operador, pila, entorno)
			if err != nil {
				return Dato{}, err
			}
		}
		Push(pila, dato)
	}

	resultado, err := Peek(pila)
	if err != nil {
		return Dato{}, err
	}
	if len(pila.elementos) != 1 {
		return Dato{}, fmt.Errorf("la expresión dejó %d valores en la pila", len(pila.elementos))
	}
	return resolverDato(resultado, entorno)
}

func aplicarOperador(operador Operador, pila *Pila, entorno *Entorno) (Dato, error) {
	if len(pila.elementos) < 2 {
		return Dato{}, fmt.Errorf("error: pila insuficiente (stack underflow)")
	}
	derecho, _ := Pop(pila)
	izquierdo, _ := Pop(pila)

	derecho, err := resolverDato(derecho, entorno)
	if err != nil {
		return Dato{}, err
	}
	if operador == opAsignar {
		if izquierdo.Tipo != TipoIdentificador {
			return Dato{}, fmt.Errorf("error de tipos: = requiere un nombre de variable a la izquierda")
		}
		Definir(entorno, izquierdo.Texto, derecho)
		return derecho, nil
	}

	izquierdo, err = resolverDato(izquierdo, entorno)
	if err != nil {
		return Dato{}, err
	}
	if esNumero(izquierdo) && esNumero(derecho) {
		if izquierdo.Tipo == TipoInt && derecho.Tipo == TipoInt && operador != opDividir {
			return calcularEnteros(operador, izquierdo.Entero, derecho.Entero), nil
		}
		return calcularDecimales(operador, comoDecimal(izquierdo), comoDecimal(derecho))
	}
	if (operador == opIgual || operador == opDistinto) && izquierdo.Tipo == derecho.Tipo {
		iguales := izquierdo.Booleano == derecho.Booleano && izquierdo.Texto == derecho.Texto
		return Dato{Tipo: TipoBool, Booleano: iguales == (operador == opIgual)}, nil
	}
	return Dato{}, fmt.Errorf("error de tipos: %s requiere dos números", operador)
}

func resolverDato(dato Dato, entorno *Entorno) (Dato, error) {
	if dato.Tipo != TipoIdentificador {
		return dato, nil
	}
	valor, existe := Buscar(entorno, dato.Texto)
	if !existe {
		return Dato{}, fmt.Errorf("variable no definida: %s", dato.Texto)
	}
	return valor, nil
}

func esNumero(dato Dato) bool {
	return dato.Tipo == TipoInt || dato.Tipo == TipoFloat
}

func comoDecimal(dato Dato) float64 {
	if dato.Tipo == TipoInt {
		return float64(dato.Entero)
	}
	return dato.Decimal
}

func calcularEnteros(operador Operador, izquierdo, derecho int) Dato {
	switch operador {
	case opSumar:
		return Dato{Tipo: TipoInt, Entero: izquierdo + derecho}
	case opRestar:
		return Dato{Tipo: TipoInt, Entero: izquierdo - derecho}
	case opMultiplicar:
		return Dato{Tipo: TipoInt, Entero: izquierdo * derecho}
	default:
		return compararNumeros(operador, izquierdo, derecho)
	}
}

func calcularDecimales(operador Operador, izquierdo, derecho float64) (Dato, error) {
	switch operador {
	case opSumar:
		return Dato{Tipo: TipoFloat, Decimal: izquierdo + derecho}, nil
	case opRestar:
		return Dato{Tipo: TipoFloat, Decimal: izquierdo - derecho}, nil
	case opMultiplicar:
		return Dato{Tipo: TipoFloat, Decimal: izquierdo * derecho}, nil
	case opDividir:
		if derecho == 0 {
			return Dato{}, fmt.Errorf("división por cero")
		}
		return Dato{Tipo: TipoFloat, Decimal: izquierdo / derecho}, nil
	default:
		return compararNumeros(operador, izquierdo, derecho), nil
	}
}

func compararNumeros[T int | float64](operador Operador, izquierdo, derecho T) Dato {
	var resultado bool
	switch operador {
	case opIgual:
		resultado = izquierdo == derecho
	case opDistinto:
		resultado = izquierdo != derecho
	case opMenor:
		resultado = izquierdo < derecho
	case opMayor:
		resultado = izquierdo > derecho
	}
	return Dato{Tipo: TipoBool, Booleano: resultado}
}
