package api

import "testing"

func Test_printWord(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "one"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printWord()
		})
	}
}
