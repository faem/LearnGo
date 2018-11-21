//swap two integers using address
package main

import "fmt"

func main()  {
	x := 1
	y := 2

	swap(&x,&y)
	fmt.Println(x,y)
}

func swap(x *int, y *int)  {
	temp := *x
	*x = *y
	*y = temp
}