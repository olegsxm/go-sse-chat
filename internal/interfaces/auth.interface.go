package interfaces

import "github.com/olegsxm/go-sse-chat.git/internal/models"

// Auth TODO think about abstraction
type IAuth interface {
	SignIn(login, password string) (models.AuthResponse, error)
}
