package main

import "fmt"

type Triangulo struct {
	Base   float64
	Altura float64
}

func calcularAreaTriangulo(t Triangulo) float64 {
	return (t.Base * t.Altura) / 2
}

func main() {
	triangulo := Triangulo{
		Base:   10.5,
		Altura: 8.2,
	}

	area := calcularAreaTriangulo(triangulo)
	fmt.Printf("A área do triângulo é: %.2f\n", area)
}