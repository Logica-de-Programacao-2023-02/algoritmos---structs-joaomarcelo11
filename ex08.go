package main

import (
	"fmt"
	"math"
	"time"
)

type Viagem struct {
	Origem  string
	Destino string
	Data    time.Time
	Preco   float64
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	var viagemMaisCara Viagem
	precoMaximo := math.Inf(-1)

	for _, viagem := range viagens {
		if viagem.Preco > precoMaximo {
			viagemMaisCara = viagem
			precoMaximo = viagem.Preco
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: time.Now(), Preco: 150.0},
		{Origem: "Nova York", Destino: "Los Angeles", Data: time.Now(), Preco: 300.0},
		{Origem: "Londres", Destino: "Paris", Data: time.Now(), Preco: 200.0},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)

	fmt.Println("Viagem mais cara:")
	fmt.Println("Origem:", viagemMaisCara.Origem)
	fmt.Println("Destino:", viagemMaisCara.Destino)
	fmt.Println("Data:", viagemMaisCara.Data)
	fmt.Println("Preço:", viagemMaisCara.Preco)
}