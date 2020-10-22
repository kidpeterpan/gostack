package main

import "fmt"

// https://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go
func main() {
	mySlice := []int{0,1,2,3,4,5,6,7,8,9}

	for index, value := range mySlice {
		fmt.Println(index,value)
	}

	for _, value := range mySlice {
		fmt.Println("use only value:",value)
	}

	for index, _ := range mySlice {
		fmt.Println("use only index:",index)
	}
}
