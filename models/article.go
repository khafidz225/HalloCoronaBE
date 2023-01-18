package models

import "time"

type Article struct {
	ID          int          `json:"id" gorm:"primary_key:auto_increment"`
	Title       string       `json:"title" form:"title" gorm:"type: varchar(255)"`
	Image       string       `json:"image" form:"image" gorm:"type: varchar(255)"`
	Description string       `json:"description" form:"description" gorm:"type: varchar(255)"`
	User        UserResponse `json:"user"`
	CreatedAt   time.Time    `json:"CreatedAt"`
	UserID      int          `json:"-" `
}

func (Article) TableName() string {
	return "articles"
}
