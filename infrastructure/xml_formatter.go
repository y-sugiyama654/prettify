package infrastructure

import (
	"encoding/xml"
	"errors"
	"strings"

	"github.com/go-xmlfmt/xmlfmt"
)

type XMLFormatter struct{}

func (f *XMLFormatter) Format(input []byte) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	if strings.TrimSpace(string(input)) == "" {
		return "", errors.New("empty input")
	}

	// 入力がXMLとして有効かチェック
	var dummy interface{}
	if err := xml.Unmarshal(input, &dummy); err != nil {
		return "", errors.New("invalid XML format")
	}

	output := xmlfmt.FormatXML(string(input), "", "  ")
	if output == "" {
		return "", errors.New("invalid XML format")
	}

	return output, nil
}
