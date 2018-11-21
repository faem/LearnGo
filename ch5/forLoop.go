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

	//type 3 (array or other data types)
	var x [5]int

	for _,value:= range x{
		fmt.Println(value)
	}

	for i,value:= range x{
		fmt.Println("x[",i,"] = ",value)
	}
}
