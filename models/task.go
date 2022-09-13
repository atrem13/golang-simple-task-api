package models

import "time"

type Task struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	AssingedTo string    `json:"id" gorm:"assingedTo"`
	Task       string    `json:"id" gorm:"task"`
	Deadline   time.Time `json:"id" gorm:"deadline"`
	CreatedAt  time.Time `json:"id" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `json:"id" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
