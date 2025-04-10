package main

import (
	"covet.digital/dashboard/cmd/web/server"
)

func main() {
	app, err := server.NewApp()
	if err != nil {
		panic(err.Error())
	}
	if err := app.Run(); err != nil {
		panic(err.Error())
	}
}
