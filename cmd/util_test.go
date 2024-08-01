package cmd

import (
	tb "github.com/araaha/tb.go/taskbook"
	"testing"
)

func Test_validArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		excl     bool
		expected string
	}{
		{
			name:     "No args",
			args:     []string{},
			excl:     true,
			expected: tb.MissingID().Error(),
		},
		{
			name:     "empty arg",
			args:     []string{"", ""},
			excl:     true,
			expected: tb.InvalidID("").Error(),
		},
		{
			name:     "index too large",
			args:     []string{"100"},
			excl:     true,
			expected: tb.InvalidID(100).Error(),
		},
		{
			name:     "index too small",
			args:     []string{"-100"},
			excl:     true,
			expected: tb.InvalidID(-100).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := validArgs(tt.args, tt.excl).Error()

			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
