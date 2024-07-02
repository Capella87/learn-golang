package main

import (
	"fmt"
	"runtime"

	// Full package import form
	"math/rand"
	// securerand "crypto/rand" // More secure
	// Dot import (Can invoke methods without package name and dot
	. "time"

	// Anonymous (a.k.a blank) import
	_ "net/http/pprof"
)

const pi = 3.1416

// Constants
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

func base() {
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
	fmt.Printf("Games: %s\n", favoritegame)
	fmt.Printf("Variable values without initialization: %d %s\n", n, emptystring)
	fmt.Printf("%s is made in %d\n", lang, year)

	fmt.Println(Now())

	// Mix the new one and already declared one, but at least one variable must be a new one
	thisyear, lang := 2024, ".NET"

	fmt.Printf("In %d, I think %s and Golang are the best programming languages\n", thisyear, lang)

	// Make arbi1 and arbi2 be used... anyway!
	_, _ = arbi1, arbi2

	// Anonymous Functions
	x, y := func() (int, int) {
		fmt.Println("This function has no parameters.")
		return 8, 7
	}()

	// Has inputs and no returns
	func(a, b int) {
		// Parameters in anonymous function shadow outer scope
		fmt.Println("a*a + b*b =", a*a+b*b)
	}(x, y) // Pass two parameters

	// No inputs, no returns
	func() {
		fmt.Println("Ken Thompson, Rob Pike, and Robert Griesemer made Go language.")
	}()

	r := rand.Uint32()
	fmt.Printf("Next pseudo-random number is %v (%v)\n", r, r)

	fmt.Printf("The type of variable r is %T\n", r)

	// Range
	for i := range 10 {
		fmt.Print(i, "\n")
	}

	// Switch
	fmt.Print("You're running Go on ")
	switch whichOs := runtime.GOOS; whichOs {
	case "windows":
		fmt.Println("Microsoft Windows")
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	}
	fmt.Println("You can see all possible cases by `go tool dist list`")
}
