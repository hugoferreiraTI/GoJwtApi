package usecases

import (
	"errors"
	"fmt"
	jwtconfig "go_jwt/config/jwt_config"
	"go_jwt/model"
	"go_jwt/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthuseCase struct {
	repo            *repository.UserRepository
	jwtSecret       []byte
	tokenExpiration time.Duration
}

func NewAuthUseCase(repo *repository.UserRepository, secret []byte, expiry time.Duration) *AuthuseCase {
	return &AuthuseCase{
		repo:            repo,
		jwtSecret:       secret,
		tokenExpiration: expiry,
	}
}

func (u *AuthuseCase) GetTokenExpiration() time.Duration {
	return u.tokenExpiration
}

// criando o usuário
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

// func de login
func (u *AuthuseCase) Login(input model.UserLogin) (string, error) {
	user, err := u.repo.GetByEmail(input.Email)
	if err != nil {
		return "", errors.New("Credenciais inválidas")
	}
	if !jwtconfig.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", errors.New("Credenciais inválidas")
	}
	return u.generateToken(user.ID, user.Email)
}

// gera o token com assinatura HS256
func (uc *AuthuseCase) generateToken(userID int, email string) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"iat":     now.Unix(),
		"exp":     now.Add(uc.tokenExpiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Print(token)
	return token.SignedString(uc.jwtSecret)
}

// atualiza o token
func (u *AuthuseCase) RefreshToken(userID int) (string, error) {
	user, err := u.repo.GetById(userID)

	if err != nil {
		return "", err
	}

	return u.generateToken(user.ID, user.Email)
}
