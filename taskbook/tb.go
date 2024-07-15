package tb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
	fmt.Printf("Created task: %v", task.ID)
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
	b.update(id, modifyItem)
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
		fmt.Errorf("%v", err)
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

func (b *Book) Print() {
	for _, item := range b.items {
		fmt.Println(item.GetBaseItem().ID)
		fmt.Println(item.GetBaseItem().Description)
		fmt.Println(item.GetBaseItem().Date)
		fmt.Println(item.GetBaseItem().Boards)
	}
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

func (b *Book) AllID(at bool) []string {
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

func (b *Book) AllBoard() []string {
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
