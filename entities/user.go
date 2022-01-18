package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//USERID AUTO GENERATE
	Email           string
	Password        string
	Name            string
	HandphoneNumber string
}
