package main

import (
	"fmt"
	"log"

	"github.com/saadullahsaeed/go-patari-api/lib/patari"
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
}
