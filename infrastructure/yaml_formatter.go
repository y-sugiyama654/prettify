package infrastructure

import (
	"errors"

	"gopkg.in/yaml.v2"
)

type YAMLFormatter struct{}

func (f *YAMLFormatter) Format(input []byte) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	var data interface{}
	err := yaml.Unmarshal(input, &data)
	if err != nil {
		return "", err
	}

	if data == nil {
		return "", errors.New("invalid YAML: empty document")
	}

	output, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
