package main

import "testing"

func TestMin(t *testing.T){
	x := min([]int{1,2,3,4,5,6})

	if x<1 || x>5{
		t.Error("Out of Range")
	}
}

func TestMax(t *testing.T){
	x := max([]int{0,1,2,3,4,5,6})

	if x<1 || x>5{
		t.Error("Out of Range")
	}
}