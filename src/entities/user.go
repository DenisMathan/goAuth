package entities

import (
	"database/sql"
	"time"
)

type User struct {
	ID       uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Email    string `gorm:"column:email" json:"email"`
	UserName string `gorm:"column:user" json:"userName"`
	Surame   string `gorm:"column:surname" json:"surname"`
	Name     string `gorm:"column:name" json:"name"`
	Mobile   string `gorm:"column:mobile" json:"mobile"`

	AuthService  string `gorm:"column:authService" json:"authService"`
	Locale       string `gorm:"column:locale" json:"locale"`
	Verified     bool   `gorm:"column:verified" json:"verified"`
	RefreshToken string `gorm:"column:refreshToken" json:"-"`

	CreatedAt sql.NullTime `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt *time.Time   `gorm:"column:deletedAt" json:"deletedAt"`
}

type Token struct {
	Token          string    `gorm:"column:token" json:"token"`
	UserID         uint      `gorm:"column:userId" json:"userId"`
	UserEmail      string    `gorm:"-" json"-"`
	ExpirationDate time.Time `gorm:"column:expirationDate" json:"expirationDate"`
}
