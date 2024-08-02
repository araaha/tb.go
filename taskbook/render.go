package tb

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/spf13/viper"
)

var (
	ErrMissingBoards       = "No boards were provided"
	ErrMissingDesc         = "No description was provided"
	ErrMissingID           = "No id was provided"
	ErrInvalidID           = "Unable to find item with id:"
	ErrInvalidIDArgNumber  = "More than one id was given as input"
	ErrInvalidPriority     = "Priority can only be 1, 2 or 3"
	ErrItemAlreadyDeleted  = "Item has already been deleted"
	ErrItemAlreadyArchived = "Item has already been archived"
	ErrItemNotArchived     = "Item has not been archived:"
	ErrItemIsNote          = "Item is a note:"
)

var (
	completeIcon = setColor("", "green", false)
	errorIcon    = setColor("", "red", false)
)

// DisplayByDate displays either archived items or non-archived items grouped by date
func (b *Book) DisplayByDate(a bool) {
	itemsByDate, sortedDates := b.groupByDate(a)

	for _, date := range sortedDates {
		displayCategory(date, itemsByDate[date])
		for _, item := range itemsByDate[date] {
			displayItemByDate(item)
		}
	}
	color.Printf("\n")
	displayStats(b.Items)
}

// DisplayByBoard displays items grouped by board
func (b *Book) DisplayByBoard() {
	itemsByBoard, sortedBoards := b.groupByBoard()

	for _, board := range sortedBoards {
		displayCategory(board, itemsByBoard[board])
		for _, item := range itemsByBoard[board] {
			displayItemByBoard(item)
		}
	}
	//Add newline at the end of loop
	color.Printf("\n")
	displayStats(b.Items)
}

// DisplayByBoardList displays items grouped by a list of Boards
func (b *Book) DisplayByBoardList(boards []string) {
	var sortedBoardList []string

	itemsByBoard, sortedBoards := b.groupByBoard()

	// Create set of boards
	bm := make(map[string]struct{})
	for _, b := range boards {
		bm[b] = struct{}{}
	}

	for _, b := range sortedBoards {
		if _, ok := bm[b]; ok {
			sortedBoardList = append(sortedBoardList, b)
		}
	}

	if len(sortedBoardList) == 0 {
		return
	}

	for _, board := range sortedBoardList {
		displayCategory(board, itemsByBoard[board])
		for _, item := range itemsByBoard[board] {
			displayItemByBoard(item)
		}
	}

	//Add newline at the end of loop
	color.Printf("\n")
	displayStats(b.Items)
}

// displayItemByBoard displays item by board
func displayItemByBoard(item Item) {
	id, icon, desc := buildMessage(item)
	age := getAge(item)
	star := getStar(item)

	alignedID := fmt.Sprintf("%5d.", id)
	ageDays := fmt.Sprintf("%vd", age)

	if age < 1 {
		color.Printf("%s %s %s %s\n", gray(alignedID), icon, desc, star)
	} else {
		color.Printf("%s %s %s %s %s\n", gray(alignedID), icon, desc, gray(ageDays), star)
	}

}

// displayItemByDate displays item by board
func displayItemByDate(item Item) {
	id, icon, desc := buildMessage(item)
	star := getStar(item)

	alignedID := fmt.Sprintf("%5d.", id)

	boards := ""
	bi := item.GetBaseItem()
	for i := 0; i < len(bi.Boards); i++ {
		if bi.Boards[i] != "My Board" {
			if i == len(bi.Boards)-1 {
				boards += bi.Boards[i]
			} else {
				boards += bi.Boards[i] + " "
			}
		}
	}

	color.Printf("%s %s %s %s %s\n", gray(alignedID), icon, desc, gray(boards), yellow(star))
}

// displayStats displays stats of []Item
func displayStats(items []Item) {
	comp, inprog, pending, notes := getStats(items)
	tasks := comp + inprog + pending

	frac := gray(color.Sprintf("%d of %d tasks complete", comp, tasks))

	color.Printf("  %s\n", frac)
	color.Printf("  %s %s%s %s%s %s%s %s\n\n", green(comp), gray("done · "), cyan(inprog), gray("in-progress · "), magenta(pending), gray("pending · "), blue(notes), gray("notes"))
}

// displayCategory displays category of []Item
func displayCategory(cat string, items []Item) {
	stats := buildCategory(items)

	color.Printf("\n")
	color.Printf("  %s %s\n", whiteUnderscore(cat), gray(stats))
}

