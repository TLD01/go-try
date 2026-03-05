package notam_type

import (
	"aerowatch.com/api/constants"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type NotamType struct {
	constants.StringConstant
	Code        string
	Description string
}

var (
	NEW          = NotamType{constants.NewStringConstant("NEW"), "N", "Ny NOTAM"}
	REPLACEMENT  = NotamType{constants.NewStringConstant("REPLACEMENT"), "R", "Replacement"}
	CANCELLATION = NotamType{constants.NewStringConstant("CANCELLATION"), "C", "Cancellation"}
)

func All() []NotamType {
	return []NotamType{NEW, REPLACEMENT, CANCELLATION}
}

func (n *NotamType) UnmarshalJSON(data []byte) error {
	notamType, err := constants.UnmarshalJSON(All(), data)
	if err != nil {
		return err
	}
	*n = *notamType
	return nil
}

func (n *NotamType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	notamType, err := constants.UnmarshalBSONValue(All(), t, data)
	if err != nil {
		return err
	}
	*n = *notamType
	return nil
}
