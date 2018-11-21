package main

import "fmt"

func main() {
	fmt.Println(fibonacci(12))
}

//recursion
func fibonacci(x int) int{
	if x==0{
		return 0
	}
	if x==1{
		return 1
	}

	return fibonacci(x-2)+fibonacci(x-1)
}
