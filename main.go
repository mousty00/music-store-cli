package main

import (
	"bufio"
	"go-starter/data"
	"go-starter/ui"
	"os"
)

func main() {
	artists := data.SeedArtists()
	scanner := bufio.NewScanner(os.Stdin)
	ui.RunStore(artists, scanner)
}
