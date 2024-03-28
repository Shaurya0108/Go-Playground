package main

import "fmt"

func main() {
	mainChannel := make(chan string) // channel opened

	go func() {
		mainChannel <- "Hello, World!"
	}()

	res := <- mainChannel // channel closed

	fmt.Println(res)
}