package main

import "fmt"

var s = []string{"h", "i", "j", "k", "l", "m", "n", "o", "p", "u"}

func chgSlice(s *[]string) {
	(*s)[5] = "hello"
	(*s)[7] = "world"
	(*s)[9] = "hi"
}

func main() {
	fmt.Printf("Original slice: %s\n", s)
	chgSlice(&s)
	fmt.Printf("Updated slice: %s\n", s)
}
