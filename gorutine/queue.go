package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 5)

	go func() {
		time.Sleep(3 * time.Second)
		c <- "foo"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c <- "bar"
	}()

	var cmds []string
	cmds = append(cmds, <-c)

wait_some:
	for {
		select {
		case cmd := <-c:
			cmds = append(cmds, cmd)
		case <-time.After(1 * time.Second):
			break wait_some
		}
	}
	for _, cmd := range cmds {
		fmt.Println(cmd)

	}
}
