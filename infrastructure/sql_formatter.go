package infrastructure

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xwb1989/sqlparser"
)

// import (
// 	"strings"
// 	"unicode"

// 	"github.com/cockroachdb/cockroachdb-parser/pkg/sql/sem/tree"
// 	"github.com/mjibson/sqlfmt"
// )

type SQLFormatter struct{}

// func (f *SQLFormatter) Format(input []byte) (string, error) {
// 	stmt := string(input)
// 	cfg := tree.DefaultPrettyCfg()
// 	//cfg.Indent = 2       // インデント幅
// 	cfg.LineWidth = 50   // 改行される最大幅
// 	cfg.Simplify = false // 簡略化しない
// 	cfg.UseTabs = false  // タブではなくスペースを使用

// 	formatted, err := sqlfmt.FmtSQL(cfg, []string{stmt})
// 	if err != nil {
// 		return "", err
// 	}

// 	return strings.TrimRightFunc(formatted, unicode.IsSpace), nil
// }

func (f *SQLFormatter) Format(input []byte) (string, error) {
	stmt, err := sqlparser.Parse(string(input))
	if err != nil {
		return "", fmt.Errorf("SQLの解析エラー: %v", err)
	}
	// 解析されたSQLを整形する
	formattedSQL := sqlparser.String(stmt)

	// 改行を挿入する箇所を定義
	replacements := map[string]string{
		`(?i)\s+SELECT`:     "\nSELECT",
		`(?i)\s+FROM`:       "\nFROM",
		`(?i)\s+WHERE`:      "\nWHERE",
		`(?i)\s+GROUP BY`:   "\nGROUP BY",
		`(?i)\s+ORDER BY`:   "\nORDER BY",
		`(?i)\s+LIMIT`:      "\nLIMIT",
		`(?i)\s+JOIN`:       "\nJOIN",
		`(?i)\s+INNER JOIN`: "\nINNER JOIN",
		`(?i)\s+LEFT JOIN`:  "\nLEFT JOIN",
		`(?i)\s+RIGHT JOIN`: "\nRIGHT JOIN",
		`(?i)\s+ON`:         "\nON",
		`(?i)\s+AND`:        "\nAND",
		`(?i)\s+OR`:         "\nOR",
	}

	// 正規表現で置換処理を行う
	for pattern, replacement := range replacements {
		re := regexp.MustCompile(pattern)
		formattedSQL = re.ReplaceAllString(formattedSQL, replacement)
	}

	// 不要な空白や改行を削除
	formattedSQL = strings.TrimSpace(formattedSQL)
	formattedSQL = strings.ReplaceAll(formattedSQL, "\n", "\n    ")

	fmt.Println("=====START=====")
	fmt.Println(formattedSQL)
	fmt.Println("=====END=====")

	return formattedSQL, nil
}
