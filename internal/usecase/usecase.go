package usecase

import (
	"errors"
	"strconv"
	"time"

	"github.com/fandyputram/go-project-template/internal/entity"
	"github.com/fandyputram/go-project-template/internal/repository"
	"github.com/fandyputram/go-project-template/pkg/hash"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key")

type Usecase interface {
	GetUser(id string) (*entity.User, error)
	Login(username, password string) (string, error)
	Register(username, password, email string) error
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
	user, err := u.userRepo.GetUserByCredentials(username)
	if err != nil {
		return "", err
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
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

func (u *usecase) Register(username, password, email string) error {
	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		return err
	}

	user := &entity.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}

	return u.userRepo.CreateUser(user)
}
