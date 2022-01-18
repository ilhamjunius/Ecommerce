package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//USERID AUTO GENERATE
	ID              uint
	Email           string
	Password        [32]byte
	Name            string
	HandphoneNumber string
	Order           []Order
}
