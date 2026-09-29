package model

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

type User struct {
	ID           int       `json:"Id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // "-" means this won't be included in JSON
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserRegister struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *UserRegister) Validate() error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		fmt.Println("EMAIL INVÁLIDO")
		return errors.New("Formato de email inválido")
	}
	return nil
}
