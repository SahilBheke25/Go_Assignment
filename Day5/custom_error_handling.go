package main

import "fmt"

/*
2. In the given code, the accessSlice function accepts a slice and index.
If the value is present in the slice at that index, the program should print the following.

"item <index>, value <value present at the index>"

But if the index does not hold any value,
it will lead to an index out of range panic in our program.


So in order to safeguard our program from panicking, add a condition to handle the condition if

index > lengthOfSlice - 1

and return an error from the accessSlice function if the above condition is met.
The error message should be the following :

"length of the slice should be more than index"

Complete the given program to return an error from the accessSlice function and handle the returned error inside the main function to print the message.

Example Test Case 1 :
Input: 3
Output:
item 3, value 6
*/

func accessSlice(slice []int, index int) error {
	if index < 0 || index >= len(slice) {
		return fmt.Errorf("length of the slice should be more than index")
	}
	fmt.Printf("item %v, value %v\n", index, slice[index])
	return nil
}
func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var index int
	fmt.Print("Enter index: ")
	_, err1 := fmt.Scanf("%v", &index)
	if err1 != nil {
		fmt.Println("Enter valid integer input")
		return
	}
	err2 := accessSlice(slice, index)
	if err2 != nil {
		fmt.Println("Error Occured:", err2)
	}
}
