package main

import (
	"fmt"
	"github.com/chaseisabelle/stop"
	"time"
)

func main() {
	stop.Listen()

	c := make(chan struct{})

	for i := range make([]int, 5) {
		go func(i int) {
			for {
				if stop.Interrupted() {
					fmt.Printf("stopping %+v...\n", i)

					break
				}

				fmt.Printf("sleeping %+v...\n", i)

				time.Sleep(1 * time.Second)
			}

			c <- struct{}{}
		}(i)
	}

	i := 0

	for range c {
		i++

		if i >= 5 {
			break
		}
	}

	println("stopped")
}
