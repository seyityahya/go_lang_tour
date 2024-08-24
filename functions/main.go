package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")
	// var numbers = []int{1, 2, 3, 4, 5}
	// fmt.Println(add(numbers))
	fmt.Println(sum(1, 2, 3, 4, 5))
}

func add(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sum(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
