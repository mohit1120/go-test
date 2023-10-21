package main

import (
	"fmt"
	"go-test/check"
)

func main() {
	fmt.Println("Hello world")
	result := check.Check("abc", "abc")
	fmt.Println(result)
}
