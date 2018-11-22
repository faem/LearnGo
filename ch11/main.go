package main

import (
	"fmt"
)
import u "LearnGO/ch11/util" //alias

func main(){
	x := u.Min([]int{1,2,3,1,56,0,3})
	//y := util.max([]int{1,2,3,1,56,0,3})
	//^this will not work as max function is declared using lowercase 'm'.
	// So it is invisible to other function
	fmt.Println(x)
}
