package sorting

// Quicksort is an efficient, general-purpose sorting algorithm. The algorithm sorts in-place, and is not stable.
//
// This implementation uses the Hoare partition scheme. Hoare's scheme is more efficient than Lomuto's partition
// scheme because it does three times fewer swaps on average. Also, the implementation given creates a balanced
// partition even when all values are equal, which Lomuto's scheme does not. Like Lomuto's partition scheme,
// Hoare's partitioning also would cause Quicksort to degrade to O(n²) for already sorted input, if the pivot was
// chosen as the first or the last element. With the middle element as the pivot, however, sorted data results with
// (almost) no swaps in equally sized partitions leading to best case behavior of Quicksort, i.e. O(n log(n)).
//
//   - Best case time complexity: O(n log n)
//   - Worst case time complexity: O(n²)
//   - Average performance: O(n log n)
//   - Space complexity: O(log n) — it must store a constant amount of information for each nested recursive call
//
// Source → https://en.wikipedia.org/wiki/Quicksort
func Quicksort[T any](array []T, compare func(a, b T) int) {
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
	// Choose the first element as the pivot. This will make the algorithm deteriorate
	// to O(n²) performance if the input is already sorted.
	//
	// This can be prevented by selecting the middle element as follows:
	//
	//   pivot := array[lo+(hi-lo)/2]
	//
	// at the cost of more complex arithmetic.
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
