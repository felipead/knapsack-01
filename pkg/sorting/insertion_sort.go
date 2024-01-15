package sorting

// InsertionSort is a simple sorting algorithm that builds the final sorted array one item at a time by comparisons.
// It is much less efficient on large lists than more advanced algorithms such as Quicksort, heapsort, or merge sort.
// However, insertion sort provides several advantages:
//
//   - Simple implementation
//   - Efficient for (quite) small data sets. Although the time complexity is O(n²), it is more efficient in practice
//     than most other simple quadratic algorithms such as selection sort or bubble sort.
//   - Adaptive, i.e., efficient for data sets that are already substantially sorted: the time complexity is O(kn)
//     when each element in the input is no more than k places away from its sorted position.
//   - Stable; i.e., does not change the relative order of elements with equal keys
//   - In-place; i.e., only requires a constant amount O(1) of additional memory space
//   - Online; i.e., can sort a list as it receives it
//
// Reference → https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort[T any](array []T, compare func(a, b T) int) {
	for i := range array {
		j := i
		// while j > 0 and A[j-1] > A[j]
		for j > 0 && compare(array[j-1], array[j]) > 0 {
			// swap A[j] and A[j-1]
			array[j], array[j-1] = array[j-1], array[j]
			j--
		}
	}
}
