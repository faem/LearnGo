package main

import "fmt"

func main() {
	x := []float64{1, 2, 3, 4, 5}

	avgClosure(x)

	//closure inside main function
	i := 5
	dec := func() int{
		i--
		return i
	}
	fmt.Println("Closure inside main function:")
	fmt.Println(dec())
	fmt.Println(dec())

	//IMPORTANT
	nextOdd := oddGenerator() // statements before return statement will only execute only once here

	fmt.Println("nextOdd")
	fmt.Println(nextOdd()) //only return statement will execute here
	fmt.Println(nextOdd()) //here too
}

//closure : When declare a function inside another function

func avgClosure(x []float64) {
	total := func(x []float64) float64 {
		temp := 0.0
		for _, value := range x {
			temp = temp + value
		}
		return temp
	}

	fmt.Println(total(x) / float64(len(x)))
}


func oddGenerator() func() int{
	fmt.Println("inside oddGenerator")
	i := int(1)
	fmt.Println(i)
	return func() int {
		i+=2
		return i
	}
}


