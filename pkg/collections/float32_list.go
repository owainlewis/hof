package collections

type Float32List struct {
	items []float32
}

// NewStringList ..
func NewFloat32List(items []float32) *Float32List {
	return &Float32List{
		items: items,
	}
}

// Get returns the underlying type value for this collection.
func (s *Float32List) Get() []float32 {
	return s.items
}

func (s *Float32List) Map(f func(float32) float32) *Float32List {
	result := []float32{}
	for _, item := range s.items {
		result = append(result, f(item))
	}

	s.items = result
	return s
}

// Filter will remove items from a sequence that match a given
// predicate function.
func (s *Float32List) Filter(f func(float32) bool) *Float32List {
	result := []float32{}
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
func (s *Float32List) Select(f func(float32) bool) *Float32List {
	result := []float32{}
	for _, item := range s.items {
		if f(item) {
			result = append(result, item)
		}
	}

	s.items = result
	return s
}

func (s *Float32List) Any(f func(float32) bool) bool {
    for _, v := range s.items {
        if f(v) {
            return true
        }
    }
    return false
}

// All returns true if all of the items in the slice satisfy the predicate f.
func (s *Float32List) All(f func(float32) bool) bool {
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
