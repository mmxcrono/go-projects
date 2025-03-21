package main

import (
	"fmt"

	"github.com/mmxcrono/go-projects/sql-api/internal/db"
)

func main() {
	db.Connect()

	albums, err := db.AlbumsByArtist("prince")

	if err != nil {
		return
	}

	for _, v := range albums {
		fmt.Printf("ID: %v, Title: %v, Price: %v\n", v.ID, v.Title, v.Price)
	}
}
