package ui

import (
	"fmt"
	"go-starter/model"
)

func PrintArtists(artists []*model.Artist) {
	fmt.Println("\n--- Artists ---")
	for i, artist := range artists {
		fmt.Printf("%d. %s\n", i+1, artist)
	}
}

func PrintAlbums(albums []*model.Album) {
	fmt.Println("\n--- Albums ---")
	for i, album := range albums {
		fmt.Printf("%d. %s\n", i+1, album)
	}
}

func PrintSongs(songs []*model.Song) {
	fmt.Println("\n--- Songs ---")
	for i, song := range songs {
		fmt.Printf("%d. %s\n", i+1, song)
	}
}
