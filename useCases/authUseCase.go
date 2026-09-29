package usecases

import (
	"errors"
	jwtconfig "go_jwt/config/jwt_config"
	"go_jwt/repository"
	"time"
)

type AuthuseCase struct {
	repo            *repository.UserRepository
	jwtSecret       []byte
	tokenExpiration time.Duration
}

func NewAuthUseCase(repo *repository.UserRepository) *AuthuseCase {
	return &AuthuseCase{
		repo: repo,
	}
}

func (uc *AuthuseCase) Register(email, password string) (int, error) {
	exists, err := uc.repo.EmailExists(email)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("Email já existe")
	}

	hashedPassword, err := jwtconfig.HashPassword(password)
	if err != nil {
		return 0, err
	}

	return uc.repo.Create(email, hashedPassword)
}
