package main

import (
	"fmt"
	"math/rand"
)

func switExamples() {
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
	if i >= 5 {
		goto Exit
	}

	// Shrink the scope of k by an explicit code block
	// Or we can declare the variable above 'goto'
	{
		k := i + i
		fmt.Println(k)
	}
	i++
	goto Next
Exit:
	fmt.Println("\nBye")
}
