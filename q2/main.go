//demonstrate go routines
package main

import (
	"fmt"
)

func sender(value chan string) {
	//write to the channel blocks
	value <- "hello from sender"
	close(value)
	fmt.Println("Goodbye")
}

func main() {
	//create a channel
	value := make(chan string)
	//go routine
	go sender(value)
	result := <-value
	fmt.Println(result)
}

/*
//buffered
func main() {
	//create a channel
	value := make(chan int, 3)
	value <- 11
	value <- 22
	value <- 33
	//read value from channel
	for i := 0; i < 3; i++ {
		fmt.Println(<-value)
	}
}
*/

/*
//unbuffered
func main() {
	//create a channel
	value := make(chan int)
	value <- 11
	fmt.Println(<-value)

}*/
