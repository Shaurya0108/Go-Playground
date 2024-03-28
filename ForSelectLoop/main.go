package main

import "fmt"

func main() {
	mainChannel := make(chan string) // channel opened

	go func() { // go routine
		mainChannel <- "Hello, World!" // channel recieves data: "Hello, World!"
	}()

	res := <- mainChannel // channel closed

	fmt.Println(res)
}