package collections

type Float64List struct {
	items []float64
}

// NewStringList ..
func NewFloat64List(items []float64) *Float64List {
	return &Float64List{
		items: items,
	}
}

// Get returns the underlying type value for this collection.
func (s *Float64List) Get() []float64 {
	return s.items
}

func (s *Float64List) Map(f func(float64) float64) *Float64List {
	result := []float64{}
	for _, item := range s.items {
		result = append(result, f(item))
	}

	s.items = result
	return s
}

// Filter will remove items from a sequence that match a given
// predicate function.
func (s *Float64List) Filter(f func(float64) bool) *Float64List {
	result := []float64{}
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
func (s *Float64List) Select(f func(float64) bool) *Float64List {
	result := []float64{}
	for _, item := range s.items {
		if f(item) {
			result = append(result, item)
		}
	}

	s.items = result
	return s
}

func (s *Float64List) Any(f func(float64) bool) bool {
    for _, v := range s.items {
        if f(v) {
            return true
        }
    }
    return false
}

// All returns true if all of the items in the slice satisfy the predicate f.
func (s *Float64List) All(f func(float64) bool) bool {
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
