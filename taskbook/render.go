package tb

import (
	"fmt"
	"github.com/gookit/color"
	"math"
	"time"
)

var (
	LightWhiteUnderscore = color.New(color.FgLightWhite, color.OpUnderscore).Render
	YellowUnderscore     = color.New(color.FgYellow, color.OpUnderscore).Render
	RedUnderscore        = color.New(color.FgRed, color.OpUnderscore).Render
	Yellow               = color.New(color.FgYellow).Render
	Red                  = color.New(color.FgRed).Render
	Gray                 = color.New(color.FgGray).Render
	Green                = color.New(color.FgGreen).Render
	Blue                 = color.New(color.FgBlue).Render
	Magenta              = color.New(color.FgMagenta).Render
)

var (
	ErrMissingBoards      = "No boards were provided"
	ErrMissingDesc        = "No description was provided"
	ErrMissingID          = "No id was provided"
	ErrInvalidID          = "Unable to find item with id:"
	ErrInvalidIDArgNumber = "More than one id was given as input"
	ErrInvalidPriority    = "Priority can only be 1, 2 or 3"
)

var (
	StarIcon     = Yellow("󰓎")
	ProgressIcon = Blue("")
	CompleteIcon = Green("")
	PendingIcon  = Magenta("")
	NoteIcon     = Blue("󰎚")
	ErrorIcon    = Red("")
)

// DisplayByDate displays items grouped by date
func (b *Book) DisplayByDate(a bool) {
	itemsByDate, sortedDates := b.groupByDate(a)

	for _, date := range sortedDates {
		displayCategory(date, itemsByDate[date])
		for _, item := range itemsByDate[date] {
			displayItemByDate(item)
		}
	}
	color.Printf("\n")
	displayStats(b.items)
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
	displayStats(b.items)
}

// displayItemByBoard displays item by board
func displayItemByBoard(item Item) {
	id, icon, desc := buildMessage(item)
	age := getAge(item)
	star := getStar(item)

	if age < 1 {
		color.Printf("%s %s %s %s\n", Gray(color.Sprintf("%5d", id), "."), icon, desc, Yellow(star))
	} else {
		color.Printf("%s %s %s %v %s\n", Gray(color.Sprintf("%5d", id), "."), icon, desc, Gray(age, "d"), Yellow(star))
	}

}

// displayItemByDate displays item by board
func displayItemByDate(item Item) {
	id, icon, desc := buildMessage(item)
	star := getStar(item)

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

	color.Printf("%s %s %s %s %s\n", Gray(color.Sprintf("%5d", id), "."), icon, desc, Gray(boards), Yellow(star))
}

// displayStats displays stats of []Item
func displayStats(items []Item) {
	comp, inprog, pending, notes := getStats(items)
	tasks := comp + inprog + pending

	frac := Gray(color.Sprintf("%d of %d tasks complete", comp, tasks))

	color.Printf("  %s\n", frac)
	color.Printf("  %s %s%s %s%s %s%s %s\n\n", Green(comp), Gray("done · "), Blue(inprog), Gray("in-progress · "), Magenta(pending), Gray("pending · "), Blue(notes), Gray("notes"))
}

// displayCategory displays category of []Item
func displayCategory(cat string, items []Item) {
	stats := buildCategory(items)

	color.Printf("\n")
	color.Printf("  %s %s\n", LightWhiteUnderscore(cat), Gray(stats))
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
			desc = Gray(desc)
		} else if !t.IsComplete && t.Priority == 2 {
			desc = YellowUnderscore(desc) + Yellow(" (!)")
		} else if !t.IsComplete && t.Priority == 3 {
			desc = RedUnderscore(desc) + Red(" (!!)")
		}
	}

	return id, icon, desc
}

// getAge calculates the age of an item in days
func getAge(item Item) (age float64) {
	diff := time.Now().UnixMilli() - item.GetBaseItem().Timestamp
	dur := time.Duration(diff) * time.Millisecond
	age = math.Floor(dur.Hours() / 24)

	return age
}

// getStatss returns the stats of []Item
func getStats(items []Item) (comp int, inprog int, pending int, notes int) {
	for _, item := range items {
		if t, ok := item.(*Task); ok {
			if t.IsComplete {
				comp++
			} else if t.InProgress {
				inprog++
			} else {
				pending++
			}
		} else {
			notes++
		}
	}
	return comp, inprog, pending, notes
}

// getIcon returns the icon of an item
func getIcon(item Item) string {
	icon := ""
	if t, ok := item.(*Task); ok {
		if t.InProgress {
			icon = ProgressIcon
			return icon
		} else if t.IsComplete {
			icon = CompleteIcon
		} else {
			icon = PendingIcon
		}
	} else {
		icon = NoteIcon
	}

	return icon
}

// getStar returns star if item is starred
func getStar(item Item) string {
	if item.GetBaseItem().IsStarred {
		return StarIcon
	}

	return ""
}

func InvalidID(id int) error {
	return fmt.Errorf("\n  %s  %s %d\n", ErrorIcon, ErrInvalidID, id)
}

func InvalidIDArgNumber() error {
	return fmt.Errorf("\n  %s  %s\n", ErrorIcon, ErrInvalidIDArgNumber)
}

func InvalidPriority() error {
	return fmt.Errorf("\n  %s  %s\n", ErrorIcon, ErrInvalidPriority)
}

func MissingID() error {
	return fmt.Errorf("\n  %s  %s\n", ErrorIcon, ErrMissingID)
}

func MissingBoards() error {
	return fmt.Errorf("\n  %s  %s\n", ErrorIcon, ErrMissingID)
}

// TODO determine if []int or []string
func markComplete(_ []int, _ bool) error {
	return nil
}

func markStarted(_ []int, _ bool) error {
	return nil
}

func markPaused(_ []int, _ bool) error {
	return nil
}

func markStarred(_ []int, _ bool) error {
	return nil
}

func itemCreated(_ int, _ bool) error {
	return nil
}

func itemEdited(_ int) error {
	return nil
}

func itemDeleted(_ []int) error {
	return nil
}

func itemMoved(_ int, _ []string) error {
	return nil
}

func itemPriority(_ int, _ int) error {
	return nil
}

func itemArchived(_ []int) error {
	return nil
}
