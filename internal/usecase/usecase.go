package usecase

import (
	"strconv"
	"time"

	"github.com/fandyputram/go-project-template/internal/entity"
	"github.com/fandyputram/go-project-template/internal/repository"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key")

type Usecase interface {
	GetUser(id string) (*entity.User, error)
	Login(username, password string) (string, error)
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

func (u *usecase) Login(username, password string) (string, error) {
	user, err := u.userRepo.GetUserByCredentials(username, password)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   user.Username,
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
