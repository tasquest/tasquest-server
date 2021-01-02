package profiles

type ProfileManager interface {
	createProfile(command CreateUserProfile) (Profile, error)
}
