package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, world!")
	var i int = 42
	var str string = strconv.Itoa(i)
	fmt.Println(str)
}
