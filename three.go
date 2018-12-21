package main

import "fmt"

func FizzBuzz(number int) string {
	var s string
	if number%3 == 0 {
		s += "Fizz"
	}

	if number%5 == 0 {
		s += "Buzz"
	}
	return s
}

func main() {
	for i := 1; i <= 100; i++ {
		fb := FizzBuzz(i)
		if len(fb) == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(fb)
		}
	}
}
