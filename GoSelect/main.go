package main

import "fmt"

func main() {
	mainChannel := make(chan string) // channel opened
	secondChannel := make(chan string) // second channel opened

	go func() { // go routine
		mainChannel <- "Hello, World!" // channel recieves data: "Hello, World!"
	}()

	go func() { // go routine 2
		secondChannel <- "Bye, World!" // channel recieves data: "Bye, World!"
	}()

	res := <- mainChannel // channel closed
	res2 := <- secondChannel // second channel closed

	fmt.Println(res)
	fmt.Println(res2)
}