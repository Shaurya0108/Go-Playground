package main

import "fmt"


func listToChannel(nums []int) <- chan int { // return a read only int channel
	out := make(chan int) // unbuffered

	go func() { // child 1 go routine
		for _, n := range nums { // for each element in the list
			out <- n // add to the channel
		}
		
		close(out) // close the out channel created in this function
	}()

	return out // the function is async, so the function does not wait for the go routine to finish
}

func sq(in <- chan int) <- chan int { // return a read only int channel
	out := make(chan int) // unbuffered

	go func() { // child 1 go routine
		for n := range in { // for each element in the in channel returned from listToChannel
			out <- n * n // add the square to the channel
		}
		
		close(out) // close the out channel created in this function
	}()

	return out // the function is async, so the function does not wait for the go routine to finish
}

// parent go routine
func main() {
	
	// input
	nums := []int{2, 3, 4, 5, 7, 2, 1}

	// stage 1
	dataChannel := listToChannel(nums) // pass the list of elements to a channel

	// stage 2
	finalChannel := sq(dataChannel) // pass the squared values of the elements to a channel

	// stage 3
	for i := range finalChannel {
		fmt.Println(i) // print the squared values from the channel
	}

	// result: the two functions wait for each other as they work in sync waiting for the reciever to recieve before sending more into the synchronous channel
	// 4
	// 9
	// 16
	// 25
	// 49
	// 4
	// 1
}