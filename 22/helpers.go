package main

import (
	"log"
	"os"
)

func GetInputFile(in string) *os.File {
	f, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
