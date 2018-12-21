package main

import "fmt"
import "strconv"

func Odd(number int) bool {
	return number%2 != 0
}

func main() {
	fmt.Print("Enter Number: ")

	var (
		input string
		data  []int
	)

	fmt.Scanln(&input)
	number, _ := strconv.Atoi(input)

	for i := 1; i <= number; i++ {
		if Odd(i) {
			data = append(data, i)
		}
	}

	fmt.Println(data[:len(data)-1])
}
