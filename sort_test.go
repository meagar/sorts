package sorts_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/meagar/sorts"
)

type IntSorter struct {
	t     *testing.T
	cmps  int
	swaps int

	ints []int
}

var _ sorts.SortInput = &IntSorter{}

func (i *IntSorter) Length() int {
	return len(i.ints)
}

func (i *IntSorter) Swap(a, b int) {
	i.t.Helper()
	i.t.Logf("swap %d (%d), %d (%d)", a, i.ints[a], b, i.ints[b])
	i.swaps++
	i.ints[a], i.ints[b] = i.ints[b], i.ints[a]
	i.t.Log(i.ints)
}

func (i *IntSorter) At(idx int) int {
	return i.ints[idx]
}

func (i *IntSorter) Done() {
	return
}

func TestBubbleSort(t *testing.T) {

	t.Run("It sorts", func(t *testing.T) {
		input := fill(5)

		want := copyInts(input)
		sort.Ints(want)

		got := IntSorter{
			t:    t,
			ints: copyInts(input),
		}

		sorts.BubbleSort(&got)

		if !match(got.ints, want) {
			t.Errorf("BubbleSort(%v):\n\tGot  %v,\n\twant %v", input, got.ints, want)
		}

	})

	var bubbleSortCases = []struct {
		input []int
		cmps  int
		swaps int
	}{
		{[]int{1, 2}, 1, 0},
		{[]int{2, 1}, 1, 1},
		{[]int{1, 2, 3}, 3, 0},
		{[]int{3, 2, 1}, 3, 3},
		{[]int{5, 2, 4, 3, 1}, 10, 8},
		{[]int{1, 2, 3, 4, 5}, 10, 0},
		{[]int{5, 4, 3, 2, 1}, 10, 10},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 45, 0},
		{[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 45, 20},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 45, 45},
	}

	for _, tc := range bubbleSortCases {
		t.Run(fmt.Sprintf("BubbleSort(%v)", tc.input), func(t *testing.T) {
			want := copyInts(tc.input)
			sort.Ints(want)

			sorter := IntSorter{
				t:    t,
				ints: copyInts(tc.input),
			}

			sorts.BubbleSort(&sorter)

			if !match(sorter.ints, want) {
				t.Errorf("BubbleSort(%v): Got %v, want %v", tc.input, sorter.ints, want)
			}

			if sorter.cmps != tc.cmps {
				t.Errorf("BubbleSort(%v): Completed with %d comparisons, expected %d", tc.input, sorter.cmps, tc.cmps)
			}
			if sorter.swaps != tc.swaps {
				t.Errorf("BubbleSort(%v): Completed with %d swaps, expected %d", tc.input, sorter.swaps, tc.swaps)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	t.Run("It sorts", func(t *testing.T) {
		input := fill(500)

		want := copyInts(input)
		sort.Ints(want)

		got := IntSorter{
			t:    t,
			ints: copyInts(input),
		}

		sorts.QuickSort(&got)

		if !match(got.ints, want) {
			t.Errorf("QuickSort(%v):\n\tGot  %v,\n\twant %v", input, got.ints, want)
		}

	})

	var testCases = []struct {
		input []int
		cmps  int
		swaps int
	}{
		{[]int{1, 2}, 3, 0},
		{[]int{2, 1}, 4, 1},
		{[]int{1, 2, 3}, 7, 0},
		{[]int{3, 2, 1}, 7, 1},
		{[]int{5, 2, 4, 3, 1}, 17, 2},
		{[]int{1, 2, 3, 4, 5}, 16, 0},
		{[]int{5, 4, 3, 2, 1}, 16, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 43, 0},
		{[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 55, 8},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 44, 5},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("QuickSort(%v)", tc.input), func(t *testing.T) {
			want := copyInts(tc.input)
			sort.Ints(want)

			sorter := IntSorter{
				t:    t,
				ints: copyInts(tc.input),
			}

			sorts.QuickSort(&sorter)

			if !match(sorter.ints, want) {
				t.Errorf("QuickSort(%v): Got %v, want %v", tc.input, sorter.ints, want)
			}

			if sorter.cmps != tc.cmps {
				t.Errorf("QuickSort(%v): Completed with %d comparisons, expected %d", tc.input, sorter.cmps, tc.cmps)
			}
			if sorter.swaps != tc.swaps {
				t.Errorf("QuickSort(%v): Completed with %d swaps, expected %d", tc.input, sorter.swaps, tc.swaps)
			}
		})
	}
}

func match(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, numA := range a {
		if numA != b[idx] {
			return false
		}
	}
	return true
}

func copyInts(ints []int) (out []int) {
	out = make([]int, len(ints))
	copy(out, ints)
	return
}

func fill(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
