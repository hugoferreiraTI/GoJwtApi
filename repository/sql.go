package repository

import (
	"database/sql"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

// verifca se o email existe
func (r *UserRepository) EmailExists(email string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	err := r.Db.QueryRow(query, email).Scan(&exists)
	return exists, err
}

// cria o usuário no banco de dados
func (r *UserRepository) Create(email, passwordHash string) (int, error) {
	var id int
	query := "INSERT INTO users (email, passwod_hash) VALUES ($1, $2) RETURNING id"
	err := r.Db.QueryRow(query, email, passwordHash).Scan(&id)
	return id, err
}
