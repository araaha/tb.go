package cmd

import (
	"testing"

	tb "github.com/araaha/tb.go/taskbook"
)

func Test_begin(t *testing.T) {
	var testTaskBook tb.Book

	_ = captureOutputTest(func() {
		testTaskBook.AddTask("1", false, []string{}, false, false, 1)
		testTaskBook.AddTask("2", false, []string{}, false, true, 1)
		testTaskBook.AddTask("3", false, []string{}, false, false, 1)
		testTaskBook.AddTask("4", false, []string{}, false, true, 1)
		testTaskBook.AddNote("note", false, []string{}, 1)
	})

	taskBook = testTaskBook
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Start task",
			args:     []string{"1"},
			expected: tb.MarkOrUnmarkAttribute([]string{"1"}, nil, "Started task", "Paused task"),
		},
		{
			name:     "Pause task",
			args:     []string{"2"},
			expected: tb.MarkOrUnmarkAttribute(nil, []string{"2"}, "Started task", "Paused task"),
		},
		{
			name:     "Start and Pause",
			args:     []string{"3", "4"},
			expected: tb.MarkOrUnmarkAttribute([]string{"3"}, []string{"4"}, "Started task", "Paused task"),
		},
		{
			name:     "note",
			args:     []string{"5"},
			expected: tb.ItemIsNote(5).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutputTest(func() {
				begin(tt.args)
			})

			if output != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, output)
			}
		})
	}
}
