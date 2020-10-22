package main

import (
	"fmt"
	"strings"
)

// https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
func main() {
	s := strings.Builder{}
	for i:= 0 ; i < 1000 ; i++ {
		s.WriteString("01")
	}
	fmt.Println(s.String())
}
