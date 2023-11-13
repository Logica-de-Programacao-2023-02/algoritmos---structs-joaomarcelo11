package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo   string
	Artista  string
	Duracao  time.Duration
}

type Playlist struct {
	Nome     string
	Musicas  []Musica
}

func imprimirPlaylist(p Playlist) {
	fmt.Printf("Nome da Playlist: %s\n", p.Nome)

	var duracaoTotal time.Duration
	for _, musica := range p.Musicas {
		fmt.Printf("Título: %s\n", musica.Titulo)
		fmt.Printf("Artista: %s\n", musica.Artista)
		fmt.Printf("Duração: %s\n", musica.Duracao)

		duracaoTotal += musica.Duracao
	}

	fmt.Printf("Tempo total da Playlist: %s\n", duracaoTotal)
}

func main() {
	musicas := []Musica{
		{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3 * time.Minute},
		{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4 * time.Minute},
		{Titulo: "Música 3", Artista: "Artista 3", Duracao: 5 * time.Minute},
	}

	playlist := Playlist{
		Nome:    "Minha Playlist",
		Musicas: musicas,
	}

	imprimirPlaylist(playlist)
}