package repository

import "github.com/fandyputram/go-project-template/internal/entity"

type UserRepository interface {
	GetUser(id int) (*entity.User, error)
	GetUserByCredentials(username, password string) (*entity.User, error)
}
