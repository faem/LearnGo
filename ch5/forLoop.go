package main

import "fmt"

func main() {
	//type 1 (like while loop)
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	//type 2
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}
