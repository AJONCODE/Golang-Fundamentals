package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Herera"), boring("Fread"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		}
	}()

	return c
}

// Fan In
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

// go run main.go
// Herera 0
// Fread 0
// Fread 1
// Herera 1
// Fread 2
// Fread 3
// Herera 2
// Herera 3
// Herera 4
// Fread 4
// You're both boring; I'm leaving.
