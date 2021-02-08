package sorting

import "testing"

func TestMergeSort(t *testing.T) {
	arr := [...]float64{54, 26, 93, 17, 77, 31, 44, 55, 20}
	t.Logf("%v", arr)
	t.Logf("%v", MergeSort(arr[:]))
}
