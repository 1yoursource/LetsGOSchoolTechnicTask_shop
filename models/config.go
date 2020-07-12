package models

import (
	"os"
	"webServer/db"
)

var server = os.Getenv("DATABASE")

var databaseName = os.Getenv("DATABASE_NAME")

var dbConnect = db.NewConnection(server, databaseName)
