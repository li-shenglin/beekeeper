package model

type Role string

var (
	Read  Role = "Read"
	Write Role = "Write"
	Admin Role = "Admin"
)

type DocType string

var (
	DIR DocType = "DIR"
	API DocType = "API"
)
