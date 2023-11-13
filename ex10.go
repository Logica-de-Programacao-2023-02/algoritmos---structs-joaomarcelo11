package main

import (
	"fmt"
)

type Filme struct {
	titulo      string
	diretor     string
	ano         int
	avaliacoes  []int
}

func (f *Filme) adicionarAvaliacao(avaliacao int) {
	f.avaliacoes = append(f.avaliacoes, avaliacao)
}

func (f *Filme) removerAvaliacao(index int) {
	if index >= 0 && index < len(f.avaliacoes) {
		f.avaliacoes = append(f.avaliacoes[:index], f.avaliacoes[index+1:]...)
	}
}

func (f *Filme) calcularMediaAvaliacoes() float64 {
	total := 0
	for _, avaliacao := range f.avaliacoes {
		total += avaliacao
	}
	return float64(total) / float64(len(f.avaliacoes))
}

func (f *Filme) imprimirInformacoes() {
	fmt.Printf("Título: %s\n", f.titulo)
	fmt.Printf("Diretor: %s\n", f.diretor)
	fmt.Printf("Ano: %d\n", f.ano)
	fmt.Printf("Média de Avaliações: %.2f\n", f.calcularMediaAvaliacoes())
}

func main() {
	filme := Filme{
		titulo:     "Inception",
		diretor:    "Christopher Nolan",
		ano:        2010,
		avaliacoes: []int{8, 9, 7, 10, 9},
	}

	filme.adicionarAvaliacao(8)
	filme.removerAvaliacao(1)
	filme.imprimirInformacoes()
}