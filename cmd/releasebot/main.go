package main

import (
	"flag"
	"os"
	"time"

	"github.com/joho/godotenv"
	zl "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/vadapavmov/releasebot/internal/bot"
	"github.com/vadapavmov/releasebot/internal/imdb"
	"github.com/vadapavmov/releasebot/internal/structs"
	"github.com/vadapavmov/releasebot/internal/tmdb"
)

func main() {
	// Load .env file
	_ = godotenv.Load()

	// Define flags
	apiKey := flag.String("tmdb-api-key", os.Getenv("TMDB_API_KEY"), "TMDB API Key")
	endpoint := flag.String("imdb-api-endpoint", os.Getenv("IMDB_API_ENDPOINT"), "IMDB API Endpoint")
	token := flag.String("token", os.Getenv("TOKEN"), "Discord Bot Token")
	guild := flag.String("guild", os.Getenv("GUILD"), "Discord Guild ID")

	flag.Parse()

	// Setup logger
	log.Logger = zl.New(zl.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	zl.SetGlobalLevel(zl.InfoLevel)

	// Validate required parameters
	if *apiKey == "" && *endpoint == "" {
		log.Fatal().Msg("one of TMDB_API_KEY or IMDB_API_ENDPOINT must be provided")
	}
	if *token == "" {
		log.Fatal().Msg("discord token missing")
	}

	// Initialize search engine
	var engine structs.SearchEngine
	if *endpoint != "" {
		engine = imdb.New(*endpoint)
		log.Info().Msg("initialized IMDB as data source")
	}
	if engine == nil && *apiKey != "" {
		engine = tmdb.New(*apiKey)
		log.Info().Msg("initialized TMDB as data source")
	}

	// Initialize Discord bot
	session := bot.New(*token, *guild, engine)

	// Start the bot
	err := session.Open()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open discord bot session")
	}
	defer session.Close()

	log.Info().Msg("bot is now running. press CTRL+C to exit")
	select {} // Block forever
}
