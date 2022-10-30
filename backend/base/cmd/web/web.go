package main

import (
	"backend/base/rest"
	"backend/web"
)

func main() {
	app := rest.New(16386)
	app.Mapping("/api/session", web.NewSessionHandler())
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
