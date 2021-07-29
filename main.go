package main

import (
	"github.com/danielmalka/dice-go/cmd"
	"log"
	"os"
)

func main() {
	command := cmd.Start()
	if err := command.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
