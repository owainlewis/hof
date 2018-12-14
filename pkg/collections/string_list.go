package collections

type StringList struct {
	items []string
}

// NewStringList ..
func NewStringList(items []string) *StringList {
	return &StringList{
		items: items,
	}
}

// Get returns the underlying type value for this collection.
func (s *StringList) Get() []string {
	return s.items
}

func (s *StringList) Map(f func(string) string) *StringList {
	result := []string{}
	for _, item := range s.items {
		result = append(result, f(item))
	}

	s.items = result
	return s
}

// Filter will remove items from a sequence that match a given
// predicate function.
func (s *StringList) Filter(f func(string) bool) *StringList {
	result := []string{}
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
func (s *StringList) Select(f func(string) bool) *StringList {
	result := []string{}
	for _, item := range s.items {
		if f(item) {
			result = append(result, item)
		}
	}

	s.items = result
	return s
}

func (s *StringList) Any(f func(string) bool) bool {
    for _, v := range s.items {
        if f(v) {
            return true
        }
    }
    return false
}

// All returns true if all of the items in the slice satisfy the predicate f.
func (s *StringList) All(f func(string) bool) bool {
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
