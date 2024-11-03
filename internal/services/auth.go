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
	"github.com/spf13/viper"
	"log/slog"
	"time"
)

type userRepository interface {
	CreateUser(ctx context.Context, login string, hash string, salt []byte) (models.User, error)
	FindUserByLogin(ctx context.Context, login string) (models.User, error)
}

type AuthService struct {
	userRepository userRepository
}

func (a AuthService) SignUp(ctx context.Context, login, password string) (models.AuthResponse, string, error) {
	user, err := a.CreateUser(ctx, login, password)
	if err != nil {
		return models.AuthResponse{}, "", err
	}

	token, refresh, err := a.CreateTokens(user)
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
		slog.Error("FindUserByLogin error", err)
		return models.AuthResponse{}, "", errors.New("internal server error") // TODO add errors
	}

	if a.hashPassword(password, u.Salt) != u.Password {
		return models.AuthResponse{}, "", errors.New("invalid credentials")
	}

	token, refresh, err := a.CreateTokens(models.UserDTO{Id: u.Id.String(), Login: u.Login})
	if err != nil {
		slog.Error("CreateTokens error", err)
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

func (a AuthService) CreateUser(ctx context.Context, login, password string) (models.UserDTO, error) {
	salt, err := a.generatePasswordSalt()
	if err != nil {
		slog.Error(err.Error())
		return models.UserDTO{}, errors.New("failed to generate password salt")
	}

	hash := a.hashPassword(password, salt)

	user, err := a.userRepository.CreateUser(ctx, login, hash, salt)
	if err != nil {
		slog.Error(err.Error())
		return models.UserDTO{}, errors.New("failed to create user")
	}

	return user.ToDTO(), nil
}

func (a AuthService) CreateTokens(user models.UserDTO) (string, string, error) {
	secret := viper.GetString("security.jwt_secret")
	tokenClaims := &jwt.ChatClaims{
		Id:    user.Id,
		Login: user.Login,
	}
	tokenClaims.ExpiresAt = jwt2.NewNumericDate(time.Now().Add(time.Minute * 1))
	token, err := jwt.CreateToken(tokenClaims, secret)
	if err != nil {
		return "", "", err
	}

	refresh := &jwt.ChatClaims{
		Id:    user.Id,
		Login: user.Login,
	}
	refresh.ExpiresAt = jwt2.NewNumericDate(time.Now().Add(time.Hour * 24))
	refreshToken, err := jwt.CreateToken(refresh, secret)

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

func (a AuthService) hashPassword(password string, salt []byte) string {
	passBytes := []byte(password)
	pepperBytes := []byte(viper.GetString("security.password_pepper")) // TODO refactor

	hasher := sha512.New()

	resBytes := make([]byte, len(passBytes)+len(pepperBytes)+len(salt))

	resBytes = append(resBytes, pepperBytes...)
	resBytes = append(resBytes, passBytes...)
	resBytes = append(resBytes, salt...)

	hasher.Write(resBytes)
	hashedPasswordBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashedPasswordBytes)
}

func (a AuthService) checkPassword(password, hashedPassword string, salt []byte) bool {
	hash := a.hashPassword(password, salt)
	return hashedPassword == hash
}

// Constructor
func newAuthService(ur userRepository) AuthService {
	return AuthService{
		userRepository: ur,
	}
}
