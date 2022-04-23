// credits : https://go.dev/tour

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	rank := 0
	n_minus_2 := 0
	n_minus_1 := 1
	return func() int {
		if rank == 0 || rank == 1{
			to_return := rank
			rank ++
			return to_return
		} else {
			fibonacci_number := n_minus_2 + n_minus_1
			n_minus_2 = n_minus_1
			n_minus_1 = fibonacci_number
			return fibonacci_number
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}