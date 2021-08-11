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
				DefaultN: 30,
				Sleep:    0,
			},
			{
				Name:     "QuickSort",
				Fn:       sorts.QuickSort,
				DefaultN: 100,
				Sleep:    0,
			},
			{
				Name:     "QuickSort",
				Fn:       sorts.QuickSort,
				DefaultN: 500,
				Sleep:    0,
			},
		},
	})
}

type IntSorter struct {
	cmps  int
	swaps int

	ints []int
}
