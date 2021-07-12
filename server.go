package main

import (
	"linebot-server/router"
	"log"
)

func main() {
	// lib.SetupRichMenu()

	router := router.SetupRouter()
	err := router.Run(":7000")
	if err != nil {
		log.Fatal(err)
	}
}
