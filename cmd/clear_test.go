package cmd

import (
	tb "github.com/araaha/tb.go/taskbook"
	"testing"
)

func Test_clear(t *testing.T) {
	var testTaskBook tb.Book

	_ = captureOutputTest(func() {
		testTaskBook.AddTask("11", false, []string{}, true, false, 1)
		testTaskBook.AddTask("12", false, []string{}, true, false, 1)
	})

	taskBook = testTaskBook

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "completed task",
			args:     nil,
			expected: tb.ItemDeleted([]string{"11", "12"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutputTest(func() {
				clear()
			})

			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
