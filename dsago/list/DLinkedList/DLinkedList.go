package DLinkedList

import "errors"

/*
DLinkedList structure
head <-> node0 <-> node1 <-> ... <-> tail
*/
type DLinkedList struct {
	head      *Node //this node does not contain user's data
	tail      *Node //this node does not contain user's data
	count     int   //keep number of items stored in the list
	itemEqual func(interface{}, interface{}) bool
	//deleteUserData func(*SLinkedList)
}

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

func NewDLinkedList(itemEqual func(lhs, rhs interface{}) bool) *DLinkedList {
	list := &DLinkedList{
		head:      &Node{},
		tail:      &Node{},
		count:     0,
		itemEqual: itemEqual,
	}

	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

// Add appends an item "v" to end of the list.
func (list *DLinkedList) Add(v interface{}) {
	node := &Node{data: v, next: list.tail, prev: list.tail.prev}

	list.tail.prev.next = node // the last node.next
	list.tail.prev = node      // update the tail.prev
	list.count += 1
}

// CopyFrom other list
func (list *DLinkedList) CopyFrom(otherList *DLinkedList) {
	//Initialize this list to the empty condition
	list.count = 0
	list.head.next = list.tail
	list.tail.prev = list.head

	//Copy pointers from "otherList"
	//list.deleteUserData = otherList.deleteUserData
	list.itemEqual = otherList.itemEqual

	//Copy data from "otherList"
	// initialize a pointer to the first node (head)
	ptr := &Node{
		data: nil,
		next: otherList.head.next,
		prev: nil,
	}
	for ptr != otherList.tail {
		list.Add(ptr.data)
		ptr = ptr.next
	}
}

func (list *DLinkedList) AddAtIndex(index int, v interface{}) error {
	if index < 0 || index > list.count {
		return errors.New("index out of range")
	}

	newNode := &Node(data: v, next:nil, prev: nill)

}
