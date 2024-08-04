package tb

import (
	"os"
	"reflect"
	"testing"
)

func TestBook_Delete(t *testing.T) {
	type fields struct {
		Items []Item
	}
	tests := []struct {
		name   string
		fields fields
		id     int
	}{
		{
			name: "delete nothing",
			fields: fields{
				Items: []Item{
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
			id: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Items: tt.fields.Items,
			}
			b.Delete(tt.id)
		})
	}
}

func TestBook_Remove(t *testing.T) {
	type fields struct {
		Items []Item
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "remove one",
			fields: fields{
				Items: []Item{
					&Task{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   true,
							Priority:    0,
						},
						IsComplete: false,
						InProgress: false,
					},
				},
			},
		},
		{
			name: "remove two",
			fields: fields{
				Items: []Item{
					&Note{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   true,
							Priority:    0,
						},
					},
					&Task{
						BaseItem: BaseItem{
							ID:          2,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   true,
							Priority:    0,
						},
						IsComplete: false,
						InProgress: false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Items: tt.fields.Items,
			}
			b.Remove()
		})
	}
}

func TestBook_Store(t *testing.T) {
	os.Setenv("TEST_MODE", "true") // Enable test mode
	defer os.Unsetenv("TEST_MODE") // Clean up environment variable

	type fields struct {
		Items []Item
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "store 1 item",
			fields: fields{
				Items: []Item{
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
			wantErr: false,
		},
		{
			name: "store 2 item",
			fields: fields{
				Items: []Item{
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
					&Task{
						BaseItem: BaseItem{
							ID:          2,
							Date:        "",
							Timestamp:   3,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Items: tt.fields.Items,
			}
			if err := b.Store(); (err != nil) != tt.wantErr {
				t.Errorf("Book.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// cross platform
func Test_getStoragePath(t *testing.T) {
	os.Setenv("TEST_MODE", "true") // Enable test mode
	defer os.Unsetenv("TEST_MODE") // Clean up environment variable

	tests := []struct {
		name string
		want string
	}{
		{
			name: "TEST_MODE",
			want: "test_storage.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStoragePath(); got != tt.want {
				t.Errorf("getStoragePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "TEST_MODE",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(); (err != nil) != tt.wantErr {
				t.Errorf("Book.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBook_Read(t *testing.T) {
	type fields struct {
		Items []Item
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Items: tt.fields.Items,
			}
			if err := b.Read(); (err != nil) != tt.wantErr {
				t.Errorf("Book.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBook_getMaxID(t *testing.T) {
	type fields struct {
		Items []Item
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "length 0",
			fields: fields{
				[]Item{},
			},
			want: 0,
		},
		{
			name: "length 1",
			fields: fields{
				[]Item{
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
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Items: tt.fields.Items,
			}
			if got := b.getMaxID(); got != tt.want {
				t.Errorf("Book.getMaxID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_GetAllID(t *testing.T) {
	type fields struct {
		Items []Item
	}
	tests := []struct {
		name   string
		fields fields
		at     bool
		a      bool
		want   []string
	}{
		{
			name: "no @",
			fields: fields{
				[]Item{
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
			at:   false,
			want: []string{"1"},
		},
		{
			name: "@",
			fields: fields{
				[]Item{
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
			at:   true,
			want: []string{"@1"},
		},
		{
			name: "multiple Items, no @",
			fields: fields{
				[]Item{
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
					&Note{
						BaseItem: BaseItem{
							ID:          2,
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
			at:   false,
			want: []string{"1", "2"},
		},
		{
			name: "multiple Items, yes @",
			fields: fields{
				[]Item{
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
					&Note{
						BaseItem: BaseItem{
							ID:          2,
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
			at:   true,
			want: []string{"@1", "@2"},
		},
		{
			name: "no archive",
			fields: fields{
				[]Item{
					&Task{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   true,
							Priority:    0,
						},
						IsComplete: false,
						InProgress: false,
					},
				},
			},
			a:    false,
			at:   false,
			want: nil,
		},
		{
			name: "yes archive",
			fields: fields{
				[]Item{
					&Task{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   true,
							Priority:    0,
						},
						IsComplete: false,
						InProgress: false,
					},
				},
			},
			at:   false,
			a:    true,
			want: []string{"1"},
		},
		{
			name: "yes archive, yes @",
			fields: fields{
				[]Item{
					&Task{
						BaseItem: BaseItem{
							ID:          1,
							Date:        "",
							Timestamp:   0,
							Description: "",
							IsStarred:   false,
							Boards:      []string{},
							IsTask:      false,
							IsArchive:   true,
							Priority:    0,
						},
						IsComplete: false,
						InProgress: false,
					},
				},
			},
			at:   true,
			a:    true,
			want: []string{"@1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Items: tt.fields.Items,
			}
			if got := b.GetAllID(tt.at, tt.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Book.GetAllID() = %v, want %v", got, tt.want)
			}
		})
	}
}
