package main

import "fmt"

func main() {
	//Declaration type 1
	var cgpa map[string]float64 	//creates nil map, must make before assigning values to it.
									//Assigning without make will cause runtime error
	cgpa = make(map[string]float64)
	cgpa["fahim"] = 3.42
	cgpa["masud"] = 3.76

	fmt.Println(cgpa["fahim"])

	//Declaration type 2
	age := map[string]int{
		"fahim": 25,
		"masud": 25,
	}
	fmt.Println(age)

	//delete element
	delete(cgpa, "masud")
	fmt.Println(cgpa)

	//better of checking if a key exists or not
	if value, flag := cgpa["fahim"]; flag {
		fmt.Println(value, flag)
	}

	if value, flag := cgpa["masud"]; flag {
		fmt.Println(value, flag)
	}

	//nested map
	nestedMap := map[string]map[string]float64{
		"fahim": map[string]float64{
			"age":  25,
			"cgpa": 3.42,
		},
		"masud": map[string]float64{
			"age":  25,
			"cgpa": 3.76,
		},
	}
	fmt.Println(nestedMap)

	if name, flag := nestedMap["fahim"]; flag {
		fmt.Println(name["age"])
	}
}
