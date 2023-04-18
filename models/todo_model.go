package models

import "time"

type Todo struct {
	ID              uint   `gorm:"primary_key" autoincrement:"true" index:"true" json:"id"`
	ActivityGroupID uint   `gorm:"not null" json:"activity_group_id"`
	Title           string `gorm:"not null" json:"title"`
	IsActive        bool   `gorm:"not null, default:true" json:"is_active"` // false = not done, true = done
	Priority        string `gorm:"not null" json:"priority"`

	CreatedAt time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (e *Todo) TableName() string {
	return "todos"
}
