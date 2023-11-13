package main

import (
	"fmt"
	"math"
)

type circulo struct {
	raio float64
}

func CalcularArea(circulo circulo) float64 {
	return math.Pi * circulo.raio * circulo.raio
}

func main() {
	var raio float64
	fmt.Println("Qual o seu raio? ")
	fmt.Scan(&raio)
	c := circulo{raio}
	area := CalcularArea(c)
	fmt.Println("Sua area é: ", area)
}