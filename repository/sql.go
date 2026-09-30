package repository

import (
	"database/sql"
	"errors"
	"go_jwt/model"
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

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	query := "SELECT id, email, passwod_hash FROM users WHERE email = $1"
	err := r.Db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetById(id int) (*model.User, error) {
	var user model.User

	query := `SELECT id, email, passwod_hash FROM users WHERE id = $1`
	err := r.Db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, errors.New("Não foi possível encontrar o ID do usuário")
	}

	return &user, nil
}
