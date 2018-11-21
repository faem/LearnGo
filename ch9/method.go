package main

import "fmt"

func main() {
	r := Rect{3,4}
	r.area()
}

func (r *Rect) area(){
	fmt.Println(r.y*r.z)
}
