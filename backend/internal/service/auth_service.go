package service

import (
	"errors"

	"shortavee/backend/internal/model"
	"shortavee/backend/internal/repository"
	"shortavee/backend/pkg/utils"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// Register logic
func (s *AuthService) Register(email, password string) (*model.User, error) {

	// cek user sudah ada
	_, err := s.userRepo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:       uuid.New(),
		Email:    email,
		Password: hashedPassword,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login logic
func (s *AuthService) Login(email, password string) (string, error) {

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// cek password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	// generate JWT
	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) GetUserByID(id uuid.UUID) (*model.User, error) {
	return s.userRepo.FindByID(id)
}
