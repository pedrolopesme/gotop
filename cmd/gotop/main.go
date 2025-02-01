package main

import (
	"log"

	"github.com/pedrolopesme/books/internal/tui"
)

func main() {
	if err := tui.StartTUI(); err != nil {
		log.Fatal(err)
	}
}
