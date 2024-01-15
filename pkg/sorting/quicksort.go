package sorting

// QuickSort is an in-place, popular sorting algorithm. This implementation uses the Hoare partition scheme.
//
// → https://en.wikipedia.org/wiki/Quicksort
func QuickSort[T any](array []T, compare func(a, b T) int) {
	quicksort(array, 0, len(array)-1, compare)
}

func quicksort[T any](array []T, lo, hi int, compare func(a, b T) int) {
	if lo >= 0 && hi >= 0 && lo < hi {
		p := partition(array, lo, hi, compare)
		quicksort(array, lo, p, compare)
		quicksort(array, p+1, hi, compare)
	}
}

func partition[T any](array []T, lo, hi int, compare func(a, b T) int) int {
	// Choose the first element as the pivot
	pivot := array[lo]

	// Left index
	i := lo - 1

	// Right index
	j := hi + 1

	for {
		// Move the left index to the right while the element at
		// the left index is less than the pivot
		for {
			i = i + 1
			if compare(array[i], pivot) >= 0 {
				break
			}
		}

		// Move the right index to the left while the element at
		// the right index is greater than the pivot
		for {
			j = j - 1
			if compare(array[j], pivot) <= 0 {
				break
			}
		}

		// If the indices crossed, return
		if i >= j {
			return j
		}

		// Swap the elements at the left and right indices
		tmp := array[i]
		array[i] = array[j]
		array[j] = tmp
	}
}
