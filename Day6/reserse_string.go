package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
2. Given a string, reverse it using one go routine.And inside go routine print reversed string and number of goroutines launched

e.g: Input: test123 output: 321tset 2
*/

func reverseString(message string, wg *sync.WaitGroup) {

	rns := []rune(message)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]

	}

	fmt.Print(string(rns), " ")
	wg.Done()
}

func main() {

	var message string
	fmt.Println("Enter string to reverse")
	_, err := fmt.Scanf("%v", &message)
	if err != nil {
		fmt.Println("Invalid input string")
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go reverseString(message, &wg)
	numGoroutines := runtime.NumGoroutine()
	wg.Wait()

	fmt.Println(numGoroutines)
}
