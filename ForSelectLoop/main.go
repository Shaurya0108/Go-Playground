package main

import "fmt"

func main() {
	// a buffer acts as a queue where you can send the buffer length of data and it is treated as a queue
	bufferedChannel := make(chan string, 3) // opened a channel accepting data of type string with a buffer length of 3
}