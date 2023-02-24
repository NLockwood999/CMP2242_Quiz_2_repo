//demonstration of wait groups
// Filename: main.go
package main

import (
	"fmt"
	"sync"
)

func alpha(WG *sync.WaitGroup) {
	// decrement the done by 1
	defer WG.Done()
	fmt.Println("this is alpha func")
}
func beta(WG *sync.WaitGroup) {
	// decrement the done by 1
	defer WG.Done()
	fmt.Println("beta ready for func")
}
func gamma(WG *sync.WaitGroup) {
	// decrement the done by 1
	defer WG.Done()
	fmt.Println("gamma executing task")
}

func main() {
	var WG sync.WaitGroup
	//we create wg,add(3) since there are 3 indpendent functions in main
	WG.Add(3)

	go alpha(&WG)
	go beta(&WG)
	go gamma(&WG)

	//application of waitgroups
	WG.Wait()

}
