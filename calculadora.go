package main

import "fmt"

type Pila struct {
	elementos []Dato
}

func Push(p *Pila, d Dato) {
	p.elementos = append(p.elementos, d)
}

func Pop(p *Pila) (Dato, error) {
	l := len(p.elementos)
	if l == 0 {
		return Dato{}, fmt.Errorf("error: pila insuficiente (stack underflow)")
	}
	ultimo := p.elementos[l-1]
	p.elementos = p.elementos[:l-1] // Recorta el último elemento
	return ultimo, nil
}

func Peek(p *Pila) (Dato, error) {
	l := len(p.elementos)
	if l == 0 {
		return Dato{}, fmt.Errorf("error: pila insuficiente (stack underflow)")
	}
	ultimo := p.elementos[l-1]
	return ultimo, nil
}
