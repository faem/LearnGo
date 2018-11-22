package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	ch := make(chan string, 1) //number of channels, if we give
	go func() {
		for {
			ch <- "AppsCode"
			ch <- "AppsCode2"
			ch <- "AppsCode3"
		}
	}()

	go func() {
		for {
			log.Println(<-ch)
			log.Println(<-ch)
			log.Println(<-ch)
			time.Sleep(time.Second * 1)
		}
	}()

	var x string
	fmt.Scanln(&x)
}
