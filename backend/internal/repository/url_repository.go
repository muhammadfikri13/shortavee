package repository

import (
	"shortavee/backend/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URLRepository struct {
	DB *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{DB: db}
}

func (r *URLRepository) Create(url *model.URL) error {
	return r.DB.Create(url).Error
}

func (r *URLRepository) FindByShortCode(code string) (*model.URL, error) {

	var url model.URL

	err := r.DB.
		Where("short_code = ?", code).
		First(&url).Error

	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *URLRepository) IncrementClickCount(id uuid.UUID) error {
	return r.DB.
		Model(&model.URL{}).
		Where("id = ?", id).
		Update("click_count", gorm.Expr("click_count + 1")).
		Error
}

func (r *URLRepository) FindAllByUserID(userID uuid.UUID) ([]model.URL, error) {
	var urls []model.URL

	err := r.DB.
		Where("user_id = ?", userID).
		Find(&urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (r *URLRepository) FindAll() ([]model.URL, error) {
	var urls []model.URL

	err := r.DB.Find(&urls).Error
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (r *URLRepository) DeleteByID(id string, userID string) error {
	return r.DB.
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&model.URL{}).Error
}
