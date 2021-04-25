package datastructures

import "testing"

func TestHeapify(t *testing.T) {
	arr := []int{-1, 2, 3, 4, 9, 8, 7, 85, 43, 45, 54, 86, 12, 22, 54, 99}
	h := Heapify(arr)
	t.Logf("%v", h)
}
