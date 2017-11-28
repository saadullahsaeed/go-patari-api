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

	res, err := pc.Search("waqt")
	fmt.Println(res.Data.Song[0].Stream)
	fmt.Println(err)
}
