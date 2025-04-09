package usecase

import (
	"encoding/json"
)

func FormatJSON(input []byte) (string, error) {
	var obj interface{}
	if err := json.Unmarshal(input, &obj); err != nil {
		return "", err
	}

	out, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}

	return string(out), nil
}
