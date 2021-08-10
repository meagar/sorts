package sorts

import "math"

type SortInput interface {
	// Length should return the number of elements in the set
	Length() int

	// Swap should swap the values at the two indexes
	Swap(idx1, idx2 int)

	// At should return the element at the given index
	At(idx int) int

	Done()
}

func BubbleSort(s SortInput) {
	for n := s.Length() - 1; n >= 0; n-- {
		for idx := 0; idx < n; idx++ {
			if s.At(idx) >= s.At(idx+1) {
				s.Swap(idx, idx+1)
			}
		}
	}
	s.Done()
}

func QuickSort(s SortInput) {
	quicksort(s, 0, s.Length()-1)
	s.Done()
}

// Translated from https://en.wikipedia.org/wiki/Quicksort
func quicksort(s SortInput, lo, hi int) {
	if lo >= 0 && hi >= 0 {
		if lo < hi {
			p := quicksort_partition(s, lo, hi)
			quicksort(s, lo, p)
			quicksort(s, p+1, hi)
		}
	}
}

func quicksort_partition(s SortInput, lo, hi int) int {
	// We save the value of the pivot.
	// This is important: we do not need the index,
	// but the value, because the position of the
	// index will change over time, so we depend on its
	// value, not its index (...I mean position)
	// pivot := s.At(int(math.Floor(float64(hi+lo) / 2.0)))
	pivot := s.At(int(math.Floor(float64(hi+lo) / 2.0)))
	i := lo - 1
	j := hi + 1

	for {
		// we start from the beginning and if we
		// find a value grater than the pivot
		// (one that should be in the right, not
		// in the left), execute the next instruction
		// do
		// 	i := i + 1
		// while A[i] < pivot
		for {
			i = i + 1
			// if s.Cmp((i) >= pivot {
			if s.At(i) >= pivot {
				break
			}
		}

		// we start from the end and if we
		// find a value smaller than the pivot
		// (one that should be in the left, not
		// in the right)...
		// do
		// 	j := j - 1
		// while A[j] > pivot
		for {
			j = j - 1
			if s.At(j) <= pivot {
				break
			}
		}

		// At this point, everything before i is < pivot,
		// and everything after j is > pivot; if i and j cross,
		// we found the pivot index and the partitioning is done.
		// if i ≥ j then
		// 	return j
		if i >= j {
			return j
		}

		// A[i] is now >= pivot, and A[j] is <= pivot (with
		// the pivot somewhere between i and j, included),
		// So we swap them, as we want the opposite
		// swap A[i] with A[j]
		s.Swap(i, j)
	}
}
