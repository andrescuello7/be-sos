package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model

	FullName string `json:"fullname"`
	Message  string `gorm:"not null" json:"message"`
}
