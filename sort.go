package sorts

import "math"

type SortInput interface {
	// Length should return the number of elements in the set
	Length() int

	// Swap should swap the values at the two indexes
	Swap(idx1, idx2 int)

	// At should return the element at the given index
	At(idx int) int

	// ResetIteration should be called at the top of each primary iteration
	ResetIteration()

	// Done should be called when the algorithm is finished sorting
	Done()
}

func BubbleSort(s SortInput) {
	for n := s.Length() - 1; n >= 0; n-- {
		for idx := 0; idx < n; idx++ {
			s.ResetIteration()
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

// https://en.wikipedia.org/wiki/Quicksort
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
		s.ResetIteration()
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

// https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort(s SortInput) {
	// i ← 1
	// while i < length(A)
	for i := 1; i < s.Length(); i++ {
		s.ResetIteration()
		// j ← i
		// while j > 0 and A[j-1] > A[j]
		for j := i; j > 0 && s.At(j-1) > s.At(j); j-- {
			// 	swap A[j] and A[j-1]
			s.Swap(j, j-1)
			// j ← j - 1
			// end while
		}
		// i ← i + 1
		// end while
	}
	s.Done()
}

func SelectionSort(s SortInput) {
	// /* a[0] to a[aLength-1] is the array to sort */
	// int i,j;
	// int aLength; // initialise to a's length
	// /* advance the position through the entire array */
	// /*   (could do i < aLength-1 because single element is also min element) */
	// for (i = 0; i < aLength-1; i++)
	// {
	for i := 0; i < s.Length()-1; i++ {
		s.ResetIteration()
		//     /* find the min element in the unsorted a[i .. aLength-1] */

		//     /* assume the min is the first element */
		//     int jMin = i;
		jMin := i
		//     /* test against elements after i to find the smallest */
		//     for (j = i+1; j < aLength; j++)
		//     {
		for j := i + 1; j < s.Length(); j++ {
			//         /* if this element is less, then it is the new minimum */
			//         if (a[j] < a[jMin])
			//         {
			if s.At(j) < s.At(jMin) {
				//             /* found new minimum; remember its index */
				//     jMin = j;
				jMin = j
				//         }
				//     }
			}
		}
		//     if (jMin != i)
		//     {
		//         swap(a[i], a[jMin]);
		//     }

		if jMin != i {
			s.Swap(i, jMin)
		}
		// }
	}
	s.Done()
}
