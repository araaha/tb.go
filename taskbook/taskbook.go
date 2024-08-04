package tb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

var (
	nextID int
)

// BaseItem is a struct that contains common fields
type BaseItem struct {
	ID          int      `json:"_id"`
	Date        string   `json:"_date"`
	Timestamp   int64    `json:"_timestamp"`
	Description string   `json:"description"`
	IsStarred   bool     `json:"isStarred"`
	Boards      []string `json:"boards"`
	IsTask      bool     `json:"_isTask"`
	IsArchive   bool     `json:"isArchive"`
	Priority    int      `json:"priority"`
}

type Task struct {
	BaseItem
	IsComplete bool `json:"isComplete"`
	InProgress bool `json:"inProgress"`
}

type Note struct {
	BaseItem
}

// Interface that will be implemented by both Task and Note
type Item interface {
	GetBaseItem() *BaseItem
}

// Method to get the embedded BaseItem for Task
func (t *Task) GetBaseItem() *BaseItem {
	return &t.BaseItem
}

// Method to get the embedded BaseItem for Note
func (n *Note) GetBaseItem() *BaseItem {
	return &n.BaseItem
}

type Book struct {
	Items []Item
}

func (b *Book) add(item Item) error {
	if item == nil {
		return fmt.Errorf("Item is nil")
	}

	nextID++
	item.GetBaseItem().ID = nextID
	item.GetBaseItem().Date = time.Now().Format("Mon Jan 02 2006")
	item.GetBaseItem().Timestamp = time.Now().UnixMilli()

	b.Items = append(b.Items, item)
	return nil
}

func createTask(desc string, star bool, boards []string, comp bool, prog bool, prio int) *Task {
	return &Task{
		BaseItem: BaseItem{
			Description: desc,
			IsStarred:   star,
			Boards:      boards,
			IsTask:      true,
			IsArchive:   false,
			Priority:    prio,
		},
		IsComplete: comp,
		InProgress: prog,
	}
}

func (b *Book) AddTask(desc string, star bool, boards []string, comp bool, prog bool, prio int) {
	task := createTask(desc, star, boards, comp, prog, prio)
	if err := b.add(task); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ItemCreated(task.ID, true))
}

func createNote(desc string, star bool, boards []string, prio int) *Note {
	return &Note{
		BaseItem: BaseItem{
			Description: desc,
			IsStarred:   star,
			Boards:      boards,
			IsTask:      false,
			IsArchive:   false,
			Priority:    prio,
		},
	}
}

func (b *Book) AddNote(desc string, star bool, boards []string, prio int) {
	note := createNote(desc, star, boards, prio)
	if err := b.add(note); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ItemCreated(note.ID, false))
}

func (b *Book) update(id int, modifyItem func(Item) Item) {
	idx, _ := b.GetIndexAndItemByID(id)
	if idx == -1 {
		return
	}
	b.Items[idx] = modifyItem(b.Items[idx])
}

// Update updates item
func (b *Book) Update(id int, modifyItem func(Item) Item) {
	taskUpdater := func(item Item) Item {
		task, ok := item.(*Task)
		if !ok {
			return item
		}

		return modifyItem(task)
	}

	noteUpdater := func(item Item) Item {
		note, ok := item.(*Note)
		if !ok {
			return item
		}

		return modifyItem(note)
	}

	_, item := b.GetIndexAndItemByID(id)
	if _, ok := item.(*Task); ok {
		b.update(id, taskUpdater)
	}
	b.update(id, noteUpdater)
}

// UpdateTask updates task
func (b *Book) UpdateTask(id int, modifyTask func(*Task) *Task) {
	taskUpdater := func(item Item) Item {
		task, ok := item.(*Task)
		if !ok {
			return item
		}

		return modifyTask(task)
	}

	b.update(id, taskUpdater)
}

// UpdateNote updates note
func (b *Book) UpdateNote(id int, modifyNote func(*Note) *Note) {
	noteUpdater := func(item Item) Item {
		note, ok := item.(*Note)
		if !ok {
			return item
		}

		return modifyNote(note)
	}

	b.update(id, noteUpdater)
}

// Delete places item in archive given ID
func (b *Book) Delete(id int) {
	modifyItem := func(item Item) Item {
		item.GetBaseItem().IsArchive = true
		return item
	}
	b.update(id, modifyItem)
}

// Remove removes every item in archive
func (b *Book) Remove() {

	var updatedItems []Item

	for _, item := range b.Items {
		if !item.GetBaseItem().IsArchive {
			updatedItems = append(updatedItems, item)
		}
	}

	b.Items = updatedItems
}

// Store stores items in storage or creates a file if necessary
func (b *Book) Store() error {
	data, err := json.MarshalIndent(b.Items, "", "    ")
	if err != nil {
		return err
	}

	// Create storage file if it does not exist
	if err := Create(); err != nil {
		return err
	}

	// Write to storage file
	return os.WriteFile(getStoragePath(), data, 0644)
}

