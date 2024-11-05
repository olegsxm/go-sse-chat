package services

import (
	"chat/internal/models"
	"chat/internal/pkg/jwt"
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"log/slog"
	"time"
)

type userRepository interface {
	CreateUser(ctx context.Context, login string, hash string, salt []byte) (models.User, error)
	FindUserByLogin(ctx context.Context, login string) (models.User, error)
}

type AuthService struct {
	userRepository userRepository
	passPepper     string
	jwtSecret      string
}

func (a AuthService) SignUp(ctx context.Context, login, password string) (models.AuthResponse, string, error) {
	user, err := a.createUser(ctx, login, password, a.passPepper)
	if err != nil {
		return models.AuthResponse{}, "", err
	}

	token, refresh, err := a.createTokens(user, a.jwtSecret)
	if err != nil {
		return models.AuthResponse{}, "", err
	}

	response := models.AuthResponse{
		Token: token,
		User:  user,
	}

	return response, refresh, nil
}

func (a AuthService) SignIn(ctx context.Context, login, password string) (models.AuthResponse, string, error) {
	u, err := a.userRepository.FindUserByLogin(ctx, login)
	if err != nil {
		slog.Error("FindUserByLogin error", err.Error())
		return models.AuthResponse{}, "", errors.New("internal server error") // TODO add errors
	}

	if a.hashPassword(password, a.passPepper, u.Salt) != u.Password {
		return models.AuthResponse{}, "", errors.New("invalid credentials")
	}

	token, refresh, err := a.createTokens(models.UserDTO{Id: u.Id.String(), Login: u.Login}, a.jwtSecret)
	if err != nil {
		slog.Error("CreateTokens error", err.Error())
		return models.AuthResponse{}, "", err
	}

	response := models.AuthResponse{
		Token: token,
		User: models.UserDTO{
			Id:    u.Id.String(),
			Login: u.Login,
		},
	}

	return response, refresh, nil
}

func (a AuthService) createUser(ctx context.Context, login, password, pepper string) (models.UserDTO, error) {
	salt, err := a.generatePasswordSalt()
	if err != nil {
		slog.Error(err.Error())
		return models.UserDTO{}, errors.New("failed to generate password salt")
	}

	hash := a.hashPassword(password, pepper, salt)

	user, err := a.userRepository.CreateUser(ctx, login, hash, salt)
	if err != nil {
		slog.Error(err.Error())
		return models.UserDTO{}, errors.New("failed to create user")
	}

	return user.ToDTO(), nil
}

func (a AuthService) createTokens(user models.UserDTO, jwtSecret string) (string, string, error) {
	tokenClaims := &jwt.ChatClaims{
		Id:    user.Id,
		Login: user.Login,
		Key:   "token",
	}
	tokenClaims.ExpiresAt = jwt2.NewNumericDate(time.Now().Add(time.Hour * 1))
	token, err := jwt.CreateToken(tokenClaims, jwtSecret)
	if err != nil {
		return "", "", err
	}

	refresh := &jwt.ChatClaims{
		Id:    user.Id,
		Login: user.Login,
		Key:   "refresh",
	}
	refresh.ExpiresAt = jwt2.NewNumericDate(time.Now().Add(time.Hour * 24))
	refreshToken, err := jwt.CreateToken(refresh, jwtSecret)

	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

// Password util
func (a AuthService) generatePasswordSalt() ([]byte, error) {
	saltSize := 16
	salt := make([]byte, saltSize)

	_, e := rand.Read(salt)
	if e != nil {
		return nil, e
	}

	return salt, nil
}

func (a AuthService) hashPassword(password, pepper string, salt []byte) string {
	passBytes := []byte(password)
	pepperBytes := []byte(pepper)

	hasher := sha512.New()

	resBytes := make([]byte, len(passBytes)+len(pepperBytes)+len(salt))

	resBytes = append(resBytes, pepperBytes...)
	resBytes = append(resBytes, passBytes...)
	resBytes = append(resBytes, salt...)

	hasher.Write(resBytes)
	hashedPasswordBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashedPasswordBytes)
}

func (a AuthService) checkPassword(password, hashedPassword, pepper string, salt []byte) bool {
	hash := a.hashPassword(password, pepper, salt)
	return hashedPassword == hash
}

// Constructor
func newAuthService(ur userRepository, jwtSecret, pepper string) AuthService {
	return AuthService{
		userRepository: ur,
		passPepper:     pepper,
		jwtSecret:      jwtSecret,
	}
}
