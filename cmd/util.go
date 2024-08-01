package cmd

import (
	"bytes"
	tb "github.com/araaha/tb.go/taskbook"
	"io"
	"os"
	"strconv"
)

// validArgs validates a list of item IDs to ensure they are valid integers,
// exist in the task book, and meet archive status conditions based on the excludeArchive flag.
func validArgs(args []string, excludeArchive bool) error {
	if len(args) == 0 {
		return tb.MissingID()
	}

	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return tb.InvalidID(arg)
		}

		_, item := taskBook.GetIndexAndItemByID(id)
		if item == nil {
			return tb.InvalidID(id)
		}

		// Exclude archived item if excludeArchive
		if excludeArchive && item.GetBaseItem().IsArchive {
			return tb.ItemAlreadyArchived()
		}

		if !excludeArchive && !item.GetBaseItem().IsArchive {
			return tb.ItemNotArchived(id)
		}
	}
	return nil
}

// captureOutputTest captures the output of f without printing to Stdout
func captureOutputTest(f func()) string {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = orig

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	r.Close()

	out := buf.String()
	if len(out) > 0 {
		return string(out[:len(out)-1])
	}
	return ""
}
