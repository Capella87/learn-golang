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
)

func main() {
	var ima complex64 = 1.2i

	fmt.Println(ima)
	fmt.Println(pi)
}
