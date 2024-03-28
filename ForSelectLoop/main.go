package main

import "fmt"

func main() {
	// a buffer acts as a queue where you can send the buffer length of data and it is treated as a queue
	bufferedChannel := make(chan string, 3) // opened a channel accepting data of type string with a buffer length of 3

	charsList := []string{"a", "b", "c"} // create a list array with data

	for _, character := range charsList { // loop over the array of character
		select {
		case bufferedChannel <- character: // send the character from the list to the channel
		}
	}

	close(bufferedChannel) // close the channel, you can still loop over the channel

	for result := range bufferedChannel { // loop over the closed channel
		fmt.Println(result)
	}

	// result: a \n b \n c
}