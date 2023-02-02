package main

import (
	"github.com/storyTunes/backend.git/cmd/postgres"
	"github.com/storyTunes/backend.git/cmd/server"
	"github.com/storyTunes/backend.git/cmd/utils"
)

func main() {
	postgres.ConnectPostgresDatabase()
	defer postgres.ClosePostgresDatabase()
	utils.InitLogger()
	postgres.InitPostgresDB()
	server.StartServer()
}
