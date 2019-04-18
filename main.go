package main

import (
	"github.com/triest/http2"
)

func main() {
	var url string = "https://swapi.co/api/planets"
	http2.GetRequwestParams(url)
}