// getStoragePath gets the storage path for our Book
func getStoragePath() string {
	if os.Getenv("TEST_MODE") == "true" {
		return "test_storage.json"
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Check for XDG_DATA_HOME
	if xdgDataHome := os.Getenv("XDG_DATA_HOME"); xdgDataHome != "" {
		storageDir := filepath.Join(xdgDataHome, "taskbook")
		storageFile := filepath.Join(storageDir, "taskbook.json")
		return storageFile
	}

	storageDir := filepath.Join(home, ".local", "share", "taskbook")
	storageFile := filepath.Join(storageDir, "storage.json")
	return storageFile
}

// func Create creates a taskbook.json if it does not exist
func Create() error {
	storageFile := getStoragePath()
	if _, err := os.Stat(storageFile); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(filepath.Dir(storageFile), 0700); err != nil {
			return err
		}
		if _, err := os.Create(storageFile); err != nil {
			return err
		}

	}

	return nil
}

// Read unmarshals storageFile
func (b *Book) Read() error {
	storageFile := getStoragePath()
	data, err := os.ReadFile(storageFile)
	if err != nil {
		return nil
	}

	if len(data) == 0 {
		return nil
	}

	var rawItems []json.RawMessage
	if err := json.Unmarshal(data, &rawItems); err != nil {
		return err
	}

	b.Items = make([]Item, len(rawItems))
	for i, raw := range rawItems {
		var base BaseItem
		if err := json.Unmarshal(raw, &base); err != nil {
			return err
		}

		if base.IsTask {
			var task Task
			if err := json.Unmarshal(raw, &task); err != nil {
				return err
			}
			b.Items[i] = &task
		} else {
			var note Note
			if err := json.Unmarshal(raw, &note); err != nil {
				return err
			}
			b.Items[i] = &note
		}
	}

	if err := json.Unmarshal(data, &b.Items); err != nil {
		fmt.Printf("%v", err)
		return err
	}

	// set nextID to max ID
	nextID = b.getMaxID()

	return nil

}

// GetIndexAndItemByID gets index, item given ID
func (b *Book) GetIndexAndItemByID(id int) (int, Item) {
	for i, item := range b.Items {
		if item.GetBaseItem().ID == id {
			return i, item
		}
	}
	return -1, nil
}

// getMaxID gets the max ID of b
func (b *Book) getMaxID() int {
	it := b.Items
	if len(it) > 0 {
		return it[len(it)-1].GetBaseItem().ID
	} else {
		return 0
	}
}

// GetAllID returns every ID. If 'at' is true, every ID will have '@' as a prefix.
// If 'a' is true, IDs in the archive are included.
func (b *Book) GetAllID(at bool, a bool) []string {
	var ids []string

	for _, item := range b.Items {
		id := strconv.Itoa(item.GetBaseItem().ID)
		isArchived := item.GetBaseItem().IsArchive

		// Determine if the current item should be included based on 'a'
		if !a && isArchived {
			continue
		}

		// Determine if the ID should be prefixed with '@'
		if at {
			id = "@" + id
		}

		ids = append(ids, id)
	}

	return ids
}

// GetAllBoard gets every board. If 'a', includes boards in the archive
func (b *Book) GetAllBoard(a bool) []string {
	var boards []string
	for _, item := range b.Items {
		if item.GetBaseItem().Boards[0] != "My Board" {
			if a || !item.GetBaseItem().IsArchive {
				boards = append(boards, item.GetBaseItem().Boards...)
			}
		}
	}
	return boards
}

// groupByBoard groups items by board
func (b *Book) groupByBoard() (map[string][]Item, []string) {
	itemsByBoard := make(map[string][]Item)
	sortedBoard := make([]string, len(itemsByBoard))

	// group items by board
	for _, item := range b.Items {
		bi := item.GetBaseItem()
		if !bi.IsArchive {
			for _, board := range bi.Boards {
				itemsByBoard[board] = append(itemsByBoard[board], item)
			}
		}
	}

	for id := range itemsByBoard {
		sortedBoard = append(sortedBoard, id)
	}

	//sorts by first element in []int w.r.t date
	sort.Slice(sortedBoard, func(i, j int) bool {
		id1 := itemsByBoard[sortedBoard[i]][0]
		id2 := itemsByBoard[sortedBoard[j]][0]
		return id1.GetBaseItem().ID < id2.GetBaseItem().ID
	})

	return itemsByBoard, sortedBoard
}

// groupByDate groups archived|non-archived items by date
func (b *Book) groupByDate(a bool) (map[string][]Item, []string) {
	itemsByDate := make(map[string][]Item)
	sortedDates := make([]string, len(itemsByDate))

	for _, item := range b.Items {
		bi := item.GetBaseItem()
		if (a && bi.IsArchive) || (!a && !bi.IsArchive) {
			itemsByDate[bi.Date] = append(itemsByDate[bi.Date], item)
		}
	}

	for id := range itemsByDate {
		sortedDates = append(sortedDates, id)
	}

	//sorts by first element in []int w.r.t date
	sort.Slice(sortedDates, func(i, j int) bool {
		id1 := itemsByDate[sortedDates[i]][0]
		id2 := itemsByDate[sortedDates[j]][0]
		return id1.GetBaseItem().ID < id2.GetBaseItem().ID
	})

	return itemsByDate, sortedDates
}
