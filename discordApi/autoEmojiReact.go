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

func (d Discord) processAutoEmojiReactCommand(interaction *discordgo.InteractionCreate) {
	// var options = [2]string{"Channel", "Emoji"}
	data := interaction.ApplicationCommandData()
	session := d.Session
	user := interaction.Member.Nick

	// perms, error := session.ApplicationCommandPermissions(appID, d.GuildId, data.ID)

	// if len(perms.Permissions) == 0 {
	// 	permissions := discordgo.ApplicationCommandPermissionsList{
	// 		Permissions: []*discordgo.ApplicationCommandPermissions{
	// 			{
	// 				ID:         config.Config.Discord.ModeratorRoleID,
	// 				Type:       discordgo.ApplicationCommandPermissionTypeRole,
	// 				Permission: true,
	// 			},
	// 			{
	// 				ID:         config.Config.Discord.MemberRoleID,
	// 				Type:       discordgo.ApplicationCommandPermissionTypeRole,
	// 				Permission: false,
	// 			},
	// 		},
	// 	},
	// 		session.ApplicationCommandPermissionsEdit(appID, d.GuildId, data.ID)

	// }

	if len(data.Options) == 0 {
		session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("sorry, %s, but i cant seem to find your options :(", user),
			},
		})
	}

	// add data type validation
	channelID := data.GetOption("Channel").StringValue()
	// assume emoji will be sent in the form of
	// <:derpdurian:1452714098847383683>
	// or just name?? unicode??
	emoji := data.GetOption("Emoji").StringValue()

	emojiName := ""
	emojiID := ""

	if strings.Contains(emoji, "<") {
		emojiArr := strings.Split(strings.Trim(emoji, "<>"), ":")
		emojiName = emojiArr[0]
		emojiID = emojiArr[1]
	} else {
		emojiID = emoji
	}

	emojiStruct := discordgo.Emoji{
		ID:   emojiID,
		Name: emojiName,
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		message := m.Message
		if message.ChannelID == channelID {
			session.MessageReactionAdd(channelID, message.ID, emojiStruct.APIName())

		}
	})

}
