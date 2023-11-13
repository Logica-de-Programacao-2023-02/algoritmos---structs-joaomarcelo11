package main

import (
	"fmt"
	"strings"
)

type Musica struct {
	Titulo   string
	Artista  string
	Duracao  int
}

type Playlist struct {
	Nome     string
	Musicas  []Musica
}

func buscarPlaylistsPorTitulo(tituloMusica string, playlists []Playlist) []Playlist {
	var playlistsEncontradas []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if strings.EqualFold(musica.Titulo, tituloMusica) {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}

	return playlistsEncontradas
}

func main() {
	musicas1 := []Musica{
		{Titulo: "Música 1", Artista: "Artista 1", Duracao: 180},
		{Titulo: "Música 2", Artista: "Artista 2", Duracao: 240},
		{Titulo: "Música 3", Artista: "Artista 3", Duracao: 300},
	}

	musicas2 := []Musica{
		{Titulo: "Música 4", Artista: "Artista 4", Duracao: 200},
		{Titulo: "Música 5", Artista: "Artista 5", Duracao: 250},
		{Titulo: "Música 6", Artista: "Artista 6", Duracao: 280},
	}

	playlist1 := Playlist{
		Nome:    "Playlist 1",
		Musicas: musicas1,
	}

	playlist2 := Playlist{
		Nome:    "Playlist 2",
		Musicas: musicas2,
	}

	playlists := []Playlist{playlist1, playlist2}

	titulo := "Música 1"
	resultado := buscarPlaylistsPorTitulo(titulo, playlists)

	if len(resultado) > 0 {
		fmt.Printf("Playlists com a música '%s':\n", titulo)
		for _, playlist := range resultado {
			fmt.Println(playlist.Nome)
		}
	} else {
		fmt.Printf("Nenhuma playlist encontrada com a música '%s'.\n", titulo)
	}
}