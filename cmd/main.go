package main

import (
	"log"

	"github.com/arafatazam/mini-twitter/app"
)

func main() {
    app := app.New()
    log.Fatal(app.Run())
}