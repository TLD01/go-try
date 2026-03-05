package altitude_unit

import (
	"aerowatch.com/api/constants"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type AltitudeUnit struct {
	constants.StringConstant
	Unit        string
	Description string
}


func (a AltitudeUnit) Equal(other AltitudeUnit) bool {
	return a.Name() == other.Name()
}

// These vars are intended as read-only constants. Do not mutate.
var (
	// GND_SFC - Ground/Surface, mark- och/eller vattenyta
	GND_SFC = AltitudeUnit{constants.NewStringConstant("GND/SFC"), "GND/SFC", "Ground/Surface, mark- och/eller vattenyta"}

	// FL - Flygnivå/Flightlevel, anges i hundra tals fot, d v s FL330=33000 ft
	FL = AltitudeUnit{constants.NewStringConstant("FL"), "FL", "Flygnivå/Flightlevel, anges i hundra tals fot"}

	// UNL - Obegränsad/Unlimited
	UNL = AltitudeUnit{constants.NewStringConstant("UNL"), "UNL", "Obegränsad/Unlimited"}

	// FT_AMSL - Feet above mean sea level (default when no unit is specified)
	FT_AMSL = AltitudeUnit{constants.NewStringConstant("FT_AMSL"), "-", "Feet above mean sea level (när ingen enhet anges)"}
)


func All() []AltitudeUnit {
	return []AltitudeUnit{GND_SFC, FL, UNL, FT_AMSL}
}

func (a *AltitudeUnit) UnmarshalJSON(data []byte) error {
	unit, err := constants.UnmarshalJSON(All(), data)
	if err != nil {
		return err
	}
	*a = *unit
	return nil
}


func (a *AltitudeUnit) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	unit, err := constants.UnmarshalBSONValue(All(), t, data)
	if err != nil {
		return err
	}
	*a = *unit
	return nil
}