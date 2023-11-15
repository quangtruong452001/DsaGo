package list

// IList defines a generic list interface in Go.
type IList interface {
	// Add appends an item "e" to the list.
	Add(e interface{})

	/*
	* AddAtIndex inserts an item "e" at location "index".
	* Location is an integer started from 0.
	 */
	AddAtIndex(index int, e interface{}) error

	/*
	* RemoveAt removes the item at location "index".
	* It returns the item stored at the index or an error if the index is invalid.
	 */
	RemoveAt(index int) (interface{}, error)

	/*
	* RemoveItem removes an item from the list.
	* It returns true if the item was found and removed; otherwise, it returns false.
	* The removeItemData function can be provided to delete the item's data stored in the list.
	 */
	RemoveItem(item interface{}, removeItemData func(interface{})) bool

	// Empty returns true if the list is empty; otherwise, it returns false.
	Empty() bool

	// Size returns the number of items stored in the list.
	Size() int

	// Clear makes the list empty by clearing all data and putting the list to the initial condition.
	Clear()

	/*
	* Get returns a reference to the item at location "index".
	* It returns an error if the index is invalid.
	 */
	Get(index int) (interface{}, error)

	/*
	* IndexOf returns the index of an item in the list.
	* If the item is not found, it returns -1.
	 */
	IndexOf(item interface{}) int

	// Contains returns true if the list contains the given item; otherwise, it returns false.
	Contains(item interface{}) bool

	/*
	* ToString returns a string describing the list.
	* item2str function can be provided to convert each item to a string.
	 */
	ToString(item2str *interface{}) string
}
