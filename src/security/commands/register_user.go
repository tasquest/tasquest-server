package commands

type RegisterUserCommand struct {
	Name     string `valid:"alphanum"`
	Surname  string `valid:"alphanum"`
	Email    string `valid:"email"`
	Password string `valid:"stringlength(9|20)"`
}
