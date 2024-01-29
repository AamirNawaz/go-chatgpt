package models

import (
	"time"

	"gorm.io/gorm"
)

type SearchHistory struct {
	Id             uint   `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	SearchKeyWords string `json:"search_key_words"`
	UserID         uint   `json:"user_id" gorm:"foreignKey:User"`
	Status         string `json:"status"`
	Result         string `json:"result"`
	SearchAt       time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
