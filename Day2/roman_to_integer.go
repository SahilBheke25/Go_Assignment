package main

import (
	"fmt"
)

func main() {

	var inputString string

	fmt.Print("Enter Roman Number: ")
	fmt.Scan(&inputString)

	var result int

	for i := 0; i < len(inputString); i++ {
		if inputString[i] == 'I' {
			if i < len(inputString)-1 {
				if inputString[i+1] == 'V' {
					result += 4
				} else if inputString[i+1] == 'X' {
					result += 9
				} else {
					result += 1
					continue
				}
				i++
			} else {
				result += 1
			}
		} else if inputString[i] == 'V' {
			result += 5
		} else if inputString[i] == 'X' {
			if i < len(inputString)-1 {
				if inputString[i+1] == 'L' {
					result += 40
				} else if inputString[i+1] == 'C' {
					result += 90
				} else {
					result += 10
					continue
				}
				i++
			} else {
				result += 10
			}
		} else if inputString[i] == 'L' {
			result += 50
		} else if inputString[i] == 'C' {
			if i < len(inputString)-1 {
				if inputString[i+1] == 'D' {
					result += 400
				} else if inputString[i+1] == 'M' {
					result += 900
				} else {
					result += 100
					continue
				}
				i++
			} else {
				result += 100
			}
		} else if inputString[i] == 'D' {
			result += 500
		} else if inputString[i] == 'M' {
			result += 1000
		}
	}

	fmt.Println("The Roman to Interger conversion is: ", result)
}
