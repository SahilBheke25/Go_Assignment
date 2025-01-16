package main

/*
1. The given program accepts an integer value between 1 to 4 from the user.
There is an option associated with each value, which is basically objects of different data types with some associated value.

Write a method named "AcceptAnything" which takes any type of data as input.

Based on the option chosen by the user, we will send different types of objects to this function and it should print the following based on its respective data type of value sent to the function "AcceptAnything".

integer :
"This is a value of type Integer, <value>"

string :
"This is a value of type String, <value>"

boolean :
"This is a value of type Boolean, <value>"

Custom data type Hello :
"This is a value of type Hello, <value>"
*/

import "fmt"

type AcceptAnything interface{}

type Hello struct {
	message string
}

func main() {
	var option int
	fmt.Println(`Enter 1-4 for following operation:
	1: To get integer object
	2: To get string object
	3: To get boolean object
	4: To get custom object`)
	_, err := fmt.Scanf("%v", &option)

	if err != nil || option < 1 || option > 4 {
		fmt.Println("Invalid Input")
		return
	}

	switch option {
	case 1:
		AcceptAnything := 4
		fmt.Printf("The object type is: %T, %v", AcceptAnything, AcceptAnything)
	case 2:
		AcceptAnything := "Let it go Sparsh Arora"
		fmt.Printf("The object type is: %T, %v", AcceptAnything, AcceptAnything)
	case 3:
		AcceptAnything := true
		fmt.Printf("The object type is: %T, %v", AcceptAnything, AcceptAnything)
	case 4:
		AcceptAnything := Hello{message: "Go with flow Sparsh Arora"}
		fmt.Printf("The object type is: %T, %v", AcceptAnything, AcceptAnything.message)
	}
}
