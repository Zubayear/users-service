package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	Name       string `json:"artist_name"`
}