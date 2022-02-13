package main

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	arr = QuickSort(arr, k, 0, len(arr)-1)
	return arr[:k]
}

func QuickSort(arr []int, k, l, r int) []int {
	if len(arr) == 0 {
		return arr
	}
	value := arr[l]
	i, j := l, r
	for i < j {
		for i < j && arr[j] >= value { // 必须先j--
			j--
		}
		for i < j && arr[i] <= value {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[l] = arr[l], arr[i]
	if i > k {
		return QuickSort(arr, k, l, i-1)
	} else if i < k {
		return QuickSort(arr, k, i+1, r)
	} else {
		return arr[:k]
	}
}
