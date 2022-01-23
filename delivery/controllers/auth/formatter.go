package auth

type LoginRequestFormat struct {
	Email    string `json:"email" form:"password"`
	Password string `json:"password" form:"password"`
}

type LoginResponseFormat struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
