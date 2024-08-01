package cmd

import (
	tb "github.com/araaha/tb.go/taskbook"
	"testing"
)

func Test_edit(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "No args",
			args:     []string{},
			expected: tb.MissingID().Error(),
		},
		{
			name:     "no @",
			args:     []string{"cool"},
			expected: tb.MissingID().Error(),
		},
		{
			name:     "@ but empty",
			args:     []string{"@", "nice"},
			expected: tb.InvalidID("").Error(),
		},
		{
			name:     "id out of bounds",
			args:     []string{"@10", "awesome"},
			expected: tb.InvalidID("10").Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutputTest(func() {
				edit(tt.args)
			})

			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
