package main

import (
	"github.com/gerbilweb/example/src"
	"github.com/gerbilweb/gerbil"
)

func main() {
	app := gerbil.New()
	defer app.Run()

	app.Route("/", &src.App{})
}
