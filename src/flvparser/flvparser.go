package main

import (
	"log"
	"media"
	"media/flv"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s flvfile", os.Args[0])
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Open %v failed", os.Args[1])
	}
	defer f.Close()

	log := media.NewLog(media.LogInfo, os.Stdout)

	parser := flv.Create(f, log)
	parser.Parser()
}
