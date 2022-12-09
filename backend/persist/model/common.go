package model

type Role string

var (
	Read         Role = "Read"
	ReadAndWrite Role = "ReadAndWrite"
	Admin        Role = "Admin"
	Owner        Role = "Owner"
)

type DocType string

var (
	DIR DocType = "DIR"

	GET    DocType = "GET"
	POST   DocType = "POST"
	PUT    DocType = "PUT"
	DELETE DocType = "DELETE"
	PATCH  DocType = "PATCH"

	HEAD    DocType = "HEAD"
	OPTIONS DocType = "OPTIONS"
	TRACE   DocType = "TRACE"
)
