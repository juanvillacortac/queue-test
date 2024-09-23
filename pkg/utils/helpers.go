package utils

import (
	"bytes"
	"encoding/json"
)

// Marshals any data structure into a `map[string]any`
func ToMap(src any) (map[string]any, error) {
	var data map[string]any
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(src); err != nil {
		return data, err
	}
	err := json.Unmarshal(b.Bytes(), &data)
	return data, err
}

// Unmarshals a `map[string]any` in the requested data structure
func UnmarshalMap(src map[string]any, dst any) error {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(src); err != nil {
		return err
	}
	return json.Unmarshal(b.Bytes(), &dst)
}
