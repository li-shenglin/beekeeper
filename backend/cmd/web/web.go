package main

import (
	"backend/base/rest"
	"backend/web/resource"
)

func main() {
	app := rest.New(16386)
	app.Mapping("/api/session", resource.NewSessionHandler())
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
