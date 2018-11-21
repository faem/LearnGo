package main

import "fmt"

type Person struct{
	Name string
}

func (p *Person) getName(){
	fmt.Println(p.Name)
}

type Teacher struct{
	Person
	schoolName string
}

func main()  {
	var t Teacher
	t.Name = "John"
	fmt.Println(t.Person.Name)
	t.Person.getName()
	//v is equivalent ^
	t.getName()
}