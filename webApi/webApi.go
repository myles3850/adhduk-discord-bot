package webapi

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

type WebCtx struct {
	Gin *gin.Engine
}

type EmojiRoleUpdate struct {
	Roles []string `json:"RoleIds" binding:"required"`
}

func Setup() *WebCtx {
	gin := gin.Default()
	return &WebCtx{Gin: gin}
}

func (ctx WebCtx) GetAllEmojis(c *gin.Context, emojis []*discordgo.Emoji) {
	c.JSON(http.StatusOK, emojis)
}

func (ctx WebCtx) GetAllRoles(c *gin.Context, roles []*discordgo.Role) {
	c.JSON(http.StatusOK, roles)
}

func (ctx WebCtx) UpdateEmojiRoles(c *gin.Context) *discordgo.EmojiParams {
	emojiSettings := EmojiRoleUpdate{}
	c.ShouldBind(&emojiSettings)

	params := &discordgo.EmojiParams{Roles: emojiSettings.Roles}
	return params
}
