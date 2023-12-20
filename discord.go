package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var base string = "https://api.themoviedb.org/3"

type Data struct {
	Adult               bool        `json:"adult"`
	BackdropPath        string      `json:"backdrop_path"`
	BelongsToCollection interface{} `json:"belongs_to_collection"`
	Budget              int         `json:"budget"`
	Genres              []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string        `json:"homepage"`
	ID                  int           `json:"id"`
	ImdbID              interface{}   `json:"imdb_id"`
	OriginalLanguage    string        `json:"original_language"`
	OriginalTitle       string        `json:"original_title"`
	Overview            string        `json:"overview"`
	Popularity          float64       `json:"popularity"`
	PosterPath          string        `json:"poster_path"`
	ProductionCompanies []interface{} `json:"production_companies"`
	ProductionCountries []interface{} `json:"production_countries"`
	ReleaseDate         string        `json:"release_date"`
	Revenue             int           `json:"revenue"`
	Runtime             int           `json:"runtime"`
	SpokenLanguages     []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}

func DiscordHandler() {

	dg, err := discordgo.New("Bot " + DISCORD)

	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "hello":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "HELLO",
				},
			})
		}
	}
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		SendMessages(s, m)
	})
	cmd, err := s.ApplicationCommandCreate(dg.State.User.ID, *GuildID, v)
	if err != nil {
		log.Panicf("Cannot create '%v' command: %v", v.Name, err)
	}

	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = dg.Open()

	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	defer dg.Close()

	fmt.Println("AetherAI is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func SendMessages(s *discordgo.Session, m *discordgo.MessageCreate) {

	//tv/21212?language=en-US' \
	searc := base + "/movie/1219926?language=en-US"
	data := SendRequest(searc, TMDB)

	var movie Data

	err := json.Unmarshal([]byte(data), &movie)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(movie.OriginalTitle)
	fmt.Println(movie.OriginalLanguage)

	if m.Author.ID == s.State.User.ID {
		return
	}

	s.ChannelMessageSend(m.ChannelID, movie.OriginalTitle)
}

func SlashCommandCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {

	fmt.Println("HELLo")
	if i.Type == discordgo.InteractionApplicationCommand {
		switch i.ApplicationCommandData().Name {
		case "hello":
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "HELLO",
				},
			})
		}
	}
}
