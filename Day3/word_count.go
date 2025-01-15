package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//3. Word Count
//Write a Program to fulfil the following conditions.
//Input: A string containing words separated by space
//Output: A slice containing words with the highest frequency in the input string.
//Note: The order of the words in the resultant slice should be the same as the order they appear in the input string.
//Condition: Words are case-sensitive. Joe is not the same as joe.
//Given Code Template: A code template is provided that takes a string as an input and turns it into a slice of strings.

//Example Test Case 1:
//Input: My name is Joe and My Father's name is also Joe
//Output: [My name is Joe]
//Here, the words "My", "name", "is", "Joe" appeared 2 times each, which is also the highest frequency of any word.
//Hence the output contains all 4 words.

//Example Test Case 2:
//Input: Europe is far but the US is farther
//Output: [is]
//Here, the word "is" appeared twice which is also the highest frequency of any word.

// Return highest frequency words in a space separated string
func highestFrequencyWords(arr []string) []string {

	var maap = map[string]int{}
	var maxCount int = -1

	// Storing count of words.
	for _, value := range arr {
		maap[value]++
		if maap[value] > maxCount {
			maxCount = maap[value]

		}
	}

	var ans []string

	// Appending most frequent words into slice.
	for _, value := range arr {

		if maap[value] == maxCount {
			maap[value] = 0
			ans = append(ans, value)
		}
	}

	return ans
}

func main() {

	// Space seperated input string.
	var word string
	scanner := bufio.NewReader(os.Stdin)
	word, _ = scanner.ReadString('\n')

	var arr = strings.Fields(word)

	var result = highestFrequencyWords(arr)

	fmt.Println(result)
}
