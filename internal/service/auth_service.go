package service

import (
	"errors"

	"github.com/ahmadalaik/be-invitation/internal/dto"
	"github.com/ahmadalaik/be-invitation/internal/models"
	"github.com/ahmadalaik/be-invitation/internal/repository"
	"github.com/ahmadalaik/be-invitation/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (string, error)
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{authRepo}
}

func (s *authService) Register(req dto.RegisterRequest) error {
	existingUser, err := s.authRepo.FindByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if existingUser != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	user := models.User{
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	return s.authRepo.Create(user)
}

func (s *authService) Login(req dto.LoginRequest) (string, error) {
	userDetail, err := s.authRepo.FindByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	if userDetail == nil {
		// return "", errors.New("invalid credentials")
		return "", errors.New("user not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password))
	if err != nil {
		// return "", errors.New("invalid credentials")
		return "", errors.New("email or password is wrong")
	}

	accessToken, err := jwt.GenerateToken(userDetail.Username)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
