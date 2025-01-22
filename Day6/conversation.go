package main

import (
	"fmt"
	"sync"
)

/*
1. Given a string containing conversation between alice and bob. In the string, if it reaches $, it means end of alice message and if it reaches #, it means end of bob's message. and if it reaches ^,
it means end of conversation ignore string after that.

Note: given string doesn't contain any spaces. If message contains two continuous conversations from one person it should be printed one after another as given in the example.

write a program to separate out messages from alice and bob. Write messages from alice and bob on seperate channels, whenever a message from alice/bob, print it in front of their name as shown in the example below:

Note: there is a single space before and after colon(:) and no space before and after comma.

e.g: "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"
output: alice : helloBob,bob : helloalice,bob : howareyou?,alice : Iamgood.howareyou?
*/

func main() {
	var message string = "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"
	aliceChannel := make(chan string)
	bobChannel := make(chan string)
	var wg sync.WaitGroup

	// Index for start of new message.
	var prevIndex int = 0

	for i := 0; i < len(message)-1; i++ {

		if message[i] == '$' {

			// If alice talks
			messageSlice := message[prevIndex:i]
			wg.Add(1)

			go func(i int) {
				if i == len(message)-2 {
					// End of conversation
					fmt.Print("alice : ", <-aliceChannel)
				} else {
					fmt.Print("alice : ", <-aliceChannel, ",")
				}
				wg.Done()
			}(i)

			aliceChannel <- messageSlice
			prevIndex = i + 1

		} else if message[i] == '#' {

			// If bob talks
			messageSlice := message[prevIndex:i]
			wg.Add(1)

			go func(i int) {
				if i == len(message)-2 {
					// End of conversation
					fmt.Print("bob : ", <-bobChannel)
				} else {
					fmt.Print("bob : ", <-bobChannel, ",")
				}
				wg.Done()
			}(i)

			bobChannel <- messageSlice
			prevIndex = i + 1
		}
		wg.Wait()
	}
}
