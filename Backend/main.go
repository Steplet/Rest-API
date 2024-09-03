package main

import (
	"log"
	"steplet/webserver/cmd/app"
)

func main() {
	store, err := app.NewPostgesStor()

	if err != nil {
		log.Fatal(err)
	}

	store.InitUserTable()

	if err != nil {
		log.Fatal(err)
	}
	server := app.NewServer(":8080", store)

	server.Run()
}
