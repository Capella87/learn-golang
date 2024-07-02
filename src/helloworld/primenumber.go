package main

import "fmt"

func FindSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; n++ {
		for i := 2; ; i++ {
			switch {
			case i*i > n:
				break Outer
			case n%i == 0:
				continue Outer
			}
		}
	}
	return n
}

func PrimeNumberExample() {
	for i := 90; i < 100; i++ {
		n := FindSmallestPrimeLargerThan(i)
		fmt.Println("The smallest prime number larger than")
		fmt.Println(i, "is", n)
	}
}
