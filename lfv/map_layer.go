package lfv

import (
	"encoding/json"
	"fmt"
)

type MapLayer struct {
	Name           string `json:"-" bson:"-"`
	LayerName      string `json:"layerName" bson:"layerName"`
	Description    string `json:"description" bson:"description"`
	FilterCriteria string `json:"filterCriteria" bson:"filterCriteria"`
	Enabled        bool   `json:"enabled" bson:"enabled"`
}

func (m MapLayer) String() string {
	return m.Name
}

func (m MapLayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Name)
}

func (m *MapLayer) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	for _, layer := range AllMapLayers {
		if layer.Name == name {
			*m = layer
			return nil
		}
	}
	return fmt.Errorf("unknown MapLayer: %s", name)
}

var (
	RSTA = MapLayer{Name: "RSTA", LayerName: "mais:RSTA", Description: "Restriktionsområden", FilterCriteria: "{LOWER:'GND/SFC}'", Enabled: true}

	// Farliga områden (Dangerous areas)
	DNGA = MapLayer{Name: "DNGA", LayerName: "mais:DNGA", Description: "Farliga områden", FilterCriteria: "{LOWER:'GND'}", Enabled: true}

	// Trafikzoner (Traffic zones)
	ATZ = MapLayer{Name: "ATZ", LayerName: "mais:ATZ", Description: "Trafikzoner", FilterCriteria: "{LOWER:'GND/SFC'}", Enabled: true}

	// Trafikinformationszoner (Traffic information zones)
	TIZ = MapLayer{Name: "TIZ", LayerName: "mais:TIZ", Description: "Trafikinformationszoner", FilterCriteria: "{LOWER:'GND/SFC'}", Enabled: true}

	// Kontrollzoner (Control zones)
	CTR = MapLayer{Name: "CTR", LayerName: "mais:CTR", Description: "Kontrollzoner", FilterCriteria: "{LOWER:'GND/SFC'}", Enabled: true}

	// Flygplatser (Airports)
	ARP = MapLayer{Name: "ARP", LayerName: "mais:ARP", Description: "Flygplatser", FilterCriteria: "{LOWER:'GND/SFC'}", Enabled: true}

	// Helikopterflygplatser (Helicopter airports)
	HKP_ARP = MapLayer{Name: "HKP_ARP", LayerName: "mais:HKP_ARP", Description: "Helikopterflygplatser", FilterCriteria: "{LOWER:'GND/SFC'}", Enabled: true}

	// 1 km radie runt helikopterflygplatser (1 km radius around helicopter airports)
	HKP1K = MapLayer{Name: "HKP1K", LayerName: "DAIM_TOPO:HKP1K", Description: "1 km radie runt helikopterflygplatser", FilterCriteria: "{LOWER:'GND/SFC'}", Enabled: true}

	// 5 km radie runt flygplatser (5 km radius around airports)
	RWY5K = MapLayer{Name: "RWY5K", LayerName: "DAIM_TOPO:RWY5K", Description: "5 km radie runt flygplatser", FilterCriteria: "{LOWER:'GND/SFC'}", Enabled: true}

	// AIP SUP, tillfälligt publicerade restriktionsområden (Temporary published restriction areas)
	SUP = MapLayer{Name: "SUP", LayerName: "DAIM_TOPO:SUP", Description: "AIP SUP, tillfälligt publicerade restriktionsområden", FilterCriteria: "{LOWER:'GND/SFC'}", Enabled: true}

	// NOTAM, tillfälligt publicerade restriktionsområden (Temporary published restriction areas)
	NOTAM = MapLayer{Name: "NOTAM", LayerName: "dynais:NOTAM", Description: "NOTAM, tillfälligt publicerade restriktionsområden", FilterCriteria: "{LOWER:'GND/SFC'}", Enabled: false}
)

var AllMapLayers = []MapLayer{RSTA, DNGA, ATZ, TIZ, CTR, ARP, HKP_ARP, HKP1K, RWY5K, SUP, NOTAM}
