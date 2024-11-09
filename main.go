package main

import "fmt"

func main() {
	var number int = 43
	var word string = "Something"
	var float float64 = 4.31
	var condition bool = false
	fmt.Printf("%d %q %.1f %v", number, word, float, condition)
}
