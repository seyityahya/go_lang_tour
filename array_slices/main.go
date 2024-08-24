package main

import "fmt"

func main() {
	// var productNames [3]string
	// productNames[0] = "Laptop"

	var productNames = []string{"Laptop", "Mouse", "Keyboard"}
	productNames = append(productNames, "Monitor")
	fmt.Println(productNames)
}
