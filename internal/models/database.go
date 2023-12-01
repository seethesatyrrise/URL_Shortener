package models

import "gorm.io/gorm"

type DBData struct {
	gorm.Model
	Link  string
	Token string
}
