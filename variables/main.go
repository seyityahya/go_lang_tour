package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var productName string
	// var quantity int
	// var discount float32

	// productName = "Laptop"
	// quantity = 10
	// discount = 0.10

	// fmt.Println(productName, reflect.TypeOf(productName))
	// fmt.Println(quantity, reflect.TypeOf(quantity))
	// fmt.Println(discount, reflect.TypeOf(discount))

	productName := "Laptop"
	fmt.Println(productName, reflect.TypeOf(productName))
}
