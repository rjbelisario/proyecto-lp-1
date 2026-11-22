# Calculadora RPN en Go

Requiere Go 1.22 o posterior. Desde la carpeta que contiene `go.mod`, ejecuta el paquete completo con:

```sh
go run .
```

`go run calculadora.go` compila únicamente el archivo de la consola. El punto en `go run .` incluye también `evaluador.go`, `tipos.go` y `entorno.go`.

Para compilar un ejecutable y luego iniciarlo:

```sh
go build -o calculadora.out .
./calculadora.out
```

## Sesión de ejemplo

El marcador `>` identifica lo que escribe el usuario.

```text
> x 5 2 + =
7
> x 5 + 2 *
24
> begin
> x 20 =
20
> vars
x = 20
> end
> vars
x = 7
> 5 true +
Error: error de tipos: + requiere dos números
> 10 2 -
8
> exit
```
