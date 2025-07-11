package handlers

import (
	"log/slog"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"course-mcp/internal/usecase/utils"
)

type TokenHandler struct {
	logger       *zerolog.Logger
	authProvider utils.AuthProvider
}

func NewTokenHandler(logger *zerolog.Logger, authProvider utils.AuthProvider) *TokenHandler {
	return &TokenHandler{
		logger:       logger,
		authProvider: authProvider,
	}
}

func (h *TokenHandler) HandleToken(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	grantType := c.PostForm("grant_type")
	code := c.PostForm("code")
	codeVerifier := c.PostForm("code_verifier")
	// clientIDParam := c.PostForm("client_id")
	redirectURI := c.PostForm("redirect_uri")
	// clientSecret := c.PostForm("client_secret")

	// h.logger.Info().Msgf("Received token request: authorization=%s, grant_type=%s, code=%s, client_id=%s, redirect_uri=%s",
	// 	authorization,
	// 	grantType, code, clientIDParam, redirectURI)

	if grantType != "authorization_code" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported grant_type"})
		return
	}
	if code == "" || redirectURI == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code, client_id, and redirect_uri are required"})
		return
	}

	tokenReq := &utils.TokenReq{
		GrantType:    grantType,
		Code:         code,
		CodeVerifier: codeVerifier,
		RedirectURI:  redirectURI,
	}

	token, err := h.authProvider.ExchangeToken(authorization, tokenReq)
	if err != nil {
		slog.Error("Token exchange failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if token == nil {
		slog.Error("Token exchange returned nil token without error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "empty token response"})
		return
	}

	accessToken := token.AccessToken

	userInfo, userErr := h.authProvider.FetchUserInfo(accessToken)
	if userErr != nil {
		slog.Error("Failed to fetch user info", "error", userErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user info", "details": userErr.Error()})
		return
	}

	slog.Info("Token response with user info",
		"email", userInfo["email"],
		"login", userInfo["login"],
	)

	c.JSON(http.StatusOK, token)
}

func (h *TokenHandler) HandleAuthorize(c *gin.Context) {
	queryString := c.Request.URL.RawQuery
	h.logger.Info().Msgf("Authorize request query parameters: %s", queryString)

	authURL, err := url.Parse("http://localhost:8081/realms/course_server/protocol/openid-connect/auth")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	authURL.RawQuery = queryString

	c.Redirect(http.StatusFound, authURL.String())
}
