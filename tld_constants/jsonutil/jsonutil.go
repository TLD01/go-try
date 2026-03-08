package jsonutil

import (
	"bytes"
	"encoding/json"
)

// JsonSerialize converts a value to a JSON string.
//
// Example:
//
//	str, err := JsonSerialize(map[string]int{"count": 42})
//	// str = `{"count":42}`
func JsonSerialize[T any](v T) (string, error) {	
	return Marshal(v)
}


func Marshal(v any) (string, error) {
	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)

	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	
	err := enc.Encode(v)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// JsonDeserialize parses a JSON string into a value.
//
// Example:
//
//	var result map[string]int
//	err := JsonDeserialize(`{"count":42}`, &result)
//	// result = map[string]int{"count":42}
func JsonDeserialize[T any](data string, v *T) error {
	return json.Unmarshal([]byte(data), v)
}
