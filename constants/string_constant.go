package constants

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/v2/x/bsonx/bsoncore"
)

// Constant is a generic interface for immutable named constants that support
// string representation, JSON marshalling, and equality comparison.
type Constant[T any] interface {
	// String returns the human-readable representation of the constant.
	String() string
	// Name returns the unique identifier of the constant used for marshalling and lookups.
	Name() string
	// MarshalJSON encodes the constant as a JSON string using its Name.
	MarshalJSON() ([]byte, error)
	// Equal reports whether the constant is equal to other.
	Equal(other T) bool
}

// StringConstant is a value type that implements Constant using a string name.
// It is intended to be embedded in domain-specific constant types.
type StringConstant struct {
	nameField string
}

// NewStringConstant creates a StringConstant with the given name.
func NewStringConstant(name string) StringConstant {
	return StringConstant{nameField: name}
}

func (a StringConstant) Equal(other StringConstant) bool {
	return a.Name() == other.Name()
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

func (s StringConstant) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bsontype.String, bsoncore.AppendString(nil, s.Name()), nil
}

func (s *StringConstant) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	s.nameField = name
	return nil
}

// UnmarshalJSON decodes a JSON string into one of the provided items by matching its Name.
// Returns an error if the JSON is invalid or no item matches.
func UnmarshalJSON[T interface{ Name() string }](items []T, data []byte) (*T, error) {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return nil, err
	}
	return Pick(items, name)
}

// Pick returns a pointer to the first item in items whose Name equals name.
// Returns an error if no match is found.
func Pick[T interface{ Name() string }](items []T, name string) (*T, error) {
	for i, item := range items {
		if item.Name() == name {
			return &items[i], nil
		}
	}
	return nil, fmt.Errorf("unknown %T: %q", *new(T), name)
}

// UnmarshalBSONValue decodes a BSON string value into one of the provided items by matching its Name.
// Returns an error if the BSON type is not a string or no item matches.
func UnmarshalBSONValue[T interface{ Name() string }](items []T, t bsontype.Type, data []byte) (*T, error) {
	if t != bsontype.String {
		return nil, fmt.Errorf("cannot unmarshal %s into StringConstant", t)
	}
	name, _, ok := bsoncore.ReadString(data)
	if !ok {
		return nil, fmt.Errorf("failed to read BSON string")
	}
	return Pick(items, name)
}
