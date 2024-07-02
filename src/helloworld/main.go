package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		// Golang has 'fallthrough' keyword to go to the next branch of switch statement
		fallthrough
	case 5, 6, 7, 8:
		n := 99
		fmt.Println("n =", n)
		fallthrough
	default:
		fmt.Println("n =", n)
	}

	i := 0
Next:
	fmt.Println(i)
	i++
	if i < 5 {
		goto Next
	}
}
