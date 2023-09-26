package main

import "fmt"

func main() {
    // Declare and initialize an array of integers
    var numbers [5]int
    numbers[0] = 1
    numbers[1] = 2
    numbers[2] = 3
    numbers[3] = 4
    numbers[4] = 5

    // Declare and initialize an array of strings
    fruits := [3]string{"Apple", "Banana", "Cherry"}

    // Declare and initialize an array with inference of size
    colors := [...]string{"Red", "Green", "Blue"}

    // Accessing elements
    fmt.Println(numbers[0]) // Prints 1
    fmt.Println(fruits[1])  // Prints "Banana"
    fmt.Println(colors[2])  // Prints "Blue"

    // Length of an array
    fmt.Println(len(numbers)) // Prints 5
    fmt.Println
