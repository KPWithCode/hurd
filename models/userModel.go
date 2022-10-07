package models

import "gorm.io/gorm"

type User struct {
	// In Gorm we need to have Model Fields Capitalised
	gorm.Model
	ID			int			`json:"ID" gorm:"primary_key"`		
	Username	string		`json:"username" `
	Email		string		`json:"email"`
	Password	string		`json:"password"`
}
