package main

import "fmt"

/*
1. In the given code, the accessSlice function accepts a slice and index.
If the value is present in the slice at that index, the program should print the following.

"item <index>, value <value at that index in slice>"

But if the index does not hold any value,
it will lead to an index out of range panic in our program.

Complete the given code to recover from panic inside the accessSlice function, when index out of range panic occurs.
Also, Print the following after handling panic

"internal error: <recovered panic value>"

Example Test Case 1 :
Input: 3
Output:
item 3, value 6
Testing Panic and Recover
*/

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("internal error: %v", r)
		}
	}()
	fmt.Printf("item %v, value %v\n", index, slice[index])
	fmt.Println("Testing Panic and Recover")
}
func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var index int
	fmt.Print("Enter index: ")
	_, err := fmt.Scanf("%v", &index)
	if err != nil {
		fmt.Println("Enter valid integer input")
		return
	}
	accessSlice(slice, index)
}
