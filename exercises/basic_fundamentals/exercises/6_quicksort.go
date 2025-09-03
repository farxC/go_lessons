package exercises

/*
Lumoto partition scheme

	Explanation: Pick the last element as a pivot and parts into small pieces
*/
func partition(arr []int, low int, high int) int {

	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func Quicksort(arr []int, low int, high int) []int {
	if low < high {
		pivot_idx := partition(arr, low, high)

		left := Quicksort(arr, low, pivot_idx-1)   // Left
		right := Quicksort(arr, pivot_idx+1, high) // Right

		return append(append(left, arr[pivot_idx]), right...)
	}
	return arr[low : high+1]
}
