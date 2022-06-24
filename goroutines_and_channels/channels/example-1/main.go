// An example code showing a simple use of a channel
// A buffered channel will be created and items will be written and read from it without causing a deadlock

package main

import "fmt"

func main() {
	var myChan chan int = make(chan int, 1)
	myChan <- 1
	fmt.Println("Wrote to the channel")
	fmt.Printf("Read %d from the channel\n", <-myChan)
}
