package models

// model
type User struct {
	ID       int    `gorm:"primaryKey"`
	Email    string `gorm:"unique;size:100;not null"`
	Name     string `gorm:"not null"`
	Lastname string `gorm:"not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:user;not null"`
}

// payloads
type UserRegisterPayload struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

type UserLoginPayload struct {
}
