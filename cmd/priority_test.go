package cmd

import (
	tb "github.com/araaha/tb.go/taskbook"
	"testing"
)

func Test_priority(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "No args",
			args:     []string{},
			expected: tb.InvalidPriority().Error(),
		},
		{
			name:     "no @",
			args:     []string{"cool", "nice"},
			expected: tb.MissingID().Error(),
		},
		{
			name:     "@ but empty",
			args:     []string{"@", "@"},
			expected: tb.InvalidID("").Error(),
		},
		{
			name:     "prio out of bounds",
			args:     []string{"@9", "10"},
			expected: tb.InvalidPriority().Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutputTest(func() {
				priority(tt.args)
			})

			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
