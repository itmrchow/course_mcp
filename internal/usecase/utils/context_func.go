package utils

import "context"

// context key
type TokenClaimsKey struct{}

// WithFunc

// WithTokenClaims: Sets the token claims in the context for further use.
func WithTokenClaims(ctx context.Context, claims *TokenClaims) context.Context {
	return context.WithValue(ctx, TokenClaimsKey{}, claims)
}

// GetFunc

// GetTokenClaims: Retrieves the token claims from the context.
func GetTokenClaims(ctx context.Context) (*TokenClaims, bool) {
	claims, ok := ctx.Value(TokenClaimsKey{}).(*TokenClaims)
	return claims, ok
}
