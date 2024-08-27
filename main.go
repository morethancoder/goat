package main

import (
	"context"
	"log"
	"miniapp/components"
	"os"
)

func main() {
    file, err := os.Create("index.html")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    err = components.App().Render(context.Background(), file)
    if err != nil {
        log.Fatal(err)
    }
}
