package main

import "fmt"

func main(){
	//Variable Declaration type 1
	var x, y = "Hello ", "World 1"
	x = x + y
	fmt.Println(x)

	//Variable Declaration type 2
	var z string
	z = "Hello World 2"
	fmt.Println(z)

	//Variable Declaration type 3
	a := "Hello World 3"
	fmt.Println(a)

	//Variable Declaration type 4
	var (
		v1 = 5
		v2 = 10
		v3 = 15.0
	)

	fmt.Println(float64(v1)+float64(v2)+v3) //can not add float with int so type converted

	//Constant
	const c string = "Hello World Constant"
	fmt.Println(c)



}