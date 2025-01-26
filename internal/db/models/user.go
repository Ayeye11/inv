package models

import "time"

// model
type User struct {
	ID        int       `gorm:"primaryKey"`
	Email     string    `gorm:"unique;size:100;not null"`
	Name      string    `gorm:"not null"`
	Lastname  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"default:user;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// payloads
type UserRegisterPayload struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdatePayload struct {
	Email     *string   `json:"email"`
	Name      *string   `json:"name"`
	Lastname  *string   `json:"lastname"`
	Password  *string   `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRolePayload struct {
	Role string `json:"role"`
}

// response
type ShowProfile struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// methods
func (u *User) ToShowProfile() *ShowProfile {
	return &ShowProfile{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Lastname:  u.Lastname,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}
