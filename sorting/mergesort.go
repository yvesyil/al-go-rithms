package sorting

func merge(l, r []float64) []float64 {
	var m []float64

	i := 0 // left iterator
	j := 0 // right iterator

	for i < len(l) && j < len(r) {
		if l[i] > r[j] {
			m = append(m, r[j])
			j++
		} else {
			m = append(m, l[i])
			i++
		}
	}

	for i < len(l) {
		m = append(m, l[i])
		i++
	}

	for j < len(r) {
		m = append(m, r[j])
		j++
	}

	return m
}

// MergeSort takes a float64 type slice and returns it sorted.
func MergeSort(slice []float64) []float64 {
	size := len(slice)
	if size < 2 {
		return slice
	}
	mid := size / 2

	l := slice[:mid]
	r := slice[mid:]

	l = MergeSort(l)
	r = MergeSort(r)

	return merge(l, r)
}
