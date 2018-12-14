package main

import (
	"fmt"

	"github.com/owainlewis/hof/pkg/collections"
)

func main() {
	data := []int{1, 2, 3, 4, 5}

	dl := collections.
		NewIntList(data).
		Filter(func(n int) bool { return n > 3 }).
		Map(func(n int) int { return n + 1 }).
		Get()

	fmt.Printf("Result %v\n", dl)

}
