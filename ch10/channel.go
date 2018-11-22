package main

import (
"fmt"
"time"
)
func sendAppsCode(c1 chan<- string, c2 <-chan string){ //c1 is sending, c2 is receiving
	for{
		c1 <- "AppsCode" //sending
		m := <- c2
		fmt.Println(m)
		time.Sleep(time.Second*1)
	}
}

func sendLimited(c1 chan<- string,c2 chan<- string){ //c1 & c2 are sending
	for{
		c1 <- "Limited" //sending
		c2 <- "Bangladesh"
	}

}

func printCompanyName(c chan string){ //c is bidirectional, no <- sign
	for{
		m := <- c //receiving
		fmt.Println(m)
		time.Sleep(time.Second*1)
	}
}
func main(){
	c := make(chan string) //without 'making' channel will not work
	c2 := make(chan string)
	go sendAppsCode(c,c2)
	go sendLimited(c,c2)
	go printCompanyName(c)
	var x string
	fmt.Scanln(&x) //Scanln is given so that we can see the behaviour of goroutines
}
