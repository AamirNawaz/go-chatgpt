package models

import "time"

type User struct {
	Id            uint             `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Name          string           `json:"name"`
	Email         string           `json:"email" gorm:"unique"`
	Password      []byte           `json:"password"`
	Status        string           `json:"status"`
	SearchHistory []*SearchHistory `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
