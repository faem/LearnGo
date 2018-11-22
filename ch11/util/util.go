package util

//Finds the minimum number from an array
func Min(V []int) int{
	temp := 9999

	for _,v := range V{
		if temp > v{
			temp = v
		}
	}
	return temp
}

//The below max funtion is invisible to others
func max(V []int) int{
	temp := 0

	for _,v := range V{
		if temp < v{
			temp = v
		}
	}
	return temp
}