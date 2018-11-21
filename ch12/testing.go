package main


func mn(V []int) int{
	temp := 9999

	for _,v := range V{
		if temp > v{
			temp = v
		}
	}
	return temp
}
