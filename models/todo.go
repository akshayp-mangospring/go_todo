package models

import (
	"time"
)

type Todo struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Completed  bool      `json:"completed"`
	TodoListID uint      `gorm:"index" json:"todo_list_id"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
