package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func readFile(t *testing.T, filename string) string {
	t.Helper()
	content, err := os.ReadFile(filepath.Join("testdata", filename))
	if err != nil {
		t.Fatalf("テストデータの読み込みに失敗: %v", err)
	}
	return string(content)
}

func TestFormatJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "正常系: 単純なJSONオブジェクト",
			input:   readFile(t, "simple.json"),
			want:    readFile(t, "simple_want.json"),
			wantErr: false,
		},
		{
			name:    "正常系: ネストされたJSONオブジェクト",
			input:   readFile(t, "nested.json"),
			want:    readFile(t, "nested_want.json"),
			wantErr: false,
		},
		// {
		// 	name:    "正常系: 配列を含むJSONオブジェクト",
		// 	input:   readFile(t, "array.json"),
		// 	want:    readFile(t, "array_want.json"),
		// 	wantErr: false,
		// },
		// {
		// 	name:    "異常系: 不正なJSON",
		// 	input:   `{"name":"test",}`,
		// 	want:    "",
		// 	wantErr: true,
		// },
		// {
		// 	name:    "異常系: JSONではない入力",
		// 	input:   "not a json",
		// 	want:    "",
		// 	wantErr: true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("テストを実行します。")
			got, err := FormatJSON([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 直接文字列として比較
			if got != tt.want {
				t.Errorf("FormatJSON() = %v, want %v", got, tt.want)
				t.Errorf("差分:\n- want: %v (length: %d)\n+ got: %v (length: %d)", tt.want, len(tt.want), got, len(got))
			}
		})
	}
}
