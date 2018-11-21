package main

import "fmt"

func main() {

	//Declaration type 1
	var x [5]int
	x[4] = 10
	fmt.Println(x)

	//Declaration type 2
	y := [5]float32{1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(y)
}
