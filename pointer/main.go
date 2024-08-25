package main

import "fmt"

func main() {
	// var a int
	// var p *int
	// a = 10
	// p = &a

	// a = 20
	// fmt.Println(a, p)

	// Burada a değişkeni fonksiyon içinde değişmiyor. Çünkü fonksiyon içinde a değişkeninin bir kopyası oluşturuluyor.
	var a = 10
	add12(a)
	fmt.Println(a)
	add12Pointer(&a)
	fmt.Println(a)
}
func add12(a int) {
	a = a + 12
}

func add12Pointer(a *int) {
	*a = *a + 12
}
