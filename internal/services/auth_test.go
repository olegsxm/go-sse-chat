package services

import (
	"chat/internal/models"
	"context"
	"testing"
)

type TestRepository struct {
}

func (t TestRepository) CreateUser(ctx context.Context, login string, hash string, salt []byte) (models.User, error) {
	return models.User{}, nil
}

func (t TestRepository) FindUserByLogin(ctx context.Context, login string) (models.User, error) {
	return models.User{}, nil
}

var testRepository = TestRepository{}

var testService = newAuthService(testRepository, "jwt", "pepper")

func TestGenerateSalt(t *testing.T) {
	salt, err := testService.generatePasswordSalt()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(salt) != 16 {
		t.Fatalf("salt length should be 16 bytes, got %d", len(salt))
	}
}
