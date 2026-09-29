package jwtconfig

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Jwt struct {
	Secret       string
	TokenExpiry  time.Duration
	RefreshExpir time.Duration
}

func Load() (*Jwt, error) {
	err := godotenv.Load("jwt.env")
	cfg := &Jwt{}

	if err != nil {
		fmt.Println("Error ao carregar o arquivo .env: ", err)
		return nil, nil
	}

	cfg.Secret = getEnv("SECRET", "chave-reserva-se-o-env-sumir")
	cfg.TokenExpiry = time.Hour * 24
	cfg.RefreshExpir = time.Hour * 168

	return cfg, nil
}

func getEnv(Key, defaultValue string) string {
	if value := os.Getenv(Key); value != "" {
		return value
	}
	return defaultValue
}
