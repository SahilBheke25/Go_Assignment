package main

import (
	"fmt"
	"sync"
	"time"
)

/*
In the code snippet give in the link, concurrent goroutines execution corrupts a piece of data by
accessing it simultaneously it leads in raise condition.
https://go.dev/play/p/SmLMf8hZQzs
output when you run this : 1 is Even (may vary when you run multiple times)

Update the given code to print the correct output.
output once code is corrected: 0 is Even
*/

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0
	var m sync.Mutex

	go func(m *sync.Mutex) {

		m.Lock()
		defer m.Unlock()

		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)

		if nIsEven {
			fmt.Println(n, " is even")
			return
		}

		fmt.Println(n, "is odd")

	}(&m)

	go func(m *sync.Mutex) {

		m.Lock()
		n++
		defer m.Unlock()

	}(&m)

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
