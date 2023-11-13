package main

import (
	"fmt"
	"time"
)

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) AumentarSalario(percentual float64) {
	aumento := f.Salario * (percentual / 100)
	f.Salario += aumento
}

func (f *Funcionario) DiminuirSalario(percentual float64) {
	reducao := f.Salario * (percentual / 100)
	f.Salario -= reducao
}

func (f *Funcionario) TempoServico() time.Duration {
	idadeInicioTrabalho := 18
	anosTrabalhados := f.Idade - idadeInicioTrabalho
	tempoServico := time.Duration(anosTrabalhados) * 365 * 24 * time.Hour
	return tempoServico
}

func main() {
	funcionario := Funcionario{
		Nome:    "João",
		Salario: 5000.0,
		Idade:   35,
	}

	fmt.Println("Salário inicial:", funcionario.Salario)

	funcionario.AumentarSalario(10)
	fmt.Println("Salário após aumento:", funcionario.Salario)

	funcionario.DiminuirSalario(5)
	fmt.Println("Salário após redução:", funcionario.Salario)

	tempoServico := funcionario.TempoServico()
	fmt.Println("Tempo de serviço:", tempoServico)
}