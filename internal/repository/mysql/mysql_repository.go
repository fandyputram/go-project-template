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
	// Implement database query
	return nil, nil
}
