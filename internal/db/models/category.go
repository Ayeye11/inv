package models

import "time"

// model
type Category struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"unique;size:100;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdateAt  time.Time `gorm:"autoUpdateTime"`
	Products  []Product `gorm:"foreignKey:CategoryID"`

	CategoryID int `gorm:"not null; index"`
}
