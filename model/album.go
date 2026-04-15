package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Album struct {
	id     uuid.UUID
	title  string
	artist *Artist
	price  float64
	songs  []*Song
}

func NewAlbum(title string, artist *Artist, price float64, songs []*Song) *Album {
	return &Album{
		id:     uuid.New(),
		title:  title,
		artist: artist,
		price:  price,
		songs:  songs,
	}
}

func (a *Album) GetID() uuid.UUID {
	return a.id
}

func (a *Album) GetTitle() string {
	return a.title
}

func (a *Album) GetArtist() *Artist {
	return a.artist
}

func (a *Album) GetPrice() float64 {
	return a.price
}

func (a *Album) GetSongs() []*Song {
	return a.songs
}

func (a *Album) SetTitle(title string) {
	a.title = title
}

func (a *Album) SetArtist(artist *Artist) {
	a.artist = artist
}

func (a *Album) SetPrice(price float64) {
	a.price = price
}

func (a *Album) SetSongs(songs []*Song) {
	a.songs = songs
}

func (a *Album) AddSong(song *Song) {
	a.songs = append(a.songs, song)
}

func (a *Album) RemoveSong(song *Song) {
	for i, s := range a.songs {
		if s.GetID() == song.GetID() {
			a.songs = append(a.songs[:i], a.songs[i+1:]...)
			return
		}
	}
}

func (a *Album) String() string {
	return fmt.Sprintf("Album: %s (ID: %s)", a.title, a.id)
}
