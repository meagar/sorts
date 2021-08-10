package main

import (
	"fmt"
	"math/rand"

	"github.com/meagar/sorts"
)

func main() {
	input := fill(10000)
	fmt.Println("Input Length: ", len(input))

	qsInput := make([]int, len(input))
	bsInput := make([]int, len(input))

	copy(qsInput, input)
	copy(bsInput, input)

	qsSorter := IntSorter{0, 0, qsInput}
	bsSorter := IntSorter{0, 0, bsInput}

	fmt.Println("Bubble Sort")
	sorts.BubbleSort(&bsSorter)
	fmt.Println("Comparisons:", bsSorter.cmps)
	fmt.Println("Swaps:", bsSorter.swaps)

	fmt.Println("Quick Sort")
	sorts.QuickSort(&qsSorter)
	fmt.Println("Comparisons:", qsSorter.cmps)
	fmt.Println("Swaps:", qsSorter.swaps)

}

func fill(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

type IntSorter struct {
	cmps  int
	swaps int

	ints []int
}

var _ sorts.SortInput = &IntSorter{}

func (i *IntSorter) Length() int {
	return len(i.ints)
}

func (i *IntSorter) Cmp(a, b int) bool {
	i.cmps++
	return i.ints[a] <= i.ints[b]
}

func (i *IntSorter) Swap(a, b int) {
	i.swaps++
	i.ints[a], i.ints[b] = i.ints[b], i.ints[a]
}

func (i *IntSorter) At(idx int) int {
	return i.ints[idx]
}
