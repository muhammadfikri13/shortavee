package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URL struct {
	ID     uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	UserID *uuid.UUID `json:"user_id" gorm:"type:uuid;index"`

	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code" gorm:"unique"`
	ClickCount  int       `json:"click_count"`
	CreatedAt   time.Time `json:"created_at"`
}

func (u *URL) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
