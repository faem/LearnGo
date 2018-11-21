package main

import "fmt"

func main() {
	i := -2

	if i == 0 {
		fmt.Println("Zero")
	} else if i <= 0 {
		fmt.Println("Less")
	} else {
		fmt.Println("Greater")
	}
}
