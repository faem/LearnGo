package main

import "fmt"

type Vehicle interface{
	distanceTraveled() float64
}

type Bike struct {
	distance float64
	speed float64
}

func (b Bike) distanceTraveled() float64{
	return b.distance
}

type Car struct {
	distance float64
	speed float64
}

func (c Car) distanceTraveled() float64{
	return c.distance
}


func totalDistance(vehicles ... Vehicle) float64{
	temp := 0.0
	for _,vehicle := range vehicles{
		temp = temp + vehicle.distanceTraveled()
	}
	return temp
}



//interface as fields

type MultiVehicle struct {
	vehicle []Vehicle
}


func main()  {
	var c Car
	var b Bike
	//c := new(Car)
	//b := new(Bike)
	c.distance = 14
	b.distance = 10

	fmt.Println(totalDistance(c,b)) //pass pointer

	mv := new(MultiVehicle)
	mv.vehicle = append(mv.vehicle,&c,&b)
	for _,v := range mv.vehicle{
		fmt.Println(v)
	}
}