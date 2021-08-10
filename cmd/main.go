package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/meagar/sorts"
	"github.com/meagar/sorts/game"
)

type SortCase struct {
	Name string
	fn   func(sorts.SortInput)
}

func main() {
	ebiten.SetWindowTitle("Sort Demo")
	ebiten.SetWindowSize(1024, 768)
	ebiten.RunGame(&game.Game{
		Algs: []*game.Alg{
			{
				Name:     "BubbleSort",
				Fn:       sorts.BubbleSort,
				DefaultN: 25,
				Sleep:    5,
			},
			{
				Name:     "QuickSort",
				Fn:       sorts.QuickSort,
				DefaultN: 250,
				Sleep:    3,
			},
		},
	})
}

type IntSorter struct {
	cmps  int
	swaps int

	ints []int
}

// var _ sorts.SortInput = &IntSorter{}

// func (i *IntSorter) Length() int {
// 	return len(i.ints)
// }

// func (i *IntSorter) CmpLE(a, b int) bool {
// 	i.cmps++
// 	return i.ints[a] <= i.ints[b]
// }
// func (i *IntSorter) CmpGE(a, b int) bool {
// 	i.cmps++
// 	return i.ints[a] >= i.ints[b]
// }

// func (i *IntSorter) Swap(a, b int) {
// 	i.swaps++
// 	i.ints[a], i.ints[b] = i.ints[b], i.ints[a]
// }

// func (i *IntSorter) At(idx int) int {
// 	return i.ints[idx]
// }

// func (i *IntSorter) Done() {
// 	return
// }
