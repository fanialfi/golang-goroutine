package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestRunHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("ups")

	time.Sleep(time.Second * 2)
}

func TestManyGorutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(time.Second * 3)
}
