package main

import "steplet/webserver/cmd/app"

func main() {
	server := app.NewServer(":8080")

	server.Run()
}
