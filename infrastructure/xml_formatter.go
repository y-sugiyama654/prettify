package infrastructure

import "github.com/go-xmlfmt/xmlfmt"

type XMLFormatter struct{}

func (f *XMLFormatter) Format(input []byte) (string, error) {
	output := xmlfmt.FormatXML(string(input), "", "  ")
	return string(output), nil
}
