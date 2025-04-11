package usecase

import (
	"errors"
	"prettify/domain"
	"prettify/infrastructure"
)

func FormatInput(input []byte, formatType string) (string, error) {
	var formatter domain.Formatter

	switch formatType {
	case "json":
		formatter = &infrastructure.JSONFormatter{}
	case "sql":
		formatter = &infrastructure.SQLFormatter{}
	default:
		return "", errors.New("未対応のフォーマットタイプです")
	}

	return formatter.Format(input)
}
