package main

import (
	"fmt"
)

// loop over range from 0 to 10, and at each iteration try to receive from channel and print the output, and if you can't, assign 
// the current counter value to the channel. 
// returns 0, 2, 4,.... as a consequence. 
func main() {
	c := make(chan int, 1)
	for i := 0; i < 10; i++{
		select {
			// if you can receive from c, do it and assign to i, and use it
			case j:= <- c:
				fmt.Println(j)
			// if you can't receive from c, then do something...
			default:
				c <- i
		}
	}
}