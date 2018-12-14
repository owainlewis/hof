# HOF

HOF is a set of higher order function utilities for working with core Golang types.

Golang looping constructs can often lead to verbose, hard to read code. This library makes
basic data transformation tasks easier to reason about whilst still maintaining good performance.

This library uses code generation rather than reflection so there should be no performance hit when
using these utilities.

## Example

```go
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
```
