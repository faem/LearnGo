package main

import "fmt"

func main() {
	avgVariadic(1, 2, 3, 4, 5)
}

//variadic functions
func avgVariadic(values ...int) { //... is used to indicate that values can take many value.
	// Last parameter of a function can be used like this
	total := 0
	for _, value := range values {
		total = total + value
	}
	fmt.Println(total / len(values))
}
