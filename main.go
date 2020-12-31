package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	testCond()
}

func testCond() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("adding to queue")
		queue = append(queue, struct{}{})

		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}

}
