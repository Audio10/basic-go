package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaracion de variables enteras
	base := 12 // creacion y asignacion de una variable
	var altura int = 14
	var area int //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna valores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado

	const baseCuadrado int = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El Área del cuadrado es:", areaCuadrado)

}
