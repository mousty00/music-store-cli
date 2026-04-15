package ui

import (
	"bufio"
	"fmt"
	"go-starter/cart"
	"go-starter/helper"
	"go-starter/model"
	"go-starter/payment"
)

func ChooseArtist(artists []*model.Artist, scanner *bufio.Scanner) *model.Artist {
	for {
		PrintArtists(artists)
		fmt.Print("Choose an artist (0 to exit): ")
		choice, err := helper.ReadInt(scanner)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if choice == 0 {
			return nil
		}
		if choice < 1 || choice > len(artists) {
			fmt.Printf("Please enter a number between 1 and %d.\n", len(artists))
			continue
		}
		return artists[choice-1]
	}
}

func ChooseAlbum(albums []*model.Album, scanner *bufio.Scanner) *model.Album {
	if len(albums) == 0 {
		fmt.Println("This artist has no albums.")
		return nil
	}
	for {
		PrintAlbums(albums)
		fmt.Print("Choose an album (0 to exit): ")
		choice, err := helper.ReadInt(scanner)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if choice == 0 {
			return nil
		}
		if choice < 1 || choice > len(albums) {
			fmt.Printf("Please enter a number between 1 and %d.\n", len(albums))
			continue
		}
		return albums[choice-1]
	}
}

func ChooseSong(songs []*model.Song, scanner *bufio.Scanner) *model.Song {
	if len(songs) == 0 {
		fmt.Println("This album has no songs.")
		return nil
	}
	for {
		PrintSongs(songs)
		fmt.Print("Choose a song (0 to exit): ")
		choice, err := helper.ReadInt(scanner)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if choice == 0 {
			return nil
		}
		if choice < 1 || choice > len(songs) {
			fmt.Printf("Please enter a number between 1 and %d.\n", len(songs))
			continue
		}
		return songs[choice-1]
	}
}

func ShowCartAndActions(c *cart.Cart, scanner *bufio.Scanner) {
	c.Display()
	if c.IsEmpty() {
		return
	}
	fmt.Println("\nOptions:")
	fmt.Println("1. Remove an item")
	fmt.Println("2. Proceed to checkout")
	fmt.Println("0. Back to browsing")
	fmt.Print("Choose: ")

	choice, err := helper.ReadInt(scanner)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}
	switch choice {
	case 1:
		fmt.Print("Enter item number to remove: ")
		idx, err := helper.ReadInt(scanner)
		if err != nil || idx < 1 || idx > len(c.Items) {
			fmt.Println("Invalid item number.")
			return
		}
		c.RemoveItem(idx - 1)
	case 2:
		payment.Checkout(c, scanner)
	default:
		return
	}
}

func ShowActionMenu(selectedSong *model.Song, selectedAlbum *model.Album, c *cart.Cart, scanner *bufio.Scanner) {
	for {
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Add this song to cart")
		fmt.Println("2. Add this album to cart")
		fmt.Println("3. View cart")
		fmt.Println("0. Back to artist selection")
		fmt.Print("Choose: ")

		choice, err := helper.ReadInt(scanner)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}
		switch choice {
		case 1:
			c.AddSong(selectedSong)
		case 2:
			c.AddAlbum(selectedAlbum)
		case 3:
			ShowCartAndActions(c, scanner)
		case 0:
			return
		default:
			fmt.Println("Invalid option.")
			continue
		}
		if choice == 1 || choice == 2 {
			if !helper.AskToContinue(scanner) {
				return
			}
		}
	}
}

func RunStore(artists []*model.Artist, scanner *bufio.Scanner) {
	c := cart.NewCart()

	for {
		selectedArtist := ChooseArtist(artists, scanner)
		if selectedArtist == nil {
			break
		}
		fmt.Printf("Selected artist: %s\n", selectedArtist.GetName())

		selectedAlbum := ChooseAlbum(selectedArtist.GetAlbums(), scanner)
		if selectedAlbum == nil {
			break
		}
		fmt.Printf("Selected album: %s\n", selectedAlbum.GetTitle())

		selectedSong := ChooseSong(selectedAlbum.GetSongs(), scanner)
		if selectedSong == nil {
			break
		}
		fmt.Printf("Selected song: %s\n", selectedSong.GetTitle())

		ShowActionMenu(selectedSong, selectedAlbum, c, scanner)
	}

	fmt.Println("Goodbye!")
}
