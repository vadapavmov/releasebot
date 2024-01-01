package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/adler32"

	"github.com/bwmarrin/discordgo"
)

var base string = "https://api.themoviedb.org/3"

func Format(data Data) string {

	str := `## %s
%s
- **Year**: %s
- **Language**: %s
- **Genres**: #%s
`

	//new_str := fmt.Sprintf(str, data.Title, data.Overview, data.ReleaseDate, data.OriginalLanguage, data.Genres[0].Name, "https://image.tmdb.org/t/p/original/"+data.PosterPath)

	new_str := fmt.Sprintf(str, data.Title, data.Overview, data.ReleaseDate, data.OriginalLanguage, data.Genres[0].Name)
	return new_str

}

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "release-movie",
		Description: "Basic command",

		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "movie-id",
				Required:    true,
			},
		},
	},
	{
		Name:        "release-tv",
		Description: "Basic command",

		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "tv-id",
				Required:    true,
			},
		},
	},
	{
		Name:        "release-bolly",
		Description: "Basic command",

		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "bollywood-id",
				Required:    true,
			},
		},
	},
	{
		Name:        "release-anime",
		Description: "Basic command",

		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "anime-id",
				Required:    true,
			},
		},
	},
}

//1219926

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"release-movie": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		var t string = options[0].StringValue()

		url := fmt.Sprintf("%s%s%s?language=en-US", base, "/movie/", t)
		data := SendRequest(url, TMDB)

		var movie Data

		pic := "https://image.tmdb.org/t/p/original/" + movie.PosterPath
		b, err := DownloadImage(pic)

		if err != nil {
			fmt.Println("Error:", err)
		}

		err = json.Unmarshal([]byte(data), &movie)

		if err != nil {
			fmt.Println("Error:", err)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: Format(movie),

				Files: []*discordgo.File{
					{
						ContentType: "image/jpeg",
						Name:        "pic.jpeg",
						Reader:      bytes.NewReader(b),
					},
				},
			},
		})
	},

	"release-tv": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		var t string = options[0].StringValue()

		url := fmt.Sprintf("%s%s%s?language=en-US", base, "/tv/", t)
		data := SendRequest(url, TMDB)

		var movie Data

		err := json.Unmarshal([]byte(data), &movie)
		if err != nil {
			fmt.Println("Error:", err)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: Format(movie),
			},
		})
	},
	"release-anime": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		var t string = options[0].StringValue()

		url := fmt.Sprintf("%s%s%s?language=en-US", base, "/anime/", t)
		data := SendRequest(url, TMDB)

		var movie Data

		err := json.Unmarshal([]byte(data), &movie)
		if err != nil {
			fmt.Println("Error:", err)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: Format(movie),
			},
		})
	},

	"release-bolly": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		var t string = options[0].StringValue()

		url := fmt.Sprintf("%s%s%s?language=en-US", base, "/bollywood/", t)
		data := SendRequest(url, TMDB)

		var movie Data

		err := json.Unmarshal([]byte(data), &movie)
		if err != nil {
			fmt.Println("Error:", err)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: Format(movie),
			},
		})
	},
}
