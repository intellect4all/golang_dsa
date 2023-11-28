package data_structures

// StackADT Stack is a data structure that follows the LIFO (Last In First Out) principle.
type StackADT interface {
	// Push adds an element to the top of the stack.
	Push(interface{})

	// Pop removes and return the top element of the stack.
	Pop() interface{}

	// Peek returns the top element of the stack without removing it.
	Peek() interface{}

	// IsEmpty returns true if the stack is empty, otherwise returns false.
	IsEmpty() bool

	// Size returns the number of elements in the stack.
	Size() int

	// Clear removes all elements from the stack.
	Clear()
}

type StackWithList struct {
	elements []interface{}
}

func NewStackWithList() *StackWithList {
	return &StackWithList{}
}

func (s *StackWithList) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *StackWithList) Clear() {
	// we set the elements to nil so that the garbage collector can free the memory
	s.elements = nil
}

func (s *StackWithList) Size() int {
	return len(s.elements)
}

func (s *StackWithList) Push(element interface{}) {

	// we initialize the elements slice if it is nil
	if s.elements == nil {
		s.elements = make([]interface{}, 0)
	}

	s.elements = append(s.elements, element)
}

func (s *StackWithList) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	// we store the last element in a variable
	element := s.elements[len(s.elements)-1]

	// we remove the last element from the slice by slicing the slice (pun intended)
	s.elements = s.elements[:len(s.elements)-1]

	return element
}

func (s *StackWithList) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.elements[len(s.elements)-1]
}
