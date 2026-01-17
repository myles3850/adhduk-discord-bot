package discordapi

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// TODOs
// - make this command mod-only
// - add emoji react limit per channel (1? 2?)
// - error handling for any failures
// - error message for max emoji limit
// - command for removing channel-emoji setup

func (d Discord) ProcessAutoEmojiReactCommand(interaction *discordgo.InteractionCreate) {
	data := interaction.ApplicationCommandData()
	session := d.Session

	if len(data.Options) == 0 {
		session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "sorry, but i cant seem to find your options :(",
			},
		})
	}

	channel := data.GetOption("channel").ChannelValue(session)
	emoji := data.GetOption("emoji").StringValue()

	fmt.Print(emoji)

	d.Session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Added %s reaction to channel %s", emoji, channel.Name),
		},
	})

	var emojiID string

	if strings.Contains(emoji, "<") {
		emojiID = strings.Trim(emoji, "<>")
	} else {
		emojiID = emoji
	}

	fmt.Print(emojiID)

	// for final ver - check that it doesnt respond to bots AND mods
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		message := m.Message
		if message.ChannelID == channel.ID {
			err := session.MessageReactionAdd(channel.ID, message.ID, emojiID)

			if err != nil {
				fmt.Printf("%+v", err)
				session.ChannelMessageSend(interaction.ChannelID, fmt.Sprintf("%+v", err))
			}

		}
	})

}
