package main

import (
	"flag"
	"os"
	"time"

	"github.com/joho/godotenv"
	zl "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/vadapavmov/releasebot/internal/bot"
	"github.com/vadapavmov/releasebot/internal/tmdb"
)

func main() {
	// Load .env file
	_ = godotenv.Load()

	// Define flags
	apiKey := flag.String("api_key", os.Getenv("API_KEY"), "TMDB API Key")
	token := flag.String("token", os.Getenv("TOKEN"), "Discord Bot Token")
	guild := flag.String("guild", os.Getenv("GUILD"), "Discord Guild ID")

	flag.Parse()

	// Setup logger
	log.Logger = zl.New(zl.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	zl.SetGlobalLevel(zl.InfoLevel)

	// Validate required parameters
	if *apiKey == "" || *token == "" {
		log.Fatal().Msg("API_KEY or TOKEN are missing")
	}

	// Initialize TMDB client
	tmdbClient := tmdb.New(*apiKey)

	// Initialize Discord bot
	session := bot.New(*token, *guild, tmdbClient)

	// Start the bot
	err := session.Open()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open discord bot session")
	}
	defer session.Close()

	log.Info().Msg("bot is now running. press CTRL+C to exit")
	select {} // Block forever
}
