package discordapi

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (d Discord) processAutoEmojiReactCommand(interaction *discordgo.InteractionCreate) {
	// var options = [2]string{"Channel", "Emoji"}
	data := interaction.ApplicationCommandData()
	session := d.Session
	user := interaction.Member.Nick

	if len(data.Options) == 0 {
		session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("sorry, %s, but i cant seem to find your options :(", user),
			},
		})
	}

	channelID := data.GetOption("Channel")
	emoji := data.GetOption("Emoji")

	// channel info should be an ID?

}
