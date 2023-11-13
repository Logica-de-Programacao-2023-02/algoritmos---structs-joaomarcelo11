package main

import "fmt"

type Animal struct {
	Nome   string
	Especie string
	Idade  int
	Som    string
}

func (a *Animal) ModificarSom(novoSom string) {
	a.Som = novoSom
}

func (a Animal) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Espécie: %s\n", a.Especie)
	fmt.Printf("Idade: %d\n", a.Idade)
	fmt.Printf("Som: %s\n", a.Som)
}

func main() {
	animal := Animal{
		Nome:   "Max",
		Especie: "Cachorro",
		Idade:  3,
		Som:    "Latido",
	}

	animal.ImprimirInformacoes()

	fmt.Println("Modificando o som...")
	animal.ModificarSom("Au au")

	animal.ImprimirInformacoes()
}