package tb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

var nextID int

// Base struct that contains common fields
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
	items []Item
}

func (b *Book) add(item Item) {

	nextID++
	item.GetBaseItem().ID = nextID
	item.GetBaseItem().Date = time.Now().Format("Mon Jan 02 2006")
	item.GetBaseItem().Timestamp = time.Now().UnixMilli()

	b.items = append(b.items, item)
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
	b.add(task)
	fmt.Printf("Created task: %v\n", task.ID)
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
	b.add(note)
	fmt.Printf("Created note: %d", note.ID)
}

func (b *Book) update(id int, modifyItem func(Item) Item) error {
	idx := b.getIndexByID(id)
	if idx == -1 {
		return fmt.Errorf("Unable to find item with id %d", id)
	}
	b.items[idx] = modifyItem(b.items[idx])
	return nil
}

func (b *Book) UpdateTask(id int, modifyTask func(*Task) *Task) error {
	taskUpdater := func(item Item) Item {
		task, ok := item.(*Task)
		if !ok {
			return item
		}

		return modifyTask(task)
	}

	return b.update(id, taskUpdater)
}

func (b *Book) UpdateNote(id int, modifyNote func(*Note) *Note) error {
	noteUpdater := func(item Item) Item {
		note, ok := item.(*Note)
		if !ok {
			return item
		}

		return modifyNote(note)
	}

	return b.update(id, noteUpdater)
}

func (b *Book) Delete(id int) error {
	modifyItem := func(item Item) Item {
		item.GetBaseItem().IsArchive = true
		return item
	}
	if err := b.update(id, modifyItem); err != nil {
		return err
	}
	return nil
}

func (b *Book) Clear(id int) error {
	idx := b.getIndexByID(id)
	if idx == -1 {
		return fmt.Errorf("invalid ID")
	}
	b.items = append(b.items[:idx], b.items[idx+1:]...)
	return nil
}

func (b *Book) Store(fileName string) error {
	data, err := json.MarshalIndent(b.items, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func (b *Book) Read(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		return err
	}

	var rawItems []json.RawMessage
	if err := json.Unmarshal(data, &rawItems); err != nil {
		return err
	}

	b.items = make([]Item, len(rawItems))
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
			b.items[i] = &task
		} else {
			var note Note
			if err := json.Unmarshal(raw, &note); err != nil {
				return err
			}
			b.items[i] = &note
		}
	}

	if err := json.Unmarshal(data, &b.items); err != nil {
		fmt.Printf("%v", err)
		return err
	}

	nextID = b.getMaxID()

	return nil

}

func (b *Book) getIndexByID(id int) int {
	for i, item := range b.items {
		if item.GetBaseItem().ID == id {
			return i
		}
	}
	return -1
}

func (b *Book) getMaxID() int {
	it := b.items
	if len(it) > 0 {
		return it[len(it)-1].GetBaseItem().ID
	} else {
		return 0
	}
}

func (b *Book) GetAllID(at bool) []string {
	var ids []string
	for _, item := range b.items {
		id := strconv.Itoa(item.GetBaseItem().ID)
		if at {
			ids = append(ids, "@"+id)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
}

func (b *Book) GetAllBoard() []string {
	var boards []string
	for _, item := range b.items {
		if item.GetBaseItem().Boards[0] != "My Board" {
			boards = append(boards, item.GetBaseItem().Boards...)
		}
	}
	return boards
}

// func Create creates a taskbook.json if it does not exist
func Create() {
	if _, err := os.Stat("/home/araaha/taskbook.json"); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create("/home/araaha/taskbook.json")
		if err != nil {
			panic(err)
		}
	}
}

// groupByBoard groups items by board
func (b *Book) groupByBoard() (map[string][]Item, []string) {
	itemsByBoard := make(map[string][]Item)
	sortedBoard := make([]string, len(itemsByBoard))

	for _, item := range b.items {
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

// groupByDate groups items by date
func (b *Book) groupByDate(a bool) (map[string][]Item, []string) {
	itemsByDate := make(map[string][]Item)
	sortedDates := make([]string, len(itemsByDate))

	for _, item := range b.items {
		bi := item.GetBaseItem()
		if a || !bi.IsArchive {
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
