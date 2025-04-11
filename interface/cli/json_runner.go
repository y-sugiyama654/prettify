package cli

import (
	"fmt"
	"os"
	"prettify/usecase"
)

func Run(formatType string, args []string) {
	var data []byte
	var err error

	if len(args) > 0 {
		data, err = os.ReadFile(args[0])
	} else {
		data, err = os.ReadFile("/dev/stdin")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "読み込み失敗: %v\n", err)
		os.Exit(1)
	}

	output, err := usecase.FormatInput(data, formatType)
	if err != nil {
		fmt.Fprintf(os.Stderr, "整形失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(output)
}
