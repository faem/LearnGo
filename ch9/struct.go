package main

import "fmt"

type Rect struct {
	y int
	z int
}

func main()  {
	//Initialization 1
	var c1 Rect
	fmt.Println(c1.y)

	//Initialization 2
	c2 := new(Rect)
	fmt.Println(c2.y)

	//Initialization 3
	c3 := Rect{y:5}
	fmt.Println(c3.y)

	//Initialization 4
	c4 := Rect{3,2}
	fmt.Println(c4.y)
	rectArea(c4)
}

func rectArea(c Rect){
	fmt.Println("Area =",c.y  * c.z)
}