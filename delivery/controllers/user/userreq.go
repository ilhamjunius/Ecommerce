package user

type RegisterUserRequestFormat struct {
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	Name            string `json:"name" form:"name"`
	HandphoneNumber string `json:"handphonenumber" form:"handphonenumber"`
}

type PutUserRequestFormat struct {
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	Name            string `json:"name" form:"name"`
	HandphoneNumber string `json:"handphonenumber" form:"handphonenumber"`
}
