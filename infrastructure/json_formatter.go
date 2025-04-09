package infrastructure

import (
	"encoding/json"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Format(input []byte) (string, error) {
	var obj interface{}
	if err := json.Unmarshal(input, &obj); err != nil {
		return "", err
	}
	output, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}
