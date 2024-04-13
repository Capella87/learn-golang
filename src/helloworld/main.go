package main

import (
	"fmt"
)

const pi = 3.1416

const (
	No                = !Yes
	Yes               = true
	MaxDegrees        = 360
	Unit       string = "radian"
	A, B              = int64(-3), int64(5)
	Y                 = float32(2.718)
	MaxUnit           = ^uint(0)
	Lastiota          = iota
)

func main() {
	var ima complex64 = 1.2i                               // Imaginary numbers; complex64 and complex128
	var city1, city2 string = "San Francisco", "Daly City" // Multiple declaration at the same line
	var arbi1, arbi2 = true, 1.2                           // Like "var" behavior in .NET, and we can declare variables with several types at the same line
	var favoritegame = "Fallout 4"

	var n int
	var emptystring string

	// Short variable declarartion forms
	lang, year := "Golang", 2007

	fmt.Println(ima)
	fmt.Println(pi)
	fmt.Println(Lastiota)
	fmt.Printf("Cities: %s %s\n", city1, city2)
	fmt.Println(arbi1)
	fmt.Println(arbi2)
	fmt.Printf("Games: %s\n", favoritegame)
	fmt.Printf("Variable values without initialization: %d %s\n", n, emptystring)
	fmt.Printf("%s is made in %d\n", lang, year)

	// Mix the new one and already declared one, but at least one variable must be a new one
	thisyear, lang := 2024, ".NET"

	fmt.Printf("In %d, I think %s and Golang are the best programming languages\n", thisyear, lang)
}
