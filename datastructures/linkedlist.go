package datastructures

import "fmt"

type LinkedList interface {
	Traverse() string
	Add(data interface{}) *Node
	Remove(n *Node) error
	Search(data interface{}) (*Node, bool)
	Update(n *Node, data interface{}) *Node
	Sort() *List
	Split() (*Node, *Node)
	Merge(n1 *Node, n2 *Node)
}

type Node struct {
	data interface{}
	next *Node
}

func newNode(data interface{}) *Node {
	return &Node{data: data}
}

type List struct {
	head *Node
	tail *Node
}

func NewLinkedList() *List {
	return &List{}
}

func (l *List) Traverse() string {
	var s string
	for np := l.head; np != nil; np = np.next {
		s += fmt.Sprintf("%v -> ", np.data)
	}
	return s
}

func (l *List) Add(data interface{}) *Node {
	n := newNode(data)
	if l.head == nil {
		l.head = n
	} else if l.head.next == nil {
		l.head.next = n
		l.tail = n
	} else {
		l.tail.next = n
		l.tail = l.tail.next
	}
	return n
}

func (l *List) Remove(n *Node) error {
	np := l.head
	for np != n && np != nil {
		np = np.next
	}
	if np == nil {
		return fmt.Errorf("Could not node %v", n.data)
	}
	np.next = n.next
	n.next = nil
	return nil
}

func (l *List) Search(data interface{}) (*Node, bool) {
	np := l.head
	for np != n {
		np = np.next
	}
}
