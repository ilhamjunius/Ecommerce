package entities

type User struct {
	// gorm.Model
	//USERID AUTO GENERATE
	ID              uint     `gorm:"primary_key:auto_increment" json:"user_id" form:"user_id"`
	Email           string   `gorm:"index:,unique" json:"email" form:"email"`
	Password        [32]byte `gorm:"not null" json:"password" form:"password"`
	Name            string   `json:"name" form:"name"`
	HandphoneNumber string   `json:"handphonenumber" form:"handphonenumber"`
	Role            string   `json:"role" form:"role"`
	Order           []Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
