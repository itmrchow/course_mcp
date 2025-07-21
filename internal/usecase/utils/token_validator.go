package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type TokenValidator struct {
	publicKey     *rsa.PublicKey
	issuer        string
	audience      string
	revokedTokens map[string]bool // 簡單的撤銷列表，實際應用中可能使用 Redis 或數據庫
}

type TokenClaims struct {
	jwt.RegisteredClaims
	Scope string `json:"scope"`
}

func NewTokenValidator() *TokenValidator {

	publicKey, err := parseKeycloakRSAPublicKey(viper.GetString("PUBLIC_KEY"))
	if err != nil {
		fmt.Printf("Failed to parse public key: %v\n", err)
	}

	return &TokenValidator{
		publicKey:     publicKey,
		issuer:        fmt.Sprintf("%s/realms/%s", viper.GetString("KEYCLOAK_URL"), viper.GetString("KEYCLOAK_REALM")),
		audience:      viper.GetString("KEYCLOAK_AUDIENCE"),
		revokedTokens: make(map[string]bool),
	}
}

// parseKeycloakRSAPublicKey 解析 Keycloak 提供的 base64 編碼 RSA 公鑰
func parseKeycloakRSAPublicKey(base64Encoded string) (*rsa.PublicKey, error) {
	// 清理字串（移除可能的空白字符）
	base64Encoded = strings.ReplaceAll(base64Encoded, " ", "")
	base64Encoded = strings.ReplaceAll(base64Encoded, "\n", "")
	base64Encoded = strings.ReplaceAll(base64Encoded, "\r", "")

	// 解碼 base64 字串為 DER 格式的二進制數據
	buf, err := base64.StdEncoding.DecodeString(base64Encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 key: %v", err)
	}

	// 使用 x509.ParsePKIXPublicKey 解析 DER 格式
	parsedKey, err := x509.ParsePKIXPublicKey(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PKIX public key: %v", err)
	}

	// 類型斷言確保是 RSA 公鑰
	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("unexpected key type %T, expected *rsa.PublicKey", parsedKey)
	}

	return publicKey, nil
}

// Validate 檢查 JWT token 的格式、簽名和聲明
func (tv *TokenValidator) Validate(tokenString string) (*TokenClaims, error) {
	// 1. 檢查 token 格式
	// - 移除 Bearer 前綴
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	if tokenString == "" {
		return nil, fmt.Errorf("empty token")
	}

	// 2. 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 驗證簽章算法
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tv.publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	// 3. 檢查 token 是否撤銷
	if tv.isTokenRevoked(tokenString) {
		return nil, fmt.Errorf("token has been revoked")
	}

	// 4. claims 驗證
	if err := tv.validateClaims(claims); err != nil {
		return nil, fmt.Errorf("claims validation failed: %v", err)
	}

	return claims, nil
}

// isTokenRevoked 檢查 token 是否已被撤銷（使用純 token 字串，不含 Bearer 前綴）
func (tv *TokenValidator) isTokenRevoked(tokenString string) bool {
	return tv.revokedTokens[tokenString]
}

// validateClaims 驗證 JWT claims
func (tv *TokenValidator) validateClaims(claims *TokenClaims) error {
	now := time.Now()

	// 1. 驗證必要的時間欄位
	if err := tv.validateTimeFields(claims, now); err != nil {
		return err
	}

	// 2. 驗證 issuer（如果配置了期望的 issuer）
	if err := tv.validateIssuer(claims); err != nil {
		return err
	}

	// 3. 驗證 audience（如果配置了期望的 audience）
	if err := tv.validateAudience(claims); err != nil {
		return err
	}

	// 4. 驗證 subject（基本檢查）
	if err := tv.validateSubject(claims); err != nil {
		return err
	}

	return nil
}

// validateTimeFields 驗證時間相關的 claims
func (tv *TokenValidator) validateTimeFields(claims *TokenClaims, now time.Time) error {
	// exp (過期時間) 是必須的
	if claims.ExpiresAt == nil {
		return fmt.Errorf("expiration time (exp) is required")
	}
	if claims.ExpiresAt.Time.Before(now) {
		return fmt.Errorf("token has expired at %v", claims.ExpiresAt.Time)
	}

	// nbf (生效時間) 是可選的，但如果存在就必須驗證
	if claims.NotBefore != nil && claims.NotBefore.Time.After(now) {
		return fmt.Errorf("token not yet valid, becomes valid at %v", claims.NotBefore.Time)
	}

	// iat (簽發時間) 是可選的，但如果存在就必須合理
	if claims.IssuedAt != nil && claims.IssuedAt.Time.After(now) {
		return fmt.Errorf("token issued in the future at %v", claims.IssuedAt.Time)
	}

	return nil
}

// validateIssuer 驗證簽發者
func (tv *TokenValidator) validateIssuer(claims *TokenClaims) error {
	// 如果我們配置了期望的 issuer，就必須驗證
	if tv.issuer != "" {
		if claims.Issuer == "" {
			return fmt.Errorf("issuer is required but not present in token")
		}
		if claims.Issuer != tv.issuer {
			return fmt.Errorf("invalid issuer: expected %s, got %s", tv.issuer, claims.Issuer)
		}
	}
	return nil
}

// validateAudience 驗證受眾
func (tv *TokenValidator) validateAudience(claims *TokenClaims) error {
	// 如果我們配置了期望的 audience，就必須驗證
	if tv.audience != "" {
		if len(claims.Audience) == 0 {
			return fmt.Errorf("audience is required but not present in token")
		}

		validAudience := false
		for _, aud := range claims.Audience {
			if aud == tv.audience {
				validAudience = true
				break
			}
		}

		if !validAudience {
			return fmt.Errorf("invalid audience: expected %s, got %v", tv.audience, claims.Audience)
		}
	}
	return nil
}

// validateSubject 驗證主體
func (tv *TokenValidator) validateSubject(claims *TokenClaims) error {
	// Subject (sub) 通常應該存在且不為空
	if claims.Subject == "" {
		return fmt.Errorf("subject (sub) is required but not present in token")
	}

	return nil
}

// RevokeToken 撤銷指定的 token（存儲時移除 Bearer 前綴）
func (tv *TokenValidator) RevokeToken(tokenString string) {
	cleanToken := strings.TrimPrefix(tokenString, "Bearer ")
	tv.revokedTokens[cleanToken] = true
}
