package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)

	err := NewCLI().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
