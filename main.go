package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	DISCORD string
	TMDB    string
	KEY     string

	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer
)

var s *discordgo.Session

func init() {

	var err error
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	DISCORD = os.Getenv("DISCORD_TOKEN")
	TMDB = os.Getenv("TOKEN")
	KEY = os.Getenv("KEY")

	s, err = discordgo.New("Bot " + DISCORD)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	myIntent := discordgo.Intent(discordgo.IntentsAll)
	fmt.Println(myIntent)
}

func main() {
	err := s.Open()

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {

		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "1184494963241779240", v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
	
		if true {
			log.Println("Removing commands...")
			for _, v := range registeredCommands {
				err := s.ApplicationCommandDelete(s.State.User.ID, "1184494963241779240", v.ID)
				if err != nil {
					log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
				}
			}
		}

}
