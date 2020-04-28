package main

import (
	"fmt"
	"runtime"
)

/*
Create a program that prints out your OS and ARCH. Use the following commands to run it
- go run
- go build
- go install
*/

func main() {
	fmt.Printf("OS: %v\nARCH: %v", runtime.GOOS, runtime.GOARCH)
}
