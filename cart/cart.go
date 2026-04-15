package cart

import (
	"fmt"
	"go-starter/model"
)

type CartItem struct {
	ItemType string
	Item     interface{}
}

type Cart struct {
	Items []CartItem
}

func NewCart() *Cart {
	return &Cart{Items: []CartItem{}}
}

func (c *Cart) AddSong(song *model.Song) {
	c.Items = append(c.Items, CartItem{ItemType: "song", Item: song})
	fmt.Printf("Added song \"%s\" to cart.\n", song.GetTitle())
}

func (c *Cart) AddAlbum(album *model.Album) {
	c.Items = append(c.Items, CartItem{ItemType: "album", Item: album})
	fmt.Printf("Added album \"%s\" to cart.\n", album.GetTitle())
}

func (c *Cart) RemoveItem(index int) error {
	if index < 0 || index >= len(c.Items) {
		return fmt.Errorf("invalid index")
	}
	removed := c.Items[index]
	c.Items = append(c.Items[:index], c.Items[index+1:]...)
	switch removed.ItemType {
	case "song":
		song := removed.Item.(*model.Song)
		fmt.Printf("Removed song \"%s\" from cart.\n", song.GetTitle())
	case "album":
		album := removed.Item.(*model.Album)
		fmt.Printf("Removed album \"%s\" from cart.\n", album.GetTitle())
	}
	return nil
}

func (c *Cart) Total() float64 {
	var total float64
	for _, item := range c.Items {
		switch item.ItemType {
		case "song":
			total += item.Item.(*model.Song).GetPrice()
		case "album":
			total += item.Item.(*model.Album).GetPrice()
		}
	}
	return total
}

func (c *Cart) Clear() {
	c.Items = []CartItem{}
	fmt.Println("Cart cleared.")
}

func (c *Cart) IsEmpty() bool {
	return len(c.Items) == 0
}

func (c *Cart) Display() {
	if c.IsEmpty() {
		fmt.Println("\n🛒 Your cart is empty.")
		return
	}
	fmt.Println("\n🛒 Your Cart:")
	fmt.Println("-----------------------------------")
	for i, item := range c.Items {
		switch item.ItemType {
		case "song":
			song := item.Item.(*model.Song)
			fmt.Printf("%d. [Song]  %s - %s  $%.2f\n",
				i+1, song.GetArtist().GetName(), song.GetTitle(), song.GetPrice())
		case "album":
			album := item.Item.(*model.Album)
			fmt.Printf("%d. [Album] %s - %s  $%.2f\n",
				i+1, album.GetArtist().GetName(), album.GetTitle(), album.GetPrice())
		}
	}
	fmt.Println("-----------------------------------")
	fmt.Printf("Total: $%.2f\n", c.Total())
}
