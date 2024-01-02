package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/vadapavmov/releasebot/internal/structs"
)

// New creates a new discord session with the given token and tmdb client
func New(token, guildID string, c structs.SearchEngine) *discordgo.Session {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create discord session")
	}
	// Register interaction
	s.AddHandler(makeInteractionHandler(c))

	// On login register commands
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info().Msgf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		registerCommands(s, guildID)
	})

	return s
}

// registerCommands registers the application commands
func registerCommands(s *discordgo.Session, guildID string) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "release-movie",
			Description: "Fetches movie for given id from tmdb and make a post in current channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "Movie ID",
					Required:    true,
				},
			},
		},
		{
			Name:        "release-tv",
			Description: "Fetches tv for given id from tmdb and make a post in current channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "TV ID",
					Required:    true,
				},
			},
		},
	}

	// Iterate over the commands and register them
	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd)
		if err != nil {
			log.Error().Err(err).Msgf("Cannot create '%v' command", cmd.Name)
		} else {
			log.Info().Msgf("Registered command '%v'", cmd.Name)
		}
	}
}

// makeInteractionHandler creates a handler function for Discord interactions
func makeInteractionHandler(c structs.SearchEngine) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		cmdName := i.ApplicationCommandData().Name
		if cmdName == "release-tv" || cmdName == "release-movie" {
			handleRelease(c, s, i, cmdName)
		}
	}
}

// handleRelease handles both movie and TV release commands
func handleRelease(c structs.SearchEngine, s *discordgo.Session, i *discordgo.InteractionCreate, commandType string) {
	id := i.ApplicationCommandData().Options[0].StringValue()
	var collection structs.Collection
	var err error

	// Immediately acknowledge the interaction
	if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	}); err != nil {
		log.Warn().Err(err).Str("command", commandType).Msg("failed to acknowledge interaction")
	}

	if commandType == "release-movie" {
		collection, err = c.GetMovie(id)
	} else {
		collection, err = c.GetTv(id)
	}

	if err != nil {
		sendErrorReply(s, i.Interaction, err)
		return
	}

	poster, err := collection.Poster()
	if err != nil {
		sendErrorReply(s, i.Interaction, err)
		return
	}
	content := Format(collection)
	if _, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &content,
		Files: []*discordgo.File{
			{
				ContentType: "image/jpeg",
				Name:        "poster.jpeg",
				Reader:      poster,
			},
		},
	}); err != nil {
		log.Warn().Err(err).Str("id", id).Str("command", commandType).Msg("failed to respond to interaction")
	}
}

// sendErrorReply sends an error reply to a Discord interaction
func sendErrorReply(s *discordgo.Session, i *discordgo.Interaction, err error) {
	errMsg := ">**failed to execute command**\n" + err.Error()
	if _, err = s.InteractionResponseEdit(i, &discordgo.WebhookEdit{
		Content: &errMsg,
	}); err != nil {
		log.Warn().Err(err).Msg("failed to send error reply")
	}
}

// Format formats the tmdb collection data for display
func Format(c structs.Collection) string {
	str := []string{
		"## %s",
		"%s",
		"",
		"- 🎭 **Genres:**  %s",
		"- 🗣 **Language:**  %s",
		"- 📅 **Release Date:**  %s",
		"- ⭐ **Rating:**  %s",
	}
	return fmt.Sprintf(strings.Join(str, "\n"), c.Name(), c.Description(), c.GenreStr(), c.Language(), c.ReleaseTime(), c.Star())
}
