package hashset

import (
	"fmt"
	"math/rand"
	"time"
)

// node (Private) - structure of each node in the linked set.
type node struct {
	data string
	next *node
}

// string (private) - Returns the data stored in a node as a string
func (n *node) string() string {
	return fmt.Sprintf(" %v", n.data)
}

// HashSet (Public) - The structure for our hashset, Since size is always
// static because a hashset is just a huge chunk of memory, I needed to
// also include the total count of items in the hashset. Collection is
// just an array of pointers to each linked set of items that hashed to
// the same value
type HashSet struct {
	items      int
	size       int
	collection []*node
}

// Init (Public) - initializes the array with whatever size is provided, This is what can be overrided by the user.
func (h *HashSet) Init(capacity int) *HashSet {
	if capacity < 0 {
		return nil
	}
	h.collection = make([]*node, capacity)
	h.size = capacity
	return h
}

// New (Public) - Returns an initialized array with default size of 10.
func New() *HashSet { return new(HashSet).Init(19) }

// Add (Public) - adds a node to the collection, also if there is
// multiple nodes at an index, it adds it to the begining for speed
func (h *HashSet) Add(item string) bool {
	err := checkForNil(item)
	if err != nil {
		return false
	}

	index := h.getIndex(hashitem(item))
	newest := &node{data: item}
	newest.next = h.collection[index]
	h.collection[index] = newest
	h.items++
	return true
}

// RemoveItem (Public) - Returns true or false if the item provided was removed.
func (h *HashSet) RemoveItem(item string) bool {
	err := checkForNil(item)
	if err != nil {
		return false
	}
	if !h.Contains(item) {
		return false
	}
	index := h.getIndex(hashitem(item))
	current := h.collection[index]
	if current.data == item {
		h.collection[index] = h.collection[index].next
		h.items--
		return true
	}
	for {
		if current.next.data == item {
			current.next = current.next.next
			h.items--
			break
		}
		current = current.next
	}
	return true
}

// Contains (Public) - Gets the index, then searches the linked set
// for the item and returns true or false if found.
func (h *HashSet) Contains(item string) bool {
	err := checkForNil(item)
	if err != nil {
		return false
	}

	index := h.getIndex(hashitem(item))
	current := h.collection[index]
	//fmt.Println(index)
	for current != nil {
		if current.data == item {
			return true
		}
		current = current.next
	}
	return false
}

// Remove (Public) - Finds a random index and returns the first removed node
func (h *HashSet) Remove() string {
	if h.items == 0 {
		return ""
	}
	var random int
	for {
		rand.Seed(time.Now().UTC().UnixNano())
		random := rand.Intn(h.size)
		if h.collection[random].data != "" {
			break
		}
	}
	removed := h.collection[random].data
	h.collection[random] = h.collection[random].next
	h.items--
	return removed
}

// Get (Public) - gets a random index and returns the first node string
func (h *HashSet) Get() string {
	if h.items == 0 {
		return ""
	}
	var random int
	for {
		rand.Seed(time.Now().UTC().UnixNano())
		random := rand.Intn(h.size)
		if h.collection[random].data != "" {
			break
		}
	}
	return h.collection[random].data
}

// Size (Public) - Returns the size of the array
func (h *HashSet) Size() int {
	return h.items
}

// Items (Public) - Returns the size of the HashSet
func (h *HashSet) Items() int {
	return h.items
}

// String (Public) - formats the map as a string
func (h *HashSet) String() string {
	if h.size == 0 {
		return "[]"
	}
	s := ""
	for i := 0; i < h.size; i++ {
		current := h.collection[i]
		if current == nil {
			continue
		}
		s += fmt.Sprintf("%v\t%v ", i, current.data)
		for current.next != nil {
			//fmt.Println(current.next.data)
			s += current.next.data + " "
			current = current.next
		}
		s += "\n"
	}
	return s
}

// checkForNil (private) - check if the key or the value is nil
func checkForNil(item string) error {
	if item == "" {
		return fmt.Errorf("Item can't be nil")
	}
	return nil
}

type hashitem string

// getIndex (Public) - Gets a hash of the string and makes
// sure it is in bounds with the size of the hashset
func (h *HashSet) getIndex(item hashitem) int {
	result := item.hashCode() % h.size
	if result < 0 {
		result += h.size
	}
	return result
}

// pow (Private) - Golang doesn't have a exponent function that deals
// with integers, so I made my own.
func pow(x, y int) int {
	if y < 0 {
		return -1
	}
	if y == 0 {
		return 1
	}
	result := x
	for i := 1; i < y; i++ {
		result *= x
	}
	return result
}

// hashCode (Private) - Golang implementation of the hashCode function
// in java. Golang has many built in hashing functions, just wanted to
// port over the one in the lab for consistency.
func (item hashitem) hashCode() int {
	size := len(item) - 1
	prime := 31

	result := 0
	for i := 0; i < len(item); i++ {
		result += int(item[i]) * pow(prime, size)
		size--
	}
	return result
}
