package main

import "fmt"

func main(){
	var xPtr *int //This doesn't assign any memory, we can't set any value for *xPtr
	x := 7
	fmt.Println(x)
	xPtr = &x //xptr now refer to the address of x
	*xPtr = 5 //as xPtr has the address of x, so x will also change
	fmt.Println(xPtr,&x,*xPtr,x) //xPtr is address of x, *xPtr is the value of x.
	// changing value of x will change *xPtr and vice versa.
	// But &xPtr is the address of the pointer itself.


	yPtr := new(int) //This assigns memory for yPtr
	*yPtr = 5 //so we can set any value for *xPtr
	fmt.Println(*yPtr)
	zero(yPtr)
	fmt.Println(*yPtr)
}

func zero(xPtr *int){
	*xPtr = 0
}
