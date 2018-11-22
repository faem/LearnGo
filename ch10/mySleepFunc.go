package main

import (
	"log"
	"time"
)

func main() {
	for {
		sleep(2)
	}

}

func sleep(t int) {
	select {
	case <-time.After(time.Duration(t) * time.Second):
		log.Println("sleeping")
	}

}
