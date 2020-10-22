package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

// https://stackoverflow.com/questions/10105935/how-to-convert-an-int-value-to-string-in-go
func main() {
	myInt := 777
	strMyInt := strconv.Itoa(myInt)
	fmt.Println(reflect.TypeOf(strMyInt)) // string
	fmt.Println(strMyInt) // 777
	// and if you want to convert string to int
	if intStrMyInt, err := strconv.Atoi(strMyInt); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(reflect.TypeOf(intStrMyInt)) // int
		fmt.Println(intStrMyInt) // 777
	}

}
