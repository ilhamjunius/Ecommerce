package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//USERID AUTO GENERATE
	Email           string
	Password        [32]byte
	Name            string
	HandphoneNumber string
}
