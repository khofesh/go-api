package forms

// SignupData :
type SignupData struct {
	Email                string `form:"email" json:"email" binding:"required,email"`
	Password             string `form:"password" json:"password" binding:"required"`
	PasswordConfirmation string `form:"passwordConfirmation" json:"passwordConfirmation" binding:"required"`
	Type                 string `form:"type" json:"type" binding:"required"`
	Demo                 bool   `form:"demo" json:"demo" binding:"required"`
}

// SigninData :
type SigninData struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
