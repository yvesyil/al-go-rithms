package datastructures

type Heap []int

func Heapify(arr []int) Heap {
	n := len(arr) - 1
	if n <= 1 {
		return Heap(arr)
	}
	for i := n / 2; i > 0; i-- {
		k := i
		v := arr[k]
		heap := false
		for !heap && 2*k <= n {
			j := 2 * k
			if j < n {
				if arr[j] < arr[j+1] {
					j++
				}
			}
			if v >= arr[j] {
				heap = true
			} else {
				arr[k] = arr[j]
				k = j
			}
		}
		arr[k] = v
	}
	return Heap(arr)
}
