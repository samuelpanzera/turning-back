package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-" gorm:"not null"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
	Role      string `json:"role" gorm:"default:'user'"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}