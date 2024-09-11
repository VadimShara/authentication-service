package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

const (
	accessToken TokenType = "access_token"
	refreshToken TokenType = "refresh_token"
)

type (
	CustomClaims struct {
		jwt.RegisteredClaims
		User      string    `json:"user"`
		Role      string    `json:"role"`
		TokenType TokenType `json:"token_type"`
	}

	TokenType string

	TokenData struct {
		UserData
		TokenType TokenType
	}

	UserData struct {
		Username string
		Role     string
	}

	AccessToken struct {
		AccessToken string `json:"access_token"`
	}

	RefreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}
)