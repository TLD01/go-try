package type_of_point

import (
	constants "github.com/TLD01/tld_constants"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type TypeOfPoint struct {
	constants.StringConstant
	Code        string `json:"code" bson:"code"`
	Description string `json:"description" bson:"description"`
}

var (
	ARP = TypeOfPoint{constants.NewStringConstant("ARP"), "A", "Airport"}
	HKP = TypeOfPoint{constants.NewStringConstant("HKP"), "H", "Helipad"}
)

func All() []TypeOfPoint {
	return []TypeOfPoint{ARP, HKP}
}

func (t *TypeOfPoint) UnmarshalJSON(data []byte) error {
	pointType, err := constants.UnmarshalJSON(All(), data)
	if err != nil {
		return err
	}
	*t = *pointType
	return nil
}

func (t *TypeOfPoint) UnmarshalBSONValue(tpe bsontype.Type, data []byte) error {
	pointType, err := constants.UnmarshalBSONValue(All(), tpe, data)
	if err != nil {
		return err
	}
	*t = *pointType
	return nil
}
