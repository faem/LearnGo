package main

import "fmt"

func main()  {
	fmt.Println(min([]int{1,2,3,4,0}))
	fmt.Println(max([]int{1,2,3,4,0}))
}

func min(V []int) int{
	temp := 9999

	for _,v := range V{
		if temp > v{
			temp = v
		}
	}
	return temp
}

func max(V []int) int{
	temp := 0

	for _,v := range V{
		if temp < v{
			temp = v
		}
	}
	return temp
}