package main

import (
	"fmt"
	"time"
)

func main() {
	// empty struct size is 0
	c := make(chan struct{}, 2) // if buffer number less than 2, len(c) always return 0
	go func() {
		fmt.Println("start")
		time.Sleep(3 * time.Second)
		fmt.Println("end")
		c <- struct{}{}
	}()
	for {
		if len(c) > 0 { // check queue number
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("something")
	}
}
