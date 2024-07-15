package tb

import (
	"github.com/gookit/color"
	"sort"
)

var ind = "  "

// groupByDate groups items by date
func (b *Book) groupByDate(a bool) (map[string][]int, []string) {
	hd := make(map[string][]int)
	var kd []string

	for _, item := range b.items {
		bi := item.GetBaseItem()
		if a || !bi.IsArchive {
			hd[bi.Date] = append(hd[bi.Date], bi.ID)
		}
	}

	for d := range hd {
		kd = append(kd, d)
	}

	//sorts by first element in []int w.r.t date
	sort.Slice(kd, func(i, j int) bool {
		return hd[kd[i]][0] < hd[kd[j]][0]
	})

	return hd, kd
}

// DisplayByDate displays items grouped by date
func (b *Book) DisplayByDate(a bool) {
	hd, kd := b.groupByDate(a)

	for _, d := range kd {
		print(d)
		print("\n")
		for _, t := range hd[d] {
			print(t)
			print("\n")
		}
	}
}

// groupByBoard groups items by board
func (b *Book) groupByBoard() (map[string][]Item, []string) {
	hb := make(map[string][]Item)
	var kb []string

	for _, item := range b.items {
		bi := item.GetBaseItem()
		if !bi.IsArchive {
			for _, board := range bi.Boards {
				hb[board] = append(hb[board], item)
			}
		}
	}

	for board := range hb {
		kb = append(kb, board)
	}

	// sorts by first element in []int w.r.t board
	sort.Slice(kb, func(i, j int) bool {
		it1 := hb[kb[i]][0]
		it2 := hb[kb[j]][0]
		return it1.GetBaseItem().ID < it2.GetBaseItem().ID
	})

	return hb, kb
}

// DisplayByBoard displays items grouped by board
func (b *Book) DisplayByBoard() {
	hb, kb := b.groupByBoard()

	lightwhite := color.New(color.FgLightWhite, color.OpUnderscore).Render
	grey := color.FgGray.Render

	color.Printf("\n")
	//TODO
	for _, board := range kb {
		color.Printf(ind+"%s\n", lightwhite(board))
		for _, t := range hb[board] {
			id := t.GetBaseItem().ID
			desc := t.GetBaseItem().Description
			color.Printf(ind+" %s%s"+"%s\n", grey(id), grey("."), desc)
		}
	}
	color.Printf("\n")
}

func (b *Book) displayStats() {

}
