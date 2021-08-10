package sorts

type SortInput interface {
	// Length should return the number of elements in the set
	Length() int

	// Cmp should return true if the element at idx1 is less than than or equal to the element at idx2
	Cmp(idx1, idx2 int) bool

	// Swap should swap the values at the two indexes
	Swap(idx1, idx2 int)

	// At should return the element at the given index
	At(idx int) int
}

func BubbleSort(s SortInput) {
	for n := s.Length() - 1; n >= 0; n-- {
		for idx := 0; idx < n; idx++ {
			if !s.Cmp(idx, idx+1) {
				s.Swap(idx, idx+1)
			}
		}
	}
}

func QuickSort(s SortInput) {
	quicksort(s, 0, s.Length()-1)
}

// The recursive component of quicksort
func quicksort(s SortInput, lo, hi int) {
	if lo >= 0 && hi >= 0 {
		if lo < hi {
			p := quicksort_partition(s, lo, hi)
			quicksort(s, lo, p-1)
			quicksort(s, p+1, hi)
		}
	}
}

func quicksort_partition(s SortInput, lo, hi int) int {
	i := lo - 1

	for j := lo; j <= hi; j++ {
		if s.Cmp(j, hi) {
			i++
			s.Swap(i, j)
		}
	}

	return i
}
