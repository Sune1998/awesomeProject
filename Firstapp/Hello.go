package main

import (
	"fmt"
	"strconv"
)

// golbal vars can be called outside of a func

// vars can be chained together in a var() block
var (
	name    string = "sune"
	address string = "slagelse"
	zipcode int    = 4200
)

// different types of vars + hello world programm
func main() {
	// tree main types of var declaration
	var j int
	j = 3
	i := 42

	var a string
	a = strconv.Itoa(zipcode) // converts adrees into string

	var h int = 33

	// adding the i j h vars together
	var t int = j + i + h

	// print sum of the three vars
	fmt.Println(t)

	// prints global vars
	fmt.Println(name + address + a)

}
