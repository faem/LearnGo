package main

import "fmt"

func main() {
	x := []float64{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	avg(x)
	fmt.Println(avg2(x))
	fmt.Println(avg3(x))
	fmt.Println(maxMin(y))

	//getting multiple value from a function
	max, min := maxMin(y)
	fmt.Println(max, min)

	//function without name
	func(){
		str := "This is a function without any name"
		fmt.Println(str)
	}()

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

	fmt.Println("fibonacci")
	fmt.Println(fibonacci(12))
}

//function without return
func avg(arr []float64) {
	total := 0.0
	for _, value := range arr {
		total = total + value
	}
	fmt.Println(total / float64(len(arr)))
}

//function with return type 1

func avg2(arr []float64) (ret float64) {
	total := 0.0
	for _, value := range arr {
		total = total + value
	}
	ret = total / float64(len(arr))
	return
}

//function with return type 2

func avg3(arr []float64) float64 {
	total := 0.0
	for _, value := range arr {
		total = total + value
	}
	return total / float64(len(arr))
}

//returning multiple values

func maxMin(arr []int) (int, int) {
	max := 0
	min := 99999
	for _, value := range arr {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return max, min
}