// buildCategory builds category stats
func buildCategory(items []Item) (stats string) {
	comp, inprog, pending, _ := getStats(items)

	tasks := comp + inprog + pending
	if tasks == 0 {
		stats = ""
	} else {
		stats = color.Sprintf("[%d/%d]", comp, tasks)
	}

	return stats
}

// buildMessage builds item message details
func buildMessage(item Item) (id int, icon string, desc string) {
	id, icon, desc = item.GetBaseItem().ID, getIcon(item), item.GetBaseItem().Description

	if t, ok := item.(*Task); ok {
		if t.IsComplete {
			desc = gray(desc)
		} else if !t.IsComplete && t.Priority == 2 {
			desc = yellowUnderscore(desc) + yellow(" (!)")
		} else if !t.IsComplete && t.Priority == 3 {
			desc = redUnderscore(desc) + red(" (!!)")
		}
	}

	return id, icon, desc
}

// getAge calculates the age of an item in days
func getAge(item Item) (age float64) {
	diff := time.Now().UnixMilli() - item.GetBaseItem().Timestamp

	//get total time difference in milliseconds
	dur := time.Duration(diff) * time.Millisecond

	//compute age in days
	age = math.Floor(dur.Hours() / 24)

	return age
}

// getStats returns the stats of []Item
func getStats(items []Item) (comp int, inprog int, pending int, notes int) {
	for _, item := range items {
		if t, ok := item.(*Task); ok && !t.IsArchive {
			if t.IsComplete {
				comp++
			} else if t.InProgress {
				inprog++
			} else {
				pending++
			}
		} else if !item.GetBaseItem().IsArchive {
			notes++
		}
	}
	return comp, inprog, pending, notes
}

// getIcon returns the icon of an item
func getIcon(item Item) string {
	icon := ""
	if t, ok := item.(*Task); ok {
		if t.IsComplete {
			icon = setColor("", "green", false)
		} else if t.InProgress {
			icon = setColor("", "cyan", false)
		} else {
			icon = setColor("", "magenta", false)
		}
	} else {
		icon = setColor("󰎚", "blue", false)
	}

	return icon
}

// getStar returns star if item is starred
func getStar(item Item) string {
	if item.GetBaseItem().IsStarred {
		return setColor("󰓎", "yellow", false)
	}
	return ""
}

// InvalidID returns an  error for an invalid ID
func InvalidID(id any) error {
	return fmt.Errorf("\n  %s  %s %v\n", errorIcon, ErrInvalidID, gray(id))
}

// InvalidPriority returns an error for an invalid priority
func InvalidIDArgNumber() error {
	return fmt.Errorf("\n  %s  %s\n", errorIcon, ErrInvalidIDArgNumber)
}

// MissingID returns an error for a missing ID
func InvalidPriority() error {
	return fmt.Errorf("\n  %s  %s\n", errorIcon, ErrInvalidPriority)
}

// MissingBoards returns an error for missing boards
func MissingID() error {
	return fmt.Errorf("\n  %s  %s\n", errorIcon, ErrMissingID)
}

// MissingBoards returns an error for missing boards
func MissingBoards() error {
	return fmt.Errorf("\n  %s  %s\n", errorIcon, ErrMissingBoards)
}

// MissingDesc returns an error for a missing description
func MissingDesc() error {
	return fmt.Errorf("\n  %s  %s\n", errorIcon, ErrMissingDesc)
}

// ItemAlreadyArchived returns an error for an already archived item
func ItemAlreadyArchived() error {
	return fmt.Errorf("\n  %s  %s\n", errorIcon, ErrItemAlreadyArchived)
}

// ItemNotArchived returns an error for an item that is not archived
func ItemNotArchived(id int) error {
	return fmt.Errorf("\n  %s  %s %s\n", errorIcon, ErrItemNotArchived, gray(id))
}

// ItemIsNote returns an error for an item that is a note
func ItemIsNote(id int) error {
	return fmt.Errorf("\n  %s  %s %s\n", errorIcon, ErrItemIsNote, gray(id))
}

// MarkOrUnmarkAttribute marks or unmarks an attribute
func MarkOrUnmarkAttribute(mark []string, unmark []string, markAttrMsg string, unmarkAttrMsg string) string {
	if len(mark) != 0 && len(unmark) == 0 {
		return fmt.Sprintf("\n  %s  \n", MarkAttribute(mark, markAttrMsg))
	}
	if len(mark) == 0 && len(unmark) != 0 {
		return fmt.Sprintf("\n  %s  \n", MarkAttribute(unmark, unmarkAttrMsg))
	}
	return fmt.Sprintf("\n  %s\n\n  %s\n", MarkAttribute(mark, markAttrMsg), MarkAttribute(unmark, unmarkAttrMsg))
}

// MarkAttribute marks an attribute
func MarkAttribute(ids []string, attrMsg string) string {
	if len(ids) == 1 {
		return fmt.Sprintf("%s  %s: %s", completeIcon, attrMsg, gray(ids[0]))
	}
	return fmt.Sprintf("%s  %s: %s", completeIcon, attrMsg+"s", gray(strings.Join(ids, ", ")))
}

// MarkRestored marks an item as restored
func MarkRestored(ids []string) string {
	if len(ids) == 1 {
		return fmt.Sprintf("\n  %s  Restored item: %s\n", completeIcon, gray(ids[0]))
	}
	return fmt.Sprintf("\n  %s  Restored items: %s\n", completeIcon, gray(strings.Join(ids, ", ")))

}

// ItemCreated returns a message for a created item
func ItemCreated(id int, isTask bool) string {
	if isTask {
		return fmt.Sprintf("\n  %s  Created task: %s\n", completeIcon, gray(id))
	}
	return fmt.Sprintf("\n  %s  Created note: %s\n", completeIcon, gray(id))
}

// ItemEdited returns a message for an edited item
func ItemEdited(id int) string {
	return fmt.Sprintf("\n  %s  Updated description of item: %s\n", completeIcon, gray(id))
}

// ItemDeleted returns a message for a deleted item
func ItemDeleted(ids []string) string {
	if len(ids) == 1 {
		return fmt.Sprintf("\n  %s  Deleted item: %s\n", completeIcon, gray(ids[0]))
	}
	return fmt.Sprintf("\n  %s  Deleted items: %s\n", completeIcon, gray(strings.Join(ids, ", ")))
}

// ItemMoved returns a message for a moved item
func ItemMoved(id int, boards []string) string {
	return fmt.Sprintf("\n  %s  Moved item %s to %s\n", completeIcon, gray(id), gray(strings.Join(boards, ", ")))
}

// ItemPriority returns a message for a priority change
func ItemPriority(id int, prioLevel string) string {
	if prioLevel == "low" {
		return fmt.Sprintf("\n  %s  Updated priority of task: %s to %s\n", completeIcon, gray(id), green(prioLevel))
	}
	if prioLevel == "medium" {
		return fmt.Sprintf("\n  %s  Updated priority of task: %s to %s\n", completeIcon, gray(id), yellow(prioLevel))
	}
	return fmt.Sprintf("\n  %s  Updated priority of task: %s to %s\n", completeIcon, gray(id), red(prioLevel))
}

// setColor sets the color of any
func setColor(s any, style string, underscore bool) string {
	if !color.Enable {
		return fmt.Sprintf("%v", s)
	}
	colors := map[string]color.Color{
		"white":   color.FgWhite,
		"red":     color.FgRed,
		"yellow":  color.FgYellow,
		"gray":    color.FgGray,
		"green":   color.FgGreen,
		"blue":    color.FgBlue,
		"magenta": color.FgMagenta,
		"cyan":    color.FgCyan,
	}

	c := viper.GetString("colors." + style)
	if isHex(c) {
		if underscore {
			return color.HEXStyle(c).SetOpts(color.Opts{color.OpUnderscore}).Sprintf("%v", s)
		}
		return color.HEXStyle(c).Sprintf("%v", s)
	}

	if underscore {
		return color.New(colors[style], color.OpUnderscore).Sprintf("%v", s)
	}
	return colors[style].Sprintf("%v", s)
}

// isHex checks if a string is a hex color
func isHex(s string) bool {
	const hp = `^#[A-Fa-f0-9]{6}|[A-Fa-f0-9]{3}$`
	m, _ := regexp.MatchString(hp, s)
	return m
}

func red(s any) string {
	return setColor(s, "red", false)
}

func yellow(s any) string {
	return setColor(s, "yellow", false)
}

func gray(s any) string {
	return setColor(s, "gray", false)
}

func green(s any) string {
	return setColor(s, "green", false)
}

func blue(s any) string {
	return setColor(s, "blue", false)
}

func magenta(s any) string {
	return setColor(s, "magenta", false)
}

func cyan(s any) string {
	return setColor(s, "cyan", false)
}

func whiteUnderscore(s any) string {
	return setColor(s, "white", true)
}

func yellowUnderscore(s any) string {
	return setColor(s, "yellow", true)
}

func redUnderscore(s any) string {
	return setColor(s, "red", true)
}
