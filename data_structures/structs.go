package data_structures

// Linked List Implementation

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (ll *LinkedList) Add(elem int) {
	new_node := &Node{elem, nil}
	if ll.head == nil && ll.head == nil {
		ll.head = new_node
	} else {
		ll.tail.next = new_node
	}
	ll.tail = new_node
	ll.length++
}

func (ll LinkedList) Search(elem int) *Node {
	for walker := ll.head; walker != nil; walker = walker.next {
		if walker.value == elem {
			return walker
		}
	}
	return nil
}
