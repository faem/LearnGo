package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "AppsCode"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "Limited"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case m3 := <-time.After(time.Second * 2): //will execute when no channel is sending
				log.Println(m3)
			case m1 := <-c1:
				log.Println(m1)
			case m2 := <-c2:
				log.Println(m2)

			}
		}
	}()

	var x string
	fmt.Scanln(&x)
}
