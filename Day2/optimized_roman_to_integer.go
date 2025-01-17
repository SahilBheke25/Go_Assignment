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

// To validate roman string.
func validate(romanNumber string) bool {
	validRomanPattern := `^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`
	pattern := regexp.MustCompile(validRomanPattern)
	if !pattern.MatchString(romanNumber) {
		fmt.Println("Invalid Roman Number")
		return false
	}
	return true
}

// To convert roman string to integer.
func convertRomanToInteger(romanNumber string) {

	var romanValues = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	len := len(romanNumber)
	var ans int
	for i := 0; i < len; i++ {

		current := string(romanNumber[i])

		value := romanValues[current]
		// If expression is in the form IV, IX, XL, XC, CD, CM. Such that if current number is less than next nummber it will subtract current number form the result.
		if i+1 < len && value < romanValues[string(romanNumber[i+1])] {
			ans -= value
		} else {
			ans += value
		}
	}

	fmt.Println("The Roman to Interger conversion is: ", ans)

}

func main() {

	fmt.Println("Converting Roman Number to Integer")

	var isRunning int = 1
	for isRunning >= 1 {

		var romanNumber string
		fmt.Print("Enter Roman Number: ")
		fmt.Scan(&romanNumber)

		romanNumber = strings.ToUpper(romanNumber)
		romanNumber = strings.Trim(romanNumber, " ")

		if !validate(romanNumber) {
			return
		}

		convertRomanToInteger(romanNumber)

		fmt.Println("If you want to CONTINUE ENTER 1, Else to END ENTER 0")
		fmt.Scan(&isRunning)
	}
}
