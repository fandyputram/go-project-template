package usecase

import (
	"strconv"

	"github.com/fandyputram/go-project-template/internal/entity"
	"github.com/fandyputram/go-project-template/internal/repository"
)

type Usecase interface {
	GetUser(id string) (*entity.User, error)
}

type usecase struct {
	userRepo repository.UserRepository
}

func NewUsecase(ur repository.UserRepository) Usecase {
	return &usecase{userRepo: ur}
}

func (u *usecase) GetUser(id string) (*entity.User, error) {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return u.userRepo.GetUser(userId)
}
