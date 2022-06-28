package main

import (
	"demo-cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Heating the engines...")

	cliApp := app.Generate()

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
