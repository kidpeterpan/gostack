package main

import "fmt"

// https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
func main() {
	firstSlice := []int{1,2,3}
	secondSlice := []int{4,5,6}
	firstSlice = append(firstSlice, secondSlice...)
	fmt.Println("concated slice...")
	fmt.Println(firstSlice)
}
