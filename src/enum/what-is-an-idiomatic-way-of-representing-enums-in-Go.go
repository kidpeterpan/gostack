package main

import "fmt"

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

const (
	a = 1 << iota
	b = 1 << iota
	c = 1 << iota
)

const (
	u = iota * 42
	v float64 = iota * 42
	w = iota * 42
)

const x = iota
const y = iota

// https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go
func main() {
	fmt.Println(c0) // 0
	fmt.Println(c1) // 1
	fmt.Println(c2) // 2

	fmt.Println(a) // 1
	fmt.Println(b) // 2
	fmt.Println(c) // 4

	fmt.Println(u) // 0
	fmt.Println(v) // 42
	fmt.Println(w) // 84

	fmt.Println(x) // 0
	fmt.Println(y) // 0
}
