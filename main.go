package main

import (
	"log"

	"github.com/valp0/go-aws-pg/app"
)

// Run server for API, handle server error
func main() {
	err := app.RunServer()
	if err != nil {
		log.Fatal(err.Error())
	}
}
