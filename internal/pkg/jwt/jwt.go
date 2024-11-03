package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ChatClaims struct {
	Id    string `json:"id"`
	Login string `json:"login"`
	Key   string `json:"key"`
	jwt.RegisteredClaims
}

func NewEchoJwtConfig(secret string) echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(ChatClaims)
		},
		SigningKey: []byte(secret),
	}
}

func CreateToken(claims *ChatClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", nil
	}

	return t, nil
}

func Parse(token string, claims *ChatClaims, secret string) error {
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return err
}
