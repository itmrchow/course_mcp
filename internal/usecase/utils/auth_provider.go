package utils

import (
	"time"

	"github.com/mark3labs/mcp-go/client/transport"
)

type AuthConfig struct {
	BaseURL      string `json:"base_url"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	Realm        string `json:"realm"`
}

type AuthProvider interface {
	GetAuthorizeURL(clientID, codeChallenge, state, redirectURI, responseType, scopes string) (string, error)

	ExchangeToken(authorization string, tokenReq *TokenReq) (*Token, error)

	FetchUserInfo(accessToken string) (map[string]interface{}, error)

	GetConfiguration() (*transport.AuthServerMetadata, error)
}

type Token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"` // Optional, may not be present in all responses
	TokenType    string    `json:"token_type"`              // e.g., "Bearer"
	ExpiresIn    int64     `json:"expires_in,omitempty"`    // Duration in seconds
	Scope        string    `json:"scope,omitempty"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
}
