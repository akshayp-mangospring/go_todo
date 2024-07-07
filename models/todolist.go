package models

import (
	"time"
)

type TodoList struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Todos     []Todo    `gorm:"foreignKey:TodoListID;constraint:OnDelete:CASCADE;" json:"todos"`
}
