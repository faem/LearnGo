package main

func main()  {
	Min([]int{1,2,3,4,0})
}

func Min(V []int) int{
	temp := 9999

	for _,v := range V{
		if temp > v{
			temp = v
		}
	}
	return temp
}
