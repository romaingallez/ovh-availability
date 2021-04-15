package main

import (
	"log"
	"os"
	"ovh-availability/api"

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

	token := os.Getenv("DISCORD_TOKEN")

	config := GetConfig()
	api.Parsing(config, token)
}
