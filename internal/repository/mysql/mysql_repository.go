package mysql

import (
	"database/sql"

	"github.com/fandyputram/go-project-template/internal/entity"
)

type MySQLRepository struct {
	DB *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{DB: db}
}

func (r *MySQLRepository) GetUser(id int) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, username, password, email FROM users WHERE id=?"
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MySQLRepository) GetUserByCredentials(username string) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, username, password, email FROM users WHERE username=?"
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MySQLRepository) CreateUser(user *entity.User) error {
	query := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	_, err := r.DB.Exec(query, user.Username, user.Password, user.Email)
	return err
}
