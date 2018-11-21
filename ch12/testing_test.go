package main

import "testing"

func TestMn(t *testing.T){
	x := mn([]int{0,1,2,3,4,5,6})

	if x<1 || x>5{
		t.Error("Out of Range")
	}
}
