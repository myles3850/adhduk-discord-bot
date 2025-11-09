package main

import (
	discordapi "choccobear.tech/emojiBot/discordApi"
	webapi "choccobear.tech/emojiBot/webApi"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	discord, err := discordapi.Setup()
	if err != nil {
		panic(1)
	}
	defer discord.Session.Close()
	discord.Session.AddHandler(discord.OnInteraction)


	web := webapi.Setup()
	server := web.Gin
	registerApiEndpoints(web.Gin, discord, web)

	server.Run()

}

func registerApiEndpoints(server *gin.Engine, discord *discordapi.Discord, web *webapi.WebCtx) {
	server.GET("emoji", func(ctx *gin.Context) {
		emojis := discord.GetAllEmojis()
		web.GetAllEmojis(ctx, emojis)
	})

	server.POST("emoji/:id/role", func(ctx *gin.Context) {
		id := ctx.Param("id")
		discord.EditEmojiRoles(id, web.UpdateEmojiRoles(ctx))
	})

	server.GET("role", func(ctx *gin.Context) {
		web.GetAllRoles(ctx, discord.GetAllRoles())
	})
}
