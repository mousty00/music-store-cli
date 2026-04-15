package model

import (
	"fmt"

	"github.com/google/uuid"
)

type Artist struct {
	id     uuid.UUID
	name   string
	albums []*Album
	songs  []*Song
}

func NewArtist(name string) *Artist {
	return &Artist{
		id:     uuid.New(),
		name:   name,
		albums: []*Album{},
		songs:  []*Song{},
	}
}

func (a *Artist) GetID() uuid.UUID {
	return a.id
}

func (a *Artist) GetName() string {
	return a.name
}

func (a *Artist) SetName(name string) {
	a.name = name
}

func (a *Artist) String() string {
	return fmt.Sprintf("Artist: %s (ID: %s)", a.name, a.id)
}

func (a *Artist) SetAlbums(albums []*Album) {
	a.albums = albums
}

func (a *Artist) GetAlbums() []*Album {
	return a.albums
}

func (a *Artist) GetSongs() []*Song {
	return a.songs
}

func (a *Artist) AddAlbum(album *Album) {
	a.albums = append(a.albums, album)
}

func (a *Artist) AddSong(song *Song) {
	a.songs = append(a.songs, song)
}

func (a *Artist) RemoveAlbum(album *Album) {
	newAlbums := make([]*Album, 0, len(a.albums))
	for _, alb := range a.albums {
		if alb.GetID() != album.GetID() {
			newAlbums = append(newAlbums, alb)
		}
	}
	a.albums = newAlbums
}
