package common

import (
	"encoding/json"
)


type Constant[T any] interface {
	String() string
	Name() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	All() []T
}

type StringConstant struct {
	nameField 	  string
}

func NewStringConstant(name string) StringConstant {
	return StringConstant{nameField: name}
}


func (s StringConstant) String() string {
	return s.Name()
}

func (s StringConstant) Name() string {
	return s.nameField
}




func (s StringConstant) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())	
}

func (s *StringConstant) UnmarshalJSON(data []byte) error {
    var name string
    if err := json.Unmarshal(data, &name); err != nil {
        return err
    }
    s.nameField = name
    return nil
}