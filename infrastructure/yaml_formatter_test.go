package infrastructure

import (
	"strings"
	"testing"
)

func TestFormatYAML(t *testing.T) {
	path := "testdata/yaml_formatter"
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "正常系: 単純なYAMLオブジェクト",
			input:   readFile(t, "simple.yaml", path),
			want:    readFile(t, "simple_want.yaml", path),
			wantErr: false,
		},
		{
			name:    "異常系: 不正なYAML形式",
			input:   "invalid: yaml: format",
			want:    "",
			wantErr: true,
		},
		{
			name:    "異常系: 空の入力",
			input:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "異常系: 不正な構文（括弧の不一致とシンボル）",
			input:   `{name: test, @invalid: true, [broken: syntax}`,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := &YAMLFormatter{}
			got, err := formatter.Format([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 改行文字を正規化して比較
			normalizedGot := strings.TrimSpace(strings.ReplaceAll(got, "\r\n", "\n"))
			normalizedWant := strings.TrimSpace(strings.ReplaceAll(tt.want, "\r\n", "\n"))

			if normalizedGot != normalizedWant {
				t.Errorf("FormatYAML()\ngot:\n%v\nwant:\n%v", normalizedGot, normalizedWant)
			}
		})
	}
}
