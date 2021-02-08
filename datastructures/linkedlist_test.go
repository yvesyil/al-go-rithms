package datastructures

import "testing"

func TestLinkedList(t *testing.T) {
	l := NewLinkedList()
	l.Add("a")
	l.Add("b")
	l.Add("c")
	l.Add("d")
	t.Log(l.Traverse())
	lr, _ := l.Reverse()
	t.Log(lr.Traverse())
}
