package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"course-mcp/internal/usecase/utils"
)

type OAuthMetadataHandler struct {
	logger       *zerolog.Logger
	authProvider utils.AuthProvider
}

func NewOAuthMetadataHandler(logger *zerolog.Logger, authProvider utils.AuthProvider) *OAuthMetadataHandler {
	return &OAuthMetadataHandler{
		logger:       logger,
		authProvider: authProvider,
	}
}

func (h *OAuthMetadataHandler) HandleOAuthProtectedResource(c *gin.Context) {
	authorizationServers := fmt.Sprintf("%s:%s/realms/course_server/.well-known/openid-configuration", viper.GetString("APP_URL"), viper.GetString("PORT"))

	metadata := &transport.OAuthProtectedResource{
		// TODO: url
		AuthorizationServers: []string{authorizationServers},
		Resource:             "Example OAuth Protected Resource",
		ResourceName:         "Example OAuth Protected Resource",
	}
	c.JSON(http.StatusOK, metadata)
}

func (h *OAuthMetadataHandler) HandleOAuthAuthorizationServer(c *gin.Context) {

	// TODO: call authserver to get metadata

	metadata, err := h.authProvider.GetConfiguration()
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get OAuth configuration")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get OAuth configuration"})
		return
	}

	// change endpoint URLs to localhost:3000
	// TODO: url
	metadata.AuthorizationEndpoint = "http://localhost:3000/authorize"
	metadata.TokenEndpoint = "http://localhost:3000/token"
	metadata.RegistrationEndpoint = "http://localhost:3000/register"
	metadata.JwksURI = "http://localhost:3000/jwks"

	c.JSON(http.StatusOK, metadata)
}
