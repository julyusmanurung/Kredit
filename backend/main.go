package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/julyusmanurung/Kredit/api"
	"github.com/julyusmanurung/Kredit/database"
)

func main() {
	db, err := database.SetupDb()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()
}
