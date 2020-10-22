package main

import (
	"fmt"
	"os"
)

// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func main() {
	if _, err := os.Stat("./src/file/test.txt"); os.IsNotExist(err) { //path start from root of project
		fmt.Println("Path does not exit!")
	} else if err == nil {
		fmt.Println("We found the file!")
	} else {

	}
}
