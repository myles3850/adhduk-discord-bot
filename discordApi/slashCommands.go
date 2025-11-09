package discordapi

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

type CommandName struct {
	wheel     string
	eightBall string
}

var names = &CommandName{wheel: "wheel", eightBall: "eight_ball"}

func (d Discord) RegisterCommands() {
	s := d.Session
	appID := s.State.User.ID
	guildID := d.GuildId
	minStringLength := 2

	commands := []*discordgo.ApplicationCommand{
		{
			Name:        names.wheel,
			Description: "Give me a selection, and ill pick one for you",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "option1",
					Description: "Choice One",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					MinLength:   &minStringLength,
				},
				{
					Name:        "option2",
					Description: "Choice Two",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					MinLength:   &minStringLength,
				},
				{
					Name:        "option3",
					Description: "Choice Three",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
					MinLength:   &minStringLength,
				},
				{
					Name:        "option4",
					Description: "Choice Four",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
					MinLength:   &minStringLength,
				},
				{
					Name:        "option5",
					Description: "Choice Five",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
					MinLength:   &minStringLength,
				},
				{
					Name:        "option6",
					Description: "Choice Six",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    false,
					MinLength:   &minStringLength,
				},
			},
		},
		{
			Name:        names.eightBall,
			Description: "ask DeeDee to shake the mystical 8 ball for you",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "question",
					Description: "give me your question so i can find out the answer",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					MinLength:   &minStringLength,
				},
			},
		},
	}

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(appID, guildID, cmd)
		if err != nil {
			fmt.Printf("❌ Cannot create command '%v': %v\n", cmd.Name, err)
		} else {
			fmt.Printf("✅ Registered command '%v'\n", cmd.Name)
		}
	}
}

func (d Discord) OnInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()

	switch data.Name {
	case names.wheel:
		d.processWheelCommand(i)
		return
	case names.eightBall:
		d.process8BallCommand(i)
		return
	}
}

//from here all functions are processing functions

func (d Discord) processWheelCommand(interaction *discordgo.InteractionCreate) {
	var options []string
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
	for _, option := range data.Options {
		options = append(options, option.StringValue())
	}

	// i know im going to forget how this shit works, it gets the length of options and passes to intn which gives a random number based on how many are in the slice
	chosen := options[rand.Intn(len(options))]

	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Interesting choices... im feeling %s this time.", chosen),
		},
	})
}

func (d Discord) process8BallCommand(interaction *discordgo.InteractionCreate) {

	data := interaction.ApplicationCommandData()
	session := d.Session
	user := interaction.Member.Nick


	question := data.Options[0].StringValue()
	ballAnswers := []string{"Yes", "No"}

	if len(question) == 0 {
		session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("sorry, %s, but i cant seem to find your question :(", user),
			},
		})

	}

	chosen := ballAnswers[rand.Intn(len(ballAnswers))]

	session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("I asked the magical 8 ball, and it said yes %s! when i asked it \"%s\" ", chosen, question),
		},
	})
}
