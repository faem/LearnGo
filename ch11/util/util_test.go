package util

import "testing"

func TestMin(t *testing.T) {
	v := Min([]int{0,1,2,4,24,2,-1})

	if v<0{
		t.Error("Expected positive number!")
	}
}