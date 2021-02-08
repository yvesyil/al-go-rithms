package datastructures

import "fmt"

/* LinkedList
* !TODO
* Sort() *List
* Split() (*Node, *Node)
* Merge(n1 *Node, n2 *Node)
 */

// Node is the linked list node.
type Node struct {
	data interface{}
	next *Node
}

func newNode(data interface{}) *Node {
	return &Node{data: data}
}

// List is the linked list struct.
type List struct {
	head *Node
	tail *Node
}

// NewLinkedList returns a reference to new nil linked list.
func NewLinkedList() *List {
	return &List{}
}

// Traverse returns a string representation of the linked list.
func (l *List) Traverse() string {
	var s string
	for np := l.head; np != nil; np = np.next {
		s += fmt.Sprintf("%v -> ", np.data)
	}
	return s
}

// Add takes a data, creates a new node and appends to the linked list, returns a reference to node.
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

// Remove removes the given node from the list, if the node is not on the list, returns an error.
func (l *List) Remove(n *Node) error {
	np := l.head
	for np != n && np != nil {
		np = np.next
	}
	if np == nil {
		return fmt.Errorf("Could not find node %v", n.data)
	}
	np.next = n.next
	n.next = nil
	return nil
}

// Search iterates the linked list, looking for a match with the given data.
// If it finds a match, returns the match with true, otherwise, returns nil with false.
func (l *List) Search(data interface{}) (*Node, bool) {
	np := l.head
	for np.data != data && np != nil {
		np = np.next
	}
	if np == nil {
		return nil, false
	}
	return np, true
}

// Update takes a node and the new data, swaps the node's data with the given one.
func (l *List) Update(n *Node, data interface{}) *Node {
	n.data = data
	return n
}

// Copy returns a new copy of the linked list.
func (l *List) Copy() *List {
	newList := NewLinkedList()
	for np := l.head; np != nil; np = np.next {
		newList.Add(np.data)
	}
	return newList
}

// Reverse returns a reversed version of the linked list
// example,
// before: a -> b -> c -> d -> nil;
// after: d -> c -> b -> a -> nil;
func (l *List) Reverse() (*List, bool) {
	if l.head == nil || l.head.next == nil {
		return nil, false
	}

	newList := l.Copy()

	var prevNp *Node = nil
	np := newList.head
	newList.tail = np
	nextNp := np.next

	for nextNp != nil {
		np.next = prevNp
		prevNp = np
		np = nextNp
		nextNp = nextNp.next
	}
	np.next = prevNp
	newList.head = np

	return newList, true
}
