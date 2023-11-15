package SLinkedList

import (
	"errors"
	"fmt"
)

/*
SLinkedList structure
head -> node0 -> node1 -> ... <-> tail
*/
type SLinkedList struct {
	head      *Node //this node does not contain user's data
	tail      *Node //this node does not contain user's data
	count     int   //keep number of items stored in the list
	itemEqual func(interface{}, interface{}) bool
	//deleteUserData func(*SLinkedList)
}

type Node struct {
	data interface{}
	next *Node
}

func NewSLinkedList(itemEqual func(lhs, rhs interface{}) bool) *SLinkedList {
	list := &SLinkedList{
		head:      &Node{},
		tail:      &Node{},
		count:     0,
		itemEqual: itemEqual,
	}

	list.head.next = list.tail
	list.tail.next = list.head
	return list
}

// Add appends an item "v" to end of the list.
func (list *SLinkedList) Add(v interface{}) {
	node := &Node{data: v, next: list.tail}

	list.tail.next.next = node // list.head.next
	list.tail.next = node      // list.head
	list.count += 1
}

// CopyFrom other list
func (list *SLinkedList) CopyFrom(otherList *SLinkedList) {
	//Initialize this list to the empty condition
	list.count = 0
	list.head.next = list.tail
	list.tail.next = list.head

	//Copy pointers from "otherList"
	//list.deleteUserData = otherList.deleteUserData
	list.itemEqual = otherList.itemEqual

	//Copy data from "otherList"
	ptr := &Node{
		data: nil,
		next: otherList.head.next,
	}
	for ptr != otherList.tail {
		list.Add(ptr.data)
		ptr = ptr.next
	}
}

// AddAtIndex inserts an item "v" at location "index".
func (list *SLinkedList) AddAtIndex(index int, v interface{}) error {
	if index < 0 || index > list.count {
		return errors.New("index out of range")
	}

	newNode := &Node{data: v, next: nil}

	if index == list.count {
		newNode.next = list.tail
		list.tail.next.next = newNode
		list.tail.next = newNode
	} else if index == 0 {
		newNode.next = list.head.next
		list.head.next = newNode
	} else {
		prevNode := list.head.next
		for i := 0; i < index-1; i++ {
			prevNode = prevNode.next
		}
		newNode.next = prevNode.next
		prevNode.next = newNode
	}

	list.count++
	return nil
}

// RemoveAt removes the item at location "index" return .
func (list *SLinkedList) RemoveAt(index int) (interface{}, error) {
	if index < 0 || index > list.count {
		return nil, errors.New("index out of range")
	}

	var data interface{}

	if index == 0 {
		tmp := list.head.next
		data = tmp.data
		list.head.next = tmp.next
	} else {
		prevNode := list.head.next
		for i := 0; i < index-1; i++ {
			prevNode = prevNode.next
		}
		tmp := prevNode.next
		data = tmp.data
		prevNode.next = tmp.next
	}
	list.count--
	return data, nil
}

// RemoveItem remove item in list has data equal input value
func (list *SLinkedList) RemoveItem(item interface{}) bool { // can be updated with itemEqual function
	find := false
	prevNode := list.head
	currNode := list.head.next

	for currNode != list.tail {
		if list.itemEqual == nil {
			find = currNode.data == item
		} else {
			find = list.itemEqual(currNode.data, item)
		}

		if find {
			if list.tail.next == currNode {
				list.tail.next = prevNode // update tail if needed
			}

			list.count-- // decrease count in list
			break
		}
		currNode = currNode.next
		prevNode = prevNode.next
	}
	return find
}

// Empty returns true if the list is empty; otherwise, it returns false.
func (list *SLinkedList) Empty() bool {
	return list.count == 0
}

// Size returns the number of items stored in the list.
func (list *SLinkedList) Size() int {
	return list.count
}

// Clear makes the list empty by clearing all data and putting the list to the initial condition.
func (list *SLinkedList) Clear() {
	list.head.next = list.tail
	list.tail.next = list.head
	list.count = 0
	//if list.deleteUserData != nil {
	//	list.deleteUserData(list)
	//}
}

// Get returns a reference to the item at location "index".
func (list *SLinkedList) Get(index int) (interface{}, error) {
	if index < 0 || index > list.count {
		return nil, errors.New("index out of range")
	}

	tmp := list.head.next
	for i := 0; i < index; i++ {
		tmp = tmp.next
	}
	return tmp.data, nil
}

// IndexOf returns the index of an item in the list.
func (list *SLinkedList) IndexOf(item interface{}) int { // can be updated with itemEqual function
	tmp := list.head.next
	at := 0
	find := false
	for tmp != list.tail {
		if list.itemEqual != nil {
			find = list.itemEqual(tmp.data, item)
		} else {
			find = tmp.data == item
		}

		if find {
			break
		}
		tmp = tmp.next
		at++
	}
	if find {
		return at
	} else {
		return -1
	}
}

// Contains returns true if the list contains the given item; otherwise, it returns false.
func (list *SLinkedList) Contains(item interface{}) bool {
	return list.IndexOf(item) != -1
}

// ToString returns a string describing the list.
func (list *SLinkedList) ToString(item2str func(interface{}) string) string {
	result := "["
	tmp := list.head.next
	for tmp.next != list.tail {
		if item2str != nil {
			result += item2str(tmp.data) + ", "
		} else {
			result += fmt.Sprintf("%v, ", tmp.data)
		}
		tmp = tmp.next
	}
	if list.count > 0 {
		if item2str != nil {
			result += item2str(tmp.data) + "]"
		} else {
			result += fmt.Sprintf("%v]", tmp.data)
		}
	} else {
		result += "]"
	}
	return result
}

//func (list *SLinkedList) equals(lhs interface{}, rhs interface{}, itemEqual func(interface{}, interface{}) bool) bool {
//	if itemEqual != nil {
//		return itemEqual(lhs, rhs) // compare items value using the provided function
//	} else {
//		return lhs == rhs // compare items address directly for equality
//	}
//}
