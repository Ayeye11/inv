package models

import "time"

// model
type Product struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"unique;size:100;not null"`
	Description *string   `gorm:"type:text"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Stock       int       `gorm:"not null;default:0"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdateAt    time.Time `gorm:"autoUpdateTime"`

	//CategoryID int `gorm:"not null;index"`

	CreatedBy int `gorm:"not null;index"`
	UpdateBy  int `gorm:"not null;index"`
}

// payloads
type AddProductPayload struct {
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *int     `json:"stock"`
}

type UpdateProductPayload struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Stock       *int     `json:"stock"`

	//CategoryID int
}
