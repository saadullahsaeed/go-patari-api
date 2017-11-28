# go-patari-api
This is a simple GoLang package to access the Patari API. 

I wrote this to get Patari playing on Sonos so I didn't focus on implementing all of the functionality available in the API and only the features I needed.

Supported API Methods:
* Search
* Get Playlist

## Installation
```go get github.com/saadullahsaeed/go-patari-api```

### Usage

```go
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

	// Search
	res, err := pc.Search("bulleya")
	fmt.Println(res.Data.Song[0])
	fmt.Println(err)

	// Get Playlist (use either playlist ID or the discovery slug)
	pd, err := pc.GetPlaylist("Rock")
	fmt.Println(pd.Songs[0])
	fmt.Println(err)
}

```



