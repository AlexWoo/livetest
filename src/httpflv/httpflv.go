package main

import (
	"log"
	"media"
	"media/flv"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s httpflv_url", os.Args[0])
	}

	res, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalf("GET \"%v\" failed", os.Args[1])
	}
	defer res.Body.Close()

	log := media.NewLog(media.LogInfo, os.Stdout)

	parser := flv.Create(res.Body, log)
	parser.Parser()
}
