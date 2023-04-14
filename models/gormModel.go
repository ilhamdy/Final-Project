package models

import "time"

type GormModel struct {
	ID        uint      `json:"id" gorm:"primaryKey:autoIncrement"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
}
