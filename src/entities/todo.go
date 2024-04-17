package entities

import (
	"database/sql"
	"time"
)

type Todo struct {
	Id        uint         `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	OwnerID   uint         `gorm:"column:ownerId" json:"-"`
	Name      string       `gorm:"column:name" json:"name"`
	Done      bool         `gorm:"column:done" json:"done"`
	Notes     string       `gorm:"column:notes" json:"notes"`
	CreatedAt sql.NullTime `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt sql.NullTime `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt *time.Time   `gorm:"column:deletedAt" json:"deletedAt"`
}

type Todos struct {
	List []Todo
}
