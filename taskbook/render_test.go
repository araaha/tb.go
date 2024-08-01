package tb

import (
	"testing"

	"github.com/gookit/color"
)

func Test_displayCategory(t *testing.T) {
	type args struct {
		cat   string
		items []Item
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			displayCategory(tt.args.cat, tt.args.items)
		})
	}
}

func Test_buildCategory(t *testing.T) {
	type args struct {
		items []Item
	}
	tests := []struct {
		name      string
		args      args
		wantStats string
	}{
		{
			name: "no tasks",
			args: args{
				items: []Item{
					&Note{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   false,
							Priority:    0,
						},
					},
				},
			},
			wantStats: "",
		},
		{
			name: "1 task",
			args: args{
				items: []Item{
					&Task{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   false,
							Priority:    0,
						},
						IsComplete: false,
						InProgress: false,
					},
				},
			},
			wantStats: color.Sprintf("[%d/%d]", 0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStats := buildCategory(tt.args.items); gotStats != tt.wantStats {
				t.Errorf("buildCategory() = %v, want %v", gotStats, tt.wantStats)
			}
		})
	}
}

func Test_buildMessage(t *testing.T) {
	type args struct {
		item Item
	}
	tests := []struct {
		name     string
		args     args
		wantId   int
		wantIcon string
		wantDesc string
	}{
		{
			name: "pending",
			args: args{
				item: &Task{
					BaseItem: BaseItem{
						ID:          1,
						Date:        "",
						Timestamp:   0,
						Description: "nice",
						IsStarred:   false,
						Boards:      []string{},
						IsTask:      false,
						IsArchive:   false,
						Priority:    0,
					},
					IsComplete: false,
					InProgress: false,
				},
			},
			wantId:   1,
			wantIcon: setColor("", "magenta", false),
			wantDesc: "nice",
		},
		{
			name: "complete",
			args: args{
				item: &Task{
					BaseItem: BaseItem{
						ID:          1,
						Date:        "",
						Timestamp:   0,
						Description: "nice",
						IsStarred:   false,
						Boards:      []string{},
						IsTask:      false,
						IsArchive:   false,
						Priority:    0,
					},
					IsComplete: true,
					InProgress: false,
				},
			},
			wantId:   1,
			wantIcon: setColor("", "green", false),
			wantDesc: gray("nice"),
		},
		{
			name: "incomplete and prio",
			args: args{
				item: &Task{
					BaseItem: BaseItem{
						ID:          1,
						Date:        "",
						Timestamp:   0,
						Description: "nice",
						IsStarred:   false,
						Boards:      []string{},
						IsTask:      false,
						IsArchive:   false,
						Priority:    2,
					},
					IsComplete: true,
					InProgress: false,
				},
			},
			wantId:   1,
			wantIcon: setColor("", "green", false),
			wantDesc: gray("nice"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotIcon, gotDesc := buildMessage(tt.args.item)
			if gotId != tt.wantId {
				t.Errorf("buildMessage() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotIcon != tt.wantIcon {
				t.Errorf("buildMessage() gotIcon = %v, want %v", gotIcon, tt.wantIcon)
			}
			if gotDesc != tt.wantDesc {
				t.Errorf("buildMessage() gotDesc = %v, want %v", gotDesc, tt.wantDesc)
			}
		})
	}
}

func Test_getAge(t *testing.T) {
	type args struct {
		item Item
	}
	tests := []struct {
		name    string
		args    args
		wantAge float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAge := getAge(tt.args.item); gotAge != tt.wantAge {
				t.Errorf("getAge() = %v, want %v", gotAge, tt.wantAge)
			}
		})
	}
}

func Test_getStats(t *testing.T) {
	type args struct {
		items []Item
	}
	tests := []struct {
		name        string
		args        args
		wantComp    int
		wantInprog  int
		wantPending int
		wantNotes   int
	}{
		{
			name: "0 everything",
			args: args{
				items: []Item{},
			},
			wantComp:    0,
			wantInprog:  0,
			wantPending: 0,
			wantNotes:   0,
		},
		{
			name: "1 task",
			args: args{
				items: []Item{
					&Task{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   false,
							Priority:    0,
						},
						IsComplete: false,
						InProgress: false,
					},
				},
			},
			wantComp:    0,
			wantInprog:  0,
			wantPending: 1,
			wantNotes:   0,
		},
		{
			name: "1 task complete",
			args: args{
				items: []Item{
					&Task{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   false,
							Priority:    0,
						},
						IsComplete: true,
						InProgress: false,
					},
				},
			},
			wantComp:    1,
			wantInprog:  0,
			wantPending: 0,
			wantNotes:   0,
		},
		{
			name: "1 task inprog",
			args: args{
				items: []Item{
					&Task{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   false,
							Priority:    0,
						},
						IsComplete: false,
						InProgress: true,
					},
				},
			},
			wantComp:    0,
			wantInprog:  1,
			wantPending: 0,
			wantNotes:   0,
		},
		{
			name: "1 note",
			args: args{
				items: []Item{
					&Note{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   false,
							Priority:    0,
						},
					},
				},
			},
			wantComp:    0,
			wantInprog:  0,
			wantPending: 0,
			wantNotes:   1,
		},
		{
			name: "1 note, 1 task",
			args: args{
				items: []Item{
					&Note{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   false,
							Priority:    0,
						},
					},
					&Task{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   false,
							Priority:    0,
						},
						IsComplete: false,
						InProgress: false,
					},
				},
			},
			wantComp:    0,
			wantInprog:  0,
			wantPending: 1,
			wantNotes:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotComp, gotInprog, gotPending, gotNotes := getStats(tt.args.items)
			if gotComp != tt.wantComp {
				t.Errorf("getStats() gotComp = %v, want %v", gotComp, tt.wantComp)
			}
			if gotInprog != tt.wantInprog {
				t.Errorf("getStats() gotInprog = %v, want %v", gotInprog, tt.wantInprog)
			}
			if gotPending != tt.wantPending {
				t.Errorf("getStats() gotPending = %v, want %v", gotPending, tt.wantPending)
			}
			if gotNotes != tt.wantNotes {
				t.Errorf("getStats() gotNotes = %v, want %v", gotNotes, tt.wantNotes)
			}
		})
	}
}

func Test_getIcon(t *testing.T) {
	type args struct {
		item Item
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIcon(tt.args.item); got != tt.want {
				t.Errorf("getIcon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStar(t *testing.T) {
	type args struct {
		item Item
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStar(tt.args.item); got != tt.want {
				t.Errorf("getStar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarkOrUnmarkAttribute(t *testing.T) {
	type args struct {
		mark          []string
		unmark        []string
		markAttrMsg   string
		unmarkAttrMsg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarkOrUnmarkAttribute(tt.args.mark, tt.args.unmark, tt.args.markAttrMsg, tt.args.unmarkAttrMsg); got != tt.want {
				t.Errorf("MarkOrUnmarkAttribute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarkAttribute(t *testing.T) {
	type args struct {
		ids     []string
		attrMsg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarkAttribute(tt.args.ids, tt.args.attrMsg); got != tt.want {
				t.Errorf("MarkAttribute() = %v, want %v", got, tt.want)
			}
		})
	}
}
