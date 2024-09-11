package jwt

import (
	"context"
	"fmt"
	"time"
	
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Service struct {
	cfg *Config
	tokensStore user.Store
}

func New(cfg *Config, tokenStore user.Store) *Service {
	return &Service{
		cfg: cfg,
		tokensStore: tokenStore,
	}
}

func (s *Service) CreateAccessToken(ctx context.Context, user UserData) (AccessToken, error) {
	tokenData := TokenData{
		UserData:  user,
		TokenType: accessToken,
	}

	token, err := s.generateToken(tokenData)

	if err != nil {
		return AccessToken{}, err
	}

	return AccessToken{AccessToken: token}, nil
}

func (s *Service) CreateRefreshToken(ctx context.Context, user UserData) (RefreshToken, error) {
	tokenData := TokenData{
		UserData:  user,
		TokenType: refreshToken,
	}

	token, err := s.generateToken(tokenData)

	if err != nil {
		return RefreshToken{}, err
	}

	return RefreshToken{RefreshToken: token}, nil
}

func (s *Service) ValidateToken(tokenString string) (*UserData, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(s.cfg.Secret), nil
    })
    if err != nil {
        return nil, ErrInvalidToken
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        role := fmt.Sprint(claims["role"])
        username := fmt.Sprint(claims["username"])
        user := &UserData{
            Username: username,
            Role:     role,
        }

		if exp, ok := claims["exp"].(float64); ok {
					if int64(exp) < time.Now().Unix() {
						return user, ErrTokenExpired 
					}
		}

        return user, nil
    }

    return nil, ErrInvalidToken
}

func (s *Service) generateToken(tokenData TokenData) (string, error) {
	var claims *CustomClaims
	switch(tokenData.TokenType){
	case "access":
		claims = &CustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(s.cfg.AccessTokenExpiration)),
				IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
				ID:        uuid.New().String(),
			},
			User:      tokenData.Username,
			Role:      tokenData.Role,
			TokenType: tokenData.TokenType,
		}
		break
	case "refresh":
		claims = &CustomClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(s.cfg.RefreshTokenExpiration)),
				IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
				ID:        uuid.New().String(),
			},
			User:      tokenData.Username,
			Role:      tokenData.Role,
			TokenType: tokenData.TokenType,
		}
		break
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.cfg.Secret))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return signed, nil
}
