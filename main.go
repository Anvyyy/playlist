package main

import (
	"log"

	"github.com/Anvyyy/playlist/internal/app"
)

// run app
func main() {
	if err := app.Run(":50051", ":50081"); err != nil {
		log.Fatal(err)
	}
}
