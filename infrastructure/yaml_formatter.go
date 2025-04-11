package infrastructure

import "gopkg.in/yaml.v2"

type YAMLFormatter struct{}

func (f *YAMLFormatter) Format(input []byte) (string, error) {
	var data interface{}
	err := yaml.Unmarshal(input, &data)
	if err != nil {
		return "", err
	}

	output, err := yaml.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
