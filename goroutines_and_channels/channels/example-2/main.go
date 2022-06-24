// An example of using goroutines along with a channel
// A goroutine producer will produce the N'th number of the fibonacci series and writes it to the channel
// The consumer will wait for M seconds for the result and if result is not ready, it will stop the program

package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan int)

	go func(limit int) {
		a, b := 0, 1
		for i := 0; i < limit-1; i++ {
			a, b = b, a+b
		}
		result <- a
	}(1000)

	select {
	case res := <-result:
		fmt.Println(res)
	case <-time.After(1 / 10 * time.Second):
		fmt.Println("Exitting...")
	}

}
