package collections

import (
	"reflect"
	"testing"
)

func TestIntListMap(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := NewIntList(tt.in).
				Map(func(n int) int { return n * 2 }).
				Get()

			if !reflect.DeepEqual(result, tt.out) {
				t.Errorf("got %+v, want %+v", result, tt.out)
			}
		})
	}
}
