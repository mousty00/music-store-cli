package data

import "go-starter/model"

func SeedArtists() []*model.Artist {
	artist := model.NewArtist("The Beatles")
	artist2 := model.NewArtist("The Rolling Stones")
	artist3 := model.NewArtist("Queen")
	artist4 := model.NewArtist("Michael Jackson")
	artist5 := model.NewArtist("Drake")
	artist6 := model.NewArtist("Travis Scott")
	artist7 := model.NewArtist("Taylor Swift")

	// The Beatles songs
	artist.AddSong(model.NewSong("Hey Jude", artist, 1.99))
	artist.AddSong(model.NewSong("Let It Be", artist, 1.99))
	artist.AddSong(model.NewSong("Yesterday", artist, 1.99))

	// The Rolling Stones songs
	artist2.AddSong(model.NewSong("Satisfaction", artist2, 1.99))
	artist2.AddSong(model.NewSong("Paint It Black", artist2, 1.99))
	artist2.AddSong(model.NewSong("Sympathy for the Devil", artist2, 1.99))

	// Queen songs
	artist3.AddSong(model.NewSong("Bohemian Rhapsody", artist3, 1.99))
	artist3.AddSong(model.NewSong("We Will Rock You", artist3, 1.99))
	artist3.AddSong(model.NewSong("We Are the Champions", artist3, 1.99))

	// Michael Jackson songs
	artist4.AddSong(model.NewSong("Billie Jean", artist4, 1.99))
	artist4.AddSong(model.NewSong("Thriller", artist4, 1.99))
	artist4.AddSong(model.NewSong("Beat It", artist4, 1.99))

	// Drake songs
	artist5.AddSong(model.NewSong("God's Plan", artist5, 1.99))
	artist5.AddSong(model.NewSong("In My Feelings", artist5, 1.99))
	artist5.AddSong(model.NewSong("Hotline Bling", artist5, 1.99))

	// Travis Scott songs
	artist6.AddSong(model.NewSong("Sicko Mode", artist6, 1.99))
	artist6.AddSong(model.NewSong("Goosebumps", artist6, 1.99))
	artist6.AddSong(model.NewSong("Highest in the Room", artist6, 1.99))

	// Taylor Swift songs
	artist7.AddSong(model.NewSong("Shake It Off", artist7, 1.99))
	artist7.AddSong(model.NewSong("Blank Space", artist7, 1.99))
	artist7.AddSong(model.NewSong("Love Story", artist7, 1.99))

	// Albums
	album := model.NewAlbum("Abbey Road", artist, 19.99, artist.GetSongs())
	album2 := model.NewAlbum("Sticky Fingers", artist2, 19.99, artist2.GetSongs())
	album3 := model.NewAlbum("A Night at the Opera", artist3, 19.99, artist3.GetSongs())
	album4 := model.NewAlbum("Thriller", artist4, 19.99, artist4.GetSongs())
	album5 := model.NewAlbum("Scorpion", artist5, 19.99, artist5.GetSongs())
	album6 := model.NewAlbum("Astroworld", artist6, 19.99, artist6.GetSongs())
	album7 := model.NewAlbum("1989", artist7, 19.99, artist7.GetSongs())

	artist.SetAlbums([]*model.Album{album})
	artist2.SetAlbums([]*model.Album{album2})
	artist3.SetAlbums([]*model.Album{album3})
	artist4.SetAlbums([]*model.Album{album4})
	artist5.SetAlbums([]*model.Album{album5})
	artist6.SetAlbums([]*model.Album{album6})
	artist7.SetAlbums([]*model.Album{album7})

	return []*model.Artist{artist, artist2, artist3, artist4, artist5, artist6, artist7}
}
