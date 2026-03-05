package map_layer

import (
	"aerowatch.com/api/constants"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type MapLayer struct {
	constants.StringConstant
	LayerName      string 
	Description    string
	FilterCriteria string
	Enabled        bool
}


var (
	RSTA = MapLayer{constants.NewStringConstant("RSTA"), "mais:RSTA", "Restriktionsområden", "{LOWER:'GND/SFC}'", true}

	// Farliga områden (Dangerous areas)
	DNGA = MapLayer{constants.NewStringConstant("DNGA"), "mais:DNGA", "Farliga områden", "{LOWER:'GND'}", true}

	// Trafikzoner (Traffic zones)
	ATZ = MapLayer{constants.NewStringConstant("ATZ"), "mais:ATZ", "Trafikzoner", "{LOWER:'GND/SFC'}", true}


	// Trafikinformationszoner (Traffic information zones)
	TIZ = MapLayer{constants.NewStringConstant("TIZ"), "mais:TIZ", "Trafikinformationszoner", "{LOWER:'GND/SFC'}", true}

	// Kontrollzoner (Control zones)
	CTR = MapLayer{constants.NewStringConstant("CTR"), "mais:CTR", "Kontrollzoner", "{LOWER:'GND/SFC'}", true}

	// Flygplatser (Airports)
	ARP = MapLayer{constants.NewStringConstant("ARP"), "mais:ARP", "Flygplatser", "{LOWER:'GND/SFC'}", true}

	// Helikopterflygplatser (Helicopter airports)
	HKP_ARP = MapLayer{constants.NewStringConstant("HKP_ARP"), "mais:HKP_ARP", "Helikopterflygplatser", "{LOWER:'GND/SFC'}", true}

	// 1 km radie runt helikopterflygplatser (1 km radius around helicopter airports)
	HKP1K = MapLayer{constants.NewStringConstant("HKP1K"), "DAIM_TOPO:HKP1K", "1 km radie runt helikopterflygplatser", "{LOWER:'GND/SFC'}", true}

	// 5 km radie runt flygplatser (5 km radius around airports)
	RWY5K = MapLayer{constants.NewStringConstant("RWY5K"), "DAIM_TOPO:RWY5K", "5 km radie runt flygplatser", "{LOWER:'GND/SFC'}", true}

	// AIP SUP, tillfälligt publicerade restriktionsområden (Temporary published restriction areas)
	SUP = MapLayer{constants.NewStringConstant("SUP"), "DAIM_TOPO:SUP", "AIP SUP, tillfälligt publicerade restriktionsområden", "{LOWER:'GND/SFC'}", true}

	// NOTAM, tillfälligt publicerade restriktionsområden (Temporary published restriction areas)
	NOTAM = MapLayer{constants.NewStringConstant("NOTAM"), "dynais:NOTAM", "NOTAM, tillfälligt publicerade restriktionsområden", "{LOWER:'GND/SFC'}", false}
)

func All() []MapLayer {
	return []MapLayer{RSTA, DNGA, ATZ, TIZ, CTR, ARP, HKP_ARP, HKP1K, RWY5K, SUP, NOTAM}
}


func (m *MapLayer) UnmarshalJSON(data []byte) error {
	layer, err := constants.UnmarshalJSON(All(), data)
	if err != nil {
		return err
	}
	*m = *layer
	return nil
}

func (m *MapLayer) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	layer, err := constants.UnmarshalBSONValue(All(), t, data)
	if err != nil {
		return err
	}
	*m = *layer
	return nil
}