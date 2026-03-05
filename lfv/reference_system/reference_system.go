package reference_system

import (
	"aerowatch.com/api/constants"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type ReferenceSystem struct {
	constants.StringConstant
	Code        string
	Description string
}

var (
	EPSG_3006   = ReferenceSystem{constants.NewStringConstant("EPSG_3006"), "3006", "SWEREF99 TM"}
	EPSG_3007   = ReferenceSystem{constants.NewStringConstant("EPSG_3007"), "3007", "SWEREF99 12 00"}
	EPSG_3008   = ReferenceSystem{constants.NewStringConstant("EPSG_3008"), "3008", "SWEREF99 13 30"}
	EPSG_3009   = ReferenceSystem{constants.NewStringConstant("EPSG_3009"), "3009", "SWEREF99 15 00"}
	EPSG_3010   = ReferenceSystem{constants.NewStringConstant("EPSG_3010"), "3010", "SWEREF99 16 30"}
	EPSG_3011   = ReferenceSystem{constants.NewStringConstant("EPSG_3011"), "3011", "SWEREF99 18 00"}
	EPSG_3012   = ReferenceSystem{constants.NewStringConstant("EPSG_3012"), "3012", "SWEREF99 14 15"}
	EPSG_3013   = ReferenceSystem{constants.NewStringConstant("EPSG_3013"), "3013", "SWEREF99 15 45"}
	EPSG_3014   = ReferenceSystem{constants.NewStringConstant("EPSG_3014"), "3014", "SWEREF99 17 15"}
	EPSG_3015   = ReferenceSystem{constants.NewStringConstant("EPSG_3015"), "3015", "SWEREF99 18 45"}
	EPSG_3016   = ReferenceSystem{constants.NewStringConstant("EPSG_3016"), "3016", "SWEREF99 20 15"}
	EPSG_3017   = ReferenceSystem{constants.NewStringConstant("EPSG_3017"), "3017", "SWEREF99 21 45"}
	EPSG_3018   = ReferenceSystem{constants.NewStringConstant("EPSG_3018"), "3018", "SWEREF99 23 15"}
	EPSG_3847   = ReferenceSystem{constants.NewStringConstant("EPSG_3847"), "3847", "GDA94 / BCSG02"}
	EPSG_3857   = ReferenceSystem{constants.NewStringConstant("EPSG_3857"), "3857", "WGS 84 / Pseudo-Mercator"}
	EPSG_4258   = ReferenceSystem{constants.NewStringConstant("EPSG_4258"), "4258", "ETRS89"}
	EPSG_4326   = ReferenceSystem{constants.NewStringConstant("EPSG_4326"), "4326", "WGS 84"}
	EPSG_4619   = ReferenceSystem{constants.NewStringConstant("EPSG_4619"), "4619", "SWEREF99"}
	EPSG_4976   = ReferenceSystem{constants.NewStringConstant("EPSG_4976"), "4976", "SWEREF99 (3D)"}
	EPSG_4977   = ReferenceSystem{constants.NewStringConstant("EPSG_4977"), "4977", "WGS 84 (3D)"}
	EPSG_6864   = ReferenceSystem{constants.NewStringConstant("EPSG_6864"), "6864", "SWEREF99 / RT90 2.5 gon V emulation"}
	EPSG_32633  = ReferenceSystem{constants.NewStringConstant("EPSG_32633"), "32633", "WGS 84 / UTM zone 33N"}
	EPSG_900913 = ReferenceSystem{constants.NewStringConstant("EPSG_900913"), "900913", "Google Maps Global Mercator (deprecated)"}
)

func All() []ReferenceSystem {
	return []ReferenceSystem{
		EPSG_3006, EPSG_3007, EPSG_3008, EPSG_3009, EPSG_3010, EPSG_3011,
		EPSG_3012, EPSG_3013, EPSG_3014, EPSG_3015, EPSG_3016, EPSG_3017,
		EPSG_3018, EPSG_3847, EPSG_3857, EPSG_4258, EPSG_4326, EPSG_4619,
		EPSG_4976, EPSG_4977, EPSG_6864, EPSG_32633, EPSG_900913,
	}
}

func (r *ReferenceSystem) UnmarshalJSON(data []byte) error {
	refSys, err := constants.UnmarshalJSON(All(), data)
	if err != nil {
		return err
	}
	*r = *refSys
	return nil
}

func (r *ReferenceSystem) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	refSys, err := constants.UnmarshalBSONValue(All(), t, data)
	if err != nil {
		return err
	}
	*r = *refSys
	return nil
}
