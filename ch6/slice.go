package main

import "fmt"

func main(){
	//Declaration
	var x []int //x has length of 0
	x = make([]int, 5) //x has length of 5
	fmt.Println(x)

	//Declaration
	y := make([]int,5,10) //slice y has length 5 from an array of capacity 10
	fmt.Println(y)

	//array to slice
	arr := [5]int{1,2,41,3,2}
	z := arr[0:5]
	fmt.Println(z[1:2])

	//append
	slicexyz := append(x,1,2)
	fmt.Println(slicexyz)

	//copy
	copySlice := make([]int, 6)
	copy(copySlice,slicexyz)
	fmt.Println(copySlice)

	//extend?
	//slicexyz = append(slicexyz, make([]int,5))


}
