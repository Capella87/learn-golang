package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Function for Goroutine
func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}

	// Notify when the task is finished.
	wg.Done() // <=> wg.Add(-1)
}

func LetsGoRoutine() {
	log.SetFlags(0)

	// Invoke Goroutines
	// We Invoked these, but no synchronization between them
	// Synchronization methods in Golang are various
	// Channel - The most popular
	// WaitGroup from sync - Simple
	wg.Add(2)
	go SayGreetings("Hi!", 10)
	go SayGreetings("Hello!", 10)

	// Block until all tasks are completed
	wg.Wait()
}
