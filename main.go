package main

import (
	"context"
	"log"
	"morethancoder/goat/components"
	"os"
)

func main() {
	file, err := os.Create("public/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = components.App().Render(context.Background(), file)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully generated index.html")
}
