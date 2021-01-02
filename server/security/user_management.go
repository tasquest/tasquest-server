package security

type UserManagement interface {
	FetchUser(id string) (User, error)
	RegisterUser(command RegisterUserCommand) (User, error)
}
