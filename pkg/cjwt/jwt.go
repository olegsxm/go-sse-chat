package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserClaims struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func NewEchoJwtConfig(secret string) echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(UserClaims)
		},
		SigningKey: []byte(secret),
	}
}

func CreateToken(claims *UserClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", nil
	}

	return t, nil
}

func Parse() {}
