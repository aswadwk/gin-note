package models

import "time"

type Activity struct {
	ID          uint       `gorm:"primary_key" autoincrement:"true" index:"true" json:"id"`
	Email       string     `gorm:"not null" json:"email"`
	Title       string     `gorm:"not null" json:"title"`
	CreatedAt   time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
}

func (e *Activity) TableName() string {
	return "activities"
}
