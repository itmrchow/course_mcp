package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/rs/zerolog"

	"course-mcp/internal/usecase/utils"
)

type OAuthMatadataHandler struct {
	logger       *zerolog.Logger
	authProvider utils.AuthProvider
}

func NewOAuthMatadataHandler(logger *zerolog.Logger, authProvider utils.AuthProvider) *OAuthMatadataHandler {
	return &OAuthMatadataHandler{
		logger:       logger,
		authProvider: authProvider,
	}
}

func (h *OAuthMatadataHandler) HandleOAuthProtectedResource(c *gin.Context) {
	metadata := &transport.OAuthProtectedResource{
		AuthorizationServers: []string{"http://localhost:3000/realms/course_server/.well-known/openid-configuration"},
		Resource:             "Example OAuth Protected Resource",
		ResourceName:         "Example OAuth Protected Resource",
	}
	c.JSON(http.StatusOK, metadata)
}

func (h *OAuthMatadataHandler) HandleOAuthAuthorizationServer(c *gin.Context) {

	// TODO: call authserver to get metadata

	metadata, err := h.authProvider.GetConfiguration()
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get OAuth configuration")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get OAuth configuration"})
		return
	}

	// change endpoint URLs to localhost:3000
	metadata.AuthorizationEndpoint = "http://localhost:3000/authorize"
	metadata.TokenEndpoint = "http://localhost:3000/token"
	metadata.RegistrationEndpoint = "http://localhost:3000/register"
	metadata.JwksURI = "http://localhost:3000/jwks"

	c.JSON(http.StatusOK, metadata)
}
