package main

import (
	"backend/base/rest"
	"backend/web/resource"
)

func main() {
	app := rest.New(16386)
	app.Mapping("/api/v1/session", resource.NewSessionHandler())

	app.Mapping("/api/v1/projects", resource.NewProjectsHandler())
	app.Mapping("/api/v1/projects/:id", resource.NewProjectHandler())

	app.Mapping("/api/v1/users", resource.NewUsersHandler())
	app.Mapping("/api/v1/users/:id", resource.NewUserHandler())
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
