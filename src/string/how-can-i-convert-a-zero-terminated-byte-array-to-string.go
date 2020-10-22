package main

import "fmt"

// https://stackoverflow.com/questions/14230145/how-can-i-convert-a-zero-terminated-byte-array-to-string

func main() {
	// one byte equals to 8 bit
	// use string() to convert byte to string
	// byte array tip is string(byteArray[:len(byteArray)]) OR [:]
	byteArray := []byte{65,66}
	fmt.Println(string(byteArray[:])) // byteArray 0 -> before 2
}
