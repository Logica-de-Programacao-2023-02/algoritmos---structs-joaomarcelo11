package main

import "fmt"

type Endereco struct {
	Rua     string
	Numero  int
	Cidade  string
	Estado  string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func imprimirEnderecoCompleto(p Pessoa) {
	fmt.Printf("Endereço completo de %s:\n", p.Nome)
	fmt.Printf("Rua: %s, %d\n", p.Endereco.Rua, p.Endereco.Numero)
	fmt.Printf("Cidade: %s\n", p.Endereco.Cidade)
	fmt.Printf("Estado: %s\n", p.Endereco.Estado)
}

func main() {
	endereco := Endereco{
		Rua:    "Rua Exemplo",
		Numero: 123,
		Cidade: "Cidade Exemplo",
		Estado: "Estado Exemplo",
	}

	pessoa := Pessoa{
		Nome:     "João",
		Idade:    25,
		Endereco: endereco,
	}

	imprimirEnderecoCompleto(pessoa)
}