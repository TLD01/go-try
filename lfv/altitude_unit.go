package lfv

import (
	"encoding/json"
	"fmt"

	"aerowatch.com/api/common"
)

type AltitudeUnit struct {
	common.StringConstant
	Unit        string
	Description string
}

// These vars are intended as read-only constants. Do not mutate.
var (
	// GND_SFC - Ground/Surface, mark- och/eller vattenyta
	GND_SFC = AltitudeUnit{common.NewStringConstant("GND/SFC"), "GND/SFC", "Ground/Surface, mark- och/eller vattenyta"}

	// FL - Flygnivå/Flightlevel, anges i hundra tals fot, d v s FL330=33000 ft
	FL = AltitudeUnit{common.NewStringConstant("FL"), "FL", "Flygnivå/Flightlevel, anges i hundra tals fot"}

	// UNL - Obegränsad/Unlimited
	UNL = AltitudeUnit{common.NewStringConstant("UNL"), "UNL", "Obegränsad/Unlimited"}

	// FT_AMSL - Feet above mean sea level (default when no unit is specified)
	FT_AMSL = AltitudeUnit{common.NewStringConstant("FT_AMSL"), "-", "Feet above mean sea level (när ingen enhet anges)"}
)

func (a AltitudeUnit) All() []AltitudeUnit {
	return []AltitudeUnit{GND_SFC, FL, UNL, FT_AMSL}
}

func (a *AltitudeUnit) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	for _, unit := range GND_SFC.All() {
		if unit.Name() == name {
			*a = unit
			return nil
		}
	}
	return fmt.Errorf("unknown AltitudeUnit: %q", name)
}