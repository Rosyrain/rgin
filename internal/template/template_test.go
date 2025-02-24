package template

import (
	"testing"
)

func TestLoadTemplate(t *testing.T) {

	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"Load main.go.tmpl", "/main.go.tmpl", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadTemplate(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
