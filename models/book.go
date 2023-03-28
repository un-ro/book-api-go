package models

import "time"

type Book struct {
	BookId    int        `gorm:"primary_key" json:"id"`
	Title     string     `gorm:"not null;unique" json:"title"`
	Author    string     `gorm:"not null" json:"author"`
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
