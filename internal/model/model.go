package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents an app user
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"` // bcrypt hashed
	Name     string `json:"name"`

	Roles []Role `gorm:"many2many:user_roles" json:"roles"`
}

// Role represents a role like admin, user, manager
type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`

	Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions"`
}

type Permission struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
}

type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}

type OAuthClient struct {
	ID        uint   `gorm:"primaryKey"`
	ClientID  string `gorm:"uniqueIndex;not null"`
	Secret    string `gorm:"not null"`
	Name      string
	CreatedAt time.Time
}
