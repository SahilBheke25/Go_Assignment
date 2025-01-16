package main

import "fmt"

/* 3. Return the slices
Complete the program to return 3 slices of a given array, based on the following conditions.
Given Array : ["qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"]
Hint: Hard-code the given array into your program.

Input: Two space-separated integers representing index1 and index2.
Output: Output will contain 3 slices
1. slice containing all the elements from the start of the array to Index1
2. slice containing all the elements from index1 to index2
3. slice containing all the elements from index2 to the end of the array
Conditions to Handle:
If Either of the input indexes is out of range of the given array, print "Incorrect Indexes" and exit the program

Example Test Case 1:
Input: 2 4
Output:
[qwe wer ert]
[ert rty tyu]
[tyu yui uio iop]

Example Test Case 2:
Input: 2 8
Output: Incorrect Indexes
*/

var array = []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

func indexValidation(index1 int, index2 int) bool {
	if index1 < 0 || index1 >= index2 || index2 > len(array) {
		fmt.Printf("Invalid index1. Please enter index between 0 to %v", len(array))
		return false
	}
	return true
}

func main() {

	var index1, index2 int
	fmt.Print("Enter Index1 and Index2: ")
	_, err := fmt.Scanf("%v %v", &index1, &index2)
	fmt.Print(err)

	if err != nil {
		fmt.Println(err)
		return
	}
	if !indexValidation(index1, index2) {
		return
	}

	// slicing operation.
	var slice1 = array[:index1+1]
	var slice2 = array[index1 : index2+1]
	var slice3 = array[index2:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
