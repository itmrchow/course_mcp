package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/mark3labs/mcp-go/client/transport"
	"github.com/spf13/viper"
)

var _ AuthProvider = &KeycloakAuthProvider{}

type KeycloakAuthProvider struct {
	KeycloakURL  string
	Realm        string
	ClientID     string
	ClientSecret string
}

func NewKeycloakAuthProvider() *KeycloakAuthProvider {
	return &KeycloakAuthProvider{
		KeycloakURL:  viper.GetString("KEYCLOAK_URL"),
		Realm:        viper.GetString("KEYCLOAK_REALM"),
		ClientID:     viper.GetString("CLIENT_ID"),
		ClientSecret: viper.GetString("CLIENT_SECRET"),
	}
}

func (k *KeycloakAuthProvider) GetAuthorizeURL(clientID string, codeChallenge string, state string, redirectURI string, responseType string, scopes string) (string, error) {
	authorizeURL := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/auth", k.KeycloakURL, k.Realm)
	u, err := url.Parse(authorizeURL)
	if err != nil {
		return "", err
	}
	values := url.Values{}
	values.Set("client_id", clientID)
	values.Set("code_challenge", codeChallenge)
	values.Set("response_type", responseType)
	if state != "" {
		values.Set("state", state)
	}
	if redirectURI != "" {
		values.Set("redirect_uri", redirectURI)
	}
	if scopes != "" {
		values.Set("scope", scopes)
	}
	u.RawQuery = values.Encode()
	return u.String(), nil

}

func (k *KeycloakAuthProvider) ExchangeToken(authorization string, tokenReq *TokenReq) (*Token, error) {
	tokenEndpoint := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", k.KeycloakURL, k.Realm)

	// 建立 form-encoded 請求體
	formData := url.Values{}
	formData.Set("grant_type", tokenReq.GrantType)
	formData.Set("redirect_uri", tokenReq.RedirectURI)
	formData.Set("code", tokenReq.Code)
	formData.Set("code_verifier", tokenReq.CodeVerifier)

	// formData.Set("client_id", tokenReq.ClientID)
	// formData.Set("client_secret", tokenReq.ClientSecret)

	// 建立請求
	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	// 設定 Basic Authentication
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 印出請求內容
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Printf("Error dumping request: %v", err)
	} else {
		log.Printf("Request dump:\n%s", string(dump))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token exchange failed: %s", string(body))
	}

	var tokenResp transport.Token
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		TokenType:    tokenResp.TokenType,
		ExpiresIn:    tokenResp.ExpiresIn,
		Scope:        tokenResp.Scope,
		// ExpiresAt is the time when the token expires
		ExpiresAt: tokenResp.ExpiresAt,
	}, nil
}

func (k *KeycloakAuthProvider) FetchUserInfo(accessToken string) (map[string]interface{}, error) {
	userInfoEndpoint := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/userinfo", k.KeycloakURL, k.Realm)
	req, err := http.NewRequest("GET", userInfoEndpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 直接從 response body 解析 JSON
	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return userInfo, nil
}

func (k *KeycloakAuthProvider) GetConfiguration() (*transport.AuthServerMetadata, error) {
	configEndpoint := fmt.Sprintf("%s/realms/%s/.well-known/openid-configuration", k.KeycloakURL, k.Realm)
	// http://localhost:8081/realms/course_server/.well-known/openid-configuration
	resp, err := http.Get(configEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error fetching configuration: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching configuration: %v", err)
	}

	var metadata transport.AuthServerMetadata
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return nil, fmt.Errorf("error decoding configuration JSON: %v", err)
	}

	return &metadata, nil
}

type TokenReq struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
	CodeVerifier string `json:"code_verifier"`
	RedirectURI  string `json:"redirect_uri"`
}
