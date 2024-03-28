package main

import "fmt"
import "time"


// child go routine
func doWork(done <- chan bool) { // function which gets a channel passed through containing boolean data
	for {
		select { // for select lets us check if the parent wants to cancel the done channel
		case something := <- done:
			fmt.Println("From done channel: ", something)
			return
		default:
			fmt.Println("... doing work")
		}
	}
}

// parent go routine
func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 2)

	close(done) // close the done channel, stops the child go routine

	time.Sleep(time.Second * 3)

	// result: ... doing work for 3 seconds till the parent go routine cancels the child go routine
}