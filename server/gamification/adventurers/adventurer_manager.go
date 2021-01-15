package adventurers

type AdventurerManager interface {
	CreateAdventurer(command CreateAdventurer) (Adventurer, error)
	UpdateAdventurer(adventurerID string, command UpdateAdventurer) (Adventurer, error)
	UpdateAdventurerExperience(command UpdateExperience) (Adventurer, error)
	FetchAdventurerForUser(userID string) (Adventurer, error)
}
