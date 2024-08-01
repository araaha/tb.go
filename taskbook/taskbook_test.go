package tb

import (
	"reflect"
	"testing"
)

func TestBook_Delete(t *testing.T) {
	type fields struct {
		Items []Item
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Items: tt.fields.Items,
			}
			b.Delete(tt.args.id)
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
		// TODO: Add test cases.
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
	type fields struct {
		Items []Item
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
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
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Create()
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
		want   []string
	}{
		{
			name: "no at",
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
			name: "at",
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
			name: "multiple at and no at",
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
			name: "multiple at and no at",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Book{
				Items: tt.fields.Items,
			}
			if got := b.GetAllID(tt.at); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Book.GetAllID() = %v, want %v", got, tt.want)
			}
		})
	}
}
