package webapi

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

type WebCtx struct {
	Gin *gin.Engine
}

func Setup() *WebCtx {
	gin := gin.Default()
	return &WebCtx{Gin: gin}
}

func (ctx *WebCtx) GetAllEmojis(c *gin.Context, emojis []*discordgo.Emoji) {
	c.JSON(http.StatusOK, emojis)
}

