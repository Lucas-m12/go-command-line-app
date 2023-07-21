package main

import (
	"command-line-app/app"
	"log"
	"os"
)

func main() {
	println("Pilot")
	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
