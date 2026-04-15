package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Song struct {
	id     uuid.UUID
	title  string
	artist *Artist
	price  float64
}

func NewSong(title string, artist *Artist, price float64) *Song {
	return &Song{
		id:     uuid.New(),
		title:  title,
		artist: artist,
		price:  price,
	}
}

func (s *Song) GetID() uuid.UUID {
	return s.id
}

func (s *Song) GetTitle() string {
	return s.title
}

func (s *Song) GetArtist() *Artist {
	return s.artist
}

func (s *Song) GetPrice() float64 {
	return s.price
}

func (s *Song) String() string {
	return fmt.Sprintf("Song: %s (ID: %s)", s.title, s.id)
}
