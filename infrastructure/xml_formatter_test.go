package infrastructure

import (
	"strings"
	"testing"
)

func TestFormatXML(t *testing.T) {
	path := "testdata/xml_formatter"
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "正常系: 空の入力",
			input:   readFile(t, "sample.xml", path),
			want:    readFile(t, "sample_want.xml", path),
			wantErr: false,
		},
		// 異常系
		{
			name:    "異常系: 空の入力",
			input:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "異常系: 不正なXML形式",
			input:   "invalid: xml: format",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := &XMLFormatter{}
			got, err := formatter.Format([]byte(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 改行文字を正規化して比較
			normalizedGot := strings.TrimSpace(strings.ReplaceAll(got, "\r\n", "\n"))
			normalizedWant := strings.TrimSpace(strings.ReplaceAll(tt.want, "\r\n", "\n"))

			if normalizedGot != normalizedWant {
				t.Errorf("FormatXML()\ngot:\n%v\nwant:\n%v", normalizedGot, normalizedWant)
			}
		})
	}
}
