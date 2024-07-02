package main

import (
	"log"
	"math/rand"
	"time"
)

// Function for Goroutine
func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
}

func LetsGoRoutine() {
	log.SetFlags(0)

	// Invoke Goroutines
	// We Invoked these, but no synchronization between them
	go SayGreetings("Hi!", 10)
	go SayGreetings("Hello!", 10)
	time.Sleep(2 * time.Second)
}
