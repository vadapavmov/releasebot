package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DISCORD string
	TMDB  string
	KEY string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	DISCORD = os.Getenv("DISCORD_TOKEN")
	TMDB = os.Getenv("TOKEN")
}
