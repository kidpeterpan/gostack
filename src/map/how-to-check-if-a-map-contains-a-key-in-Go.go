package main

import "fmt"

// https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
func main() {
	myMap := map[string]int {}
	myMap["one"] = 1
	myMap["two"] = 2

	if value, ok := myMap["one"]; ok {
		fmt.Println("the map contains:",value)
	}
}
