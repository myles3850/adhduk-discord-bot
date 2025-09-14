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
	defer discord.Session.Close()
	if err != nil {
		panic(1)
	}

	web := webapi.Setup()
	server := web.Gin

	server.GET("allEmojis", func(ctx *gin.Context) {
		emojis := discord.GetAllEmojis()
		web.GetAllEmojis(ctx, emojis)
	})

	server.GET("getEmoji", func(ctx *gin.Context) {
		
	})

	server.Run()
}
