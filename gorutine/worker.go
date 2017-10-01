package main

import (
	"fmt"
	"sync"
	"time"
)

// fetchURL is method to fetch website data and send to channel
func fetchURL(wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
	for {
		url, ok := <-c
		if !ok {
			break
		}
		fmt.Println("download: ", url)
		time.Sleep(3 * time.Second)
	}
	fmt.Println("finish")
}

func main() {
	var wg sync.WaitGroup
	c := make(chan string, 5)

	// Generateo 3 workers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go fetchURL(&wg, c)
	}

	// Execute 3 workers in the same time
	c <- "http://www.example1.com"
	c <- "http://www.example2.com"
	c <- "http://www.example3.com"
	c <- "http://www.example4.com"
	c <- "http://www.example5.com"
	close(c)

	wg.Wait()
}
