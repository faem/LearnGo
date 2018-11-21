package main

import "fmt"

func main() {
	i := 5

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("less then 5")
	case 2:
		fmt.Println("less then 5")
	case 3:
		fmt.Println("less then 5")
	case 4:
		fmt.Println("less then 5")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("greater then 5")
	}
}
