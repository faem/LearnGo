package main

import "fmt"

func main(){
	//defer works as stack, it pushes the result to a stack.
	// returns the result from stack when all other functions returned
	defer f1() //this will run at the end of all function call
	defer f2()
	f3()
	x := "why"
	fmt.Println(x)

	if x=="wh"{
		fmt.Println("if is working")
	}else{
		fmt.Println("else is working")
	}
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
}

func f1(){
	fmt.Println("f1")
}

func f2(){
	fmt.Println("f2")
}

func f3(){
	fmt.Println("f3")
}