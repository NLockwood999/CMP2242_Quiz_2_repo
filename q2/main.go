//demonstrate go routines
package main

import (
	"fmt"
)

func messenger(msg chan string) {
	//write data to the channel block
	msg <- "messages to deliver today"
	close(msg)
}

func letters(num chan int) {
	//write dats to the channel blocks
	num <- 13
	close(num)
}

func main() {
	//create a channel value that will pass int data
	quantity := make(chan int)
	//create a channel msg that will pass string data
	msg := make(chan string)

	//create a go routine called messenger that accepts msg as a parameter
	go messenger(msg)
	//create a go routine called messenger that accepts quantity as a parameter
	go letters(quantity)
	//read data from channel to a variable
	result1 := <-msg
	result2 := <-quantity
	//prints the data
	fmt.Println("I have", result2, result1)
}
