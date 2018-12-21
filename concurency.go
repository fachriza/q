package main

import (
	"fmt"
	"time"
)

func print() {
	fmt.Println("Printing from goroutine")
}

func main() {
	go print()
	time.Sleep(1 * time.Second)
	fmt.Println("Printing from main")
}
