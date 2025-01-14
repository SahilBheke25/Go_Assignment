package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Assignment for day 2:
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
// Roman numerals are usually written from largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer

func main() {

	fmt.Println("Converting Roman Number to Integer")

	var flag int = 1
	for flag >= 1 {

		var inputString string

		fmt.Print("Enter Roman Number: ")
		fmt.Scan(&inputString)

		inputString = strings.ToUpper(inputString)
		inputString = strings.Trim(inputString, " ")

		validRomanPattern := `^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`

		//Regex Compilation
		pattern := regexp.MustCompile(validRomanPattern)

		//Validation
		if !pattern.MatchString(inputString) {
			fmt.Println("Invalid Roman Number")
			return
		}

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

		fmt.Println("If you want to continue Enter any number != 0 Else to END ENTER 0")
		fmt.Scan(&flag)
		if flag >= 1 {
			continue
		} else {
			break
		}
	}

}
