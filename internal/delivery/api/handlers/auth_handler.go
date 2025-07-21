package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"course-mcp/internal/usecase/utils"
)

type AuthHandler struct {
	logger       *zerolog.Logger
	authProvider utils.AuthProvider
}

func NewAuthHandler(logger *zerolog.Logger, authProvider utils.AuthProvider) *AuthHandler {
	return &AuthHandler{
		logger:       logger,
		authProvider: authProvider,
	}
}

func (h *AuthHandler) HandleRegister(c *gin.Context) {
	clientID := viper.GetString("CLIENT_ID")
	clientSecret := viper.GetString("CLIENT_SECRET")

	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	body["client_id"] = clientID
	body["client_secret"] = clientSecret
	c.JSON(200, body)
}
