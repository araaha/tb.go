package cmd

import (
	"testing"

	tb "github.com/araaha/tb.go/taskbook"
)

func Test_check(t *testing.T) {
	var testTaskBook tb.Book

	_ = captureOutputTest(func() {
		testTaskBook.AddTask("6", false, []string{}, false, false, 1)
		testTaskBook.AddTask("7", false, []string{}, true, false, 1)
		testTaskBook.AddTask("8", false, []string{}, false, false, 1)
		testTaskBook.AddTask("9", false, []string{}, true, false, 1)
		testTaskBook.AddNote("note2", false, []string{}, 1)
	})

	taskBook = testTaskBook

	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Check task",
			args:     []string{"6"},
			expected: tb.MarkOrUnmarkAttribute([]string{"6"}, nil, "Checked task", "Unchecked task"),
		},
		{
			name:     "Uncheck task",
			args:     []string{"7"},
			expected: tb.MarkOrUnmarkAttribute(nil, []string{"7"}, "Checked task", "Unchecked task"),
		},
		{
			name:     "Start and Pause",
			args:     []string{"6", "8", "9"},
			expected: tb.MarkOrUnmarkAttribute([]string{"8"}, []string{"6", "9"}, "Checked task", "Unchecked task"),
		},
		{
			name:     "note",
			args:     []string{"10"},
			expected: tb.ItemIsNote(10).Error(),
		},
		{
			name:     "note and task",
			args:     []string{"10", "9"},
			expected: tb.ItemIsNote(10).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutputTest(func() {
				check(tt.args)
			})
			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
