package models

// UserModelValidator :
type UserModelValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required"`
		Type     string `form:"type" json:"type" binding:"required"`
		Demo     string `form:"demo" json:"demo" binding:"required"`
	} `json:"user"`
	userModel UserModel
}
