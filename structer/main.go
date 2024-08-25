package main

import "fmt"

func main() {
	var structer = Structer{id: 1, name: "Ali", age: 20, adress: Adress{city: "İstanbul", country: "Turkey"}}
	// var structer2 = Structer{id: 2, name: "Veli", age: 30, adress: Adress{city: "Ankara", country: "Turkey"}}

	changeName(&structer)
	fmt.Println(structer)
}

func changeName(s *Structer) {
	s.name = "Mehmet"
}

type Structer struct {
	id     int64
	name   string
	age    int
	adress Adress
}
type Adress struct {
	city    string
	country string
}
