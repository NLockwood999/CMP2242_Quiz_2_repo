package main

import (
	"fmt"
	"time"
)

//this funtion handles the printing
func printLetters() {
	for a := 1; a <= 4; a++ {
		time.Sleep(3 * time.Second)
		fmt.Println("Letter : ", a)
	}
}

//main itself is a go routine
func main() {
	//a function becomes go routine by adding go at the beginning
	go printLetters()
	fmt.Println("executing this line of code while printing letters")
	time.Sleep(25 * time.Second)
	fmt.Println("Your letters have been successfully printed, postman :)")
}

// a go routine can be described as an independent entity
