package main

import (
	"fmt"
	"log"

	patari "github.com/saadullahsaeed/go-patari-api/lib"
)

func main() {
	pc := patari.NewClient()
	err := pc.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	res, err := pc.Search("bulleya")
	fmt.Println(res.Data.Song[0])
	fmt.Println(err)

	pd, err := pc.GetPlaylist("Rock")
	fmt.Println(pd.Songs[0])
	fmt.Println(err)
}
