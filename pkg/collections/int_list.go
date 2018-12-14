package collections

type IntList struct {
	items []int
}

// NewStringList ..
func NewIntList(items []int) *IntList {
	return &IntList{
		items: items,
	}
}

// Get returns the underlying type value for this collection.
func (s *IntList) Get() []int {
	return s.items
}

func (s *IntList) Map(f func(int) int) *IntList {
	result := []int{}
	for _, item := range s.items {
		result = append(result, f(item))
	}

	s.items = result
	return s
}

// Filter will remove items from a sequence that match a given
// predicate function.
func (s *IntList) Filter(f func(int) bool) *IntList {
	result := []int{}
	for _, item := range s.items {
		if !f(item) {
			result = append(result, item)
		}
	}

	s.items = result
	return s
}

// Select will return items from a sequence that match a given
// predicate function.
func (s *IntList) Select(f func(int) bool) *IntList {
	result := []int{}
	for _, item := range s.items {
		if f(item) {
			result = append(result, item)
		}
	}

	s.items = result
	return s
}

func (s *IntList) Any(f func(int) bool) bool {
    for _, v := range s.items {
        if f(v) {
            return true
        }
    }
    return false
}

// All returns true if all of the items in the slice satisfy the predicate f.
func (s *IntList) All(f func(int) bool) bool {
    for _, v := range s.items {
        if !f(v) {
            return false
        }
    }
    return true
}

// TODO

// Take
// TakeWhile
// Reverse
