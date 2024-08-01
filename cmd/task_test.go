package cmd

import (
	tb "github.com/araaha/tb.go/taskbook"
	"testing"
)

func Test_task(t *testing.T) {
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
			name:     "No desc but board",
			args:     []string{"@abc"},
			expected: tb.MissingDesc().Error(),
		},
		{
			name:     "No desc but two board",
			args:     []string{"@abc", "@cba"},
			expected: tb.MissingDesc().Error(),
		},
		{
			name:     "Empty board",
			args:     []string{"@", "nice"},
			expected: tb.MissingBoards().Error(),
		},
		{
			name:     "Empty desc",
			args:     []string{""},
			expected: tb.MissingDesc().Error(),
		}, {
			name:     "task with default board",
			args:     []string{"Read hw"},
			expected: tb.ItemCreated(13, true),
		},
		{
			name:     "task with board",
			args:     []string{"@work", "complete work"},
			expected: tb.ItemCreated(14, true),
		},
		{
			name:     "long task with multiple boards",
			args:     []string{"@work", "@homework", "complete work", "hurry"},
			expected: tb.ItemCreated(15, true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutputTest(func() {
				task(tt.args)
			})

			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
