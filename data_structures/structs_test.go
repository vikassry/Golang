package data_structures

import (
	"testing"
)

func TestContainsElementsAddedToLinkedList(t *testing.T) {
	linkedList := new(LinkedList)
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)

	for i := 1; i <= 3; i++ {
		if linkedList.Search(i) == nil {
			t.Errorf("Should have contained a node with value %d but got %v", i, linkedList.Search(i))
		}
	}
}
