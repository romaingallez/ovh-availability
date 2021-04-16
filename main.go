package main

import (
	"log"
	"os"
	"ovh-availability/api"
	"ovh-availability/database"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB()

	var runParsing bool = true
	if runParsing {
		token := os.Getenv("DISCORD_TOKEN")

		config := GetConfig()

		api.Parsing(config, token)
	}

}
