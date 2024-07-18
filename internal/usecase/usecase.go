package usecase

import (
	"github.com/fandyputram/go-project-template/internal/entity"
	"github.com/fandyputram/go-project-template/internal/repository"
)

type Usecase interface {
	GetUser(id int) (*entity.User, error)
}

type usecase struct {
	userRepo repository.UserRepository
}

func NewUsecase(ur repository.UserRepository) Usecase {
	return &usecase{userRepo: ur}
}

func (u *usecase) GetUser(id int) (*entity.User, error) {
	return u.userRepo.GetUser(id)
}
