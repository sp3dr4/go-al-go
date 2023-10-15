package quicksort

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)

	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	idx := lo - 1
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx += 1
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}
	idx += 1
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func QuickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}
