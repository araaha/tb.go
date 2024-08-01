package cmd

import (
	tb "github.com/araaha/tb.go/taskbook"
	"testing"
)

func Test_note(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "No args",
			args:     []string{},
			expected: tb.MissingDesc().Error(),
		},
		{
			name:     "@ but empty",
			args:     []string{"nice", "@"},
			expected: tb.MissingBoards().Error(),
		},
		{
			name:     "empty note",
			args:     []string{""},
			expected: tb.MissingDesc().Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutputTest(func() {
				note(tt.args)
			})

			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
