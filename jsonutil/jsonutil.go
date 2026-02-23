package jsonutil

import (
	"encoding/json"
)

// JsonSerialize converts a value to a JSON string.
//
// Example:
//
//	str, err := JsonSerialize(map[string]int{"count": 42})
//	// str = `{"count":42}`
func JsonSerialize(v any) (string, error) {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// JsonDeserialize parses a JSON string into a value.
//
// Example:
//
//	var result map[string]int
//	err := JsonDeserialize(`{"count":42}`, &result)
//	// result = map[string]int{"count":42}
func JsonDeserialize(data string, v any) error {
	return json.Unmarshal([]byte(data), v)
}
