package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"flag"

	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

var (
	base string = "https://api.themoviedb.org/3"
)

var GuildID = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")

func init() { flag.Parse() }

func DiscordHandler() {
	dg, err := discordgo.New("Bot " + DISCORD)

	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "basic-command",
			Description: "Basic command",
		},
	}

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		SendMessages(s, m)
	})

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

	//var u string
	//tv/21212?language=en-US' \
	urlurl := base + "/movie/1219926?language=en-US"
	data := SendRequest(urlurl, TMDB)

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

}
