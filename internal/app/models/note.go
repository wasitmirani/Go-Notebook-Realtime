package models

import "gorm.io/gorm"

type Note struct{
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	MetaData  string `json:"meta_data"`
	Thumbnail string `json:"thumbnail"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}