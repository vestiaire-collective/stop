package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/vestiaire-collective/stop"
)

func main() {
	stop.Listen()

	w := sync.WaitGroup{}

	w.Add(5)

	for i := range make([]int, 5) {
		go func(i int) {
			defer w.Done()

			for !stop.Stopped() {
				fmt.Printf("sleeping %+v...\n", i)
				time.Sleep(1 * time.Second)
			}

			fmt.Printf("stopping %+v...\n", i)
		}(i)
	}

	w.Wait()

	println("stopped")
}
