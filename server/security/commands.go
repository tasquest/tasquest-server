package security

type RegisterUserCommand struct {
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required,gte=12,lte=25"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required,gte=12,lte=25"`
}
