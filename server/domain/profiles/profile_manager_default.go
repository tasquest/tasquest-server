package profiles

import (
	"emperror.dev/errors"
	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"tasquest.com/server/security"
)

var profileManagerOnce sync.Once
var profileManagerInstance *DefaultProfileManager

type DefaultProfileManager struct {
	ProfileRepository ProfileRepository
	UserManagement    security.UserManagement
}

func ProvideDefaultProfileManager(repository ProfileRepository, management security.UserManagement) *DefaultProfileManager {
	profileManagerOnce.Do(func() {
		profileManagerInstance = &DefaultProfileManager{
			ProfileRepository: repository,
			UserManagement:    management,
		}
	})
	return profileManagerInstance
}

func (pm *DefaultProfileManager) CreateProfile(command CreateUserProfile) (Profile, error) {
	_, usrErr := pm.UserManagement.FetchUser(command.UserID)

	if usrErr != nil {
		return Profile{}, errors.WithStack(usrErr)
	}

	existingProfile, profErr := pm.ProfileRepository.FindByUser(command.UserID)

	if !cmp.Equal(existingProfile, Profile{}) {
		return Profile{}, errors.WithStack(ErrProfileAlreadyExists)
	}

	if profErr != nil {
		return Profile{}, errors.Wrap(profErr, "Failed to fetch profile")
	}

	userID, _ := primitive.ObjectIDFromHex(command.UserID)

	profile := Profile{
		UserID:   userID,
		Name:     command.Name,
		Surname:  command.Surname,
		Birthday: primitive.NewDateTimeFromTime(command.Birthday),
	}

	return pm.ProfileRepository.Save(profile)
}
