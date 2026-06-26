package service

import (
	"shortavee/backend/internal/model"
	"shortavee/backend/internal/repository"
	"shortavee/backend/pkg/utils"

	"github.com/google/uuid"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) CreateShortURL(originalURL string, userID uuid.UUID) (*model.URL, error) {

	url := &model.URL{
		OriginalURL: originalURL,
		ShortCode:   utils.GenerateShortCode(6),
		UserID:      &userID,
	}

	err := s.repo.Create(url)

	return url, err
}

func (s *URLService) GetOriginalURL(code string) (*model.URL, error) {
	return s.repo.FindByShortCode(code)
}

func (s *URLService) IncrementClickCount(id uuid.UUID) error {
	return s.repo.IncrementClickCount(id)
}

func (s *URLService) GetURLsByUser(userID uuid.UUID) ([]model.URL, error) {
	return s.repo.FindAllByUserID(userID)
}
