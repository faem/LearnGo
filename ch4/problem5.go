package main

import "fmt"

func main(){
	var fahrenheit, celcius float64
	fmt.Scanf("%f", &fahrenheit)

	celcius = (fahrenheit - 32) * 5/9

	fmt.Println(fahrenheit,"Fahrenheit =",celcius,"Celcius")
	fmt.Printf("%.2f Fahrenheit = %.2f Celcius\n",fahrenheit, celcius)
	//fmt.Printf("%.2f Fahrenheit",fahrenheit," = %.2f Celcius", celcius) // DO NOT


}
