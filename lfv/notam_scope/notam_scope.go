package notam_scope

import (
	"aerowatch.com/api/constants"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type NotamScope struct {
	constants.StringConstant	
	Code        string `json:"code" bson:"code"`
	Description string `json:"description" bson:"description"`
}

func (n NotamScope) Equal(other NotamScope) bool {
	return n.Name() == other.Name()
}

var (
	Aerodrome  = NotamScope{constants.NewStringConstant("AERO"), "A", "Aerodrome"}
	NavWarning = NotamScope{constants.NewStringConstant("NAV_WARNING"), "W", "Nav Warning"}
	Enroute    = NotamScope{constants.NewStringConstant("EN_ROUTE"), "E", "Enroute"}
	Checklist  = NotamScope{constants.NewStringConstant("CHECKLIST"), "K", "Checklist"}
)

func All() []NotamScope {
	return []NotamScope{Aerodrome, NavWarning, Enroute, Checklist}
}


func (n *NotamScope) UnmarshalJSON(data []byte) error {
	scope, err := constants.UnmarshalJSON(All(), data)	
	if err != nil {
		return err
	}
	*n = *scope
	return nil
}


func (n *NotamScope) UnmarshalBSONValue(t bsontype.Type, data []byte) error {	
	scope, err := constants.UnmarshalBSONValue(All(), t, data)
	if err != nil {
		return err
	}
	*n = *scope
	return nil
}



