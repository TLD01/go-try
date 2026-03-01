package lfv

import "encoding/json"

type NotamScope struct {
	Name        string `json:"name" bson:"name"`
	Code        string `json:"code" bson:"code"`
	Description string `json:"description" bson:"description"`
}

func (n NotamScope) String() string {
	return n.Name
}

func (n NotamScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Name)
}

func (n *NotamScope) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	for _, scope := range AllNotamScopes {
		if scope.Name == name {
			*n = scope
			return nil
		}
	}
	return nil
}

var (
	Aerodrome  = NotamScope{Name: "AERO", Code: "A", Description: "Aerodrome"}
	NavWarning = NotamScope{Name: "NAV_WARNING", Code: "W", Description: "Nav Warning"}
	Enroute    = NotamScope{Name: "EN_ROUTE", Code: "E", Description: "Enroute"}
	Checklist  = NotamScope{Name: "CHECKLIST", Code: "K", Description: "Checklist"}
)

var AllNotamScopes = []NotamScope{Aerodrome, NavWarning, Enroute, Checklist}
