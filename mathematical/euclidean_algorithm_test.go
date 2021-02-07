package mathematical

import "testing"

func TestEuclidean(t *testing.T) {
	if got := Euclidean(60, 24); got != 12 {
		t.Errorf("Expected 12, got %d", got)
	}
	if got := Euclidean(55, 15); got != 5 {
		t.Errorf("Expected 5, got %d", got)
	}
}
func TestEuclideanRecursive(t *testing.T) {
	if got := EuclideanRecursive(60, 24); got != 12 {
		t.Errorf("Expected 12, got %d", got)
	}
	if got := EuclideanRecursive(55, 15); got != 5 {
		t.Errorf("Expected 5, got %d", got)
	}
}

func BenchmarkEuclidean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Euclidean(540205, 40321)
	}
}
func BenchmarkEuclideanRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EuclideanRecursive(540205, 40321)
	}
}
