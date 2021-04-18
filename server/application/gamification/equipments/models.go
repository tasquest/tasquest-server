package equipments

import (
	"github.com/google/uuid"
)

type Equipment struct {
	ID            uuid.UUID         `json:"id" bson:"_id"`
	Name          string            `json:"name" bson:"name"`
	Type          EquipmentType     `json:"type" bson:"type"`
	Location      EquipmentLocation `json:"location" bson:"location"`
	OffensiveRate int               `json:"offensiveRate" bson:"offensive_rate"`
	DefensiveRate int               `json:"defensiveRate" bson:"defensive_rate"`
}

type EquipmentType int

const (
	HELMET EquipmentType = iota
	CAP
)

func (eType EquipmentType) String() string {
	return [...]string{"Helmet", "Cap"}[eType]
}

type EquipmentLocation int

const (
	HEAD EquipmentLocation = iota
	TORSO
	RIGHT_HAND
	LEFT_HAND
	WAIST
	LEG
	FOOT
)

func (loc EquipmentLocation) String() string {
	return [...]string{"Head", "Torso", "Right Hand", "Left Hand", "Waist", "Leg", "Foot"}[loc]
}
