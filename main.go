package main

import (
	"fmt"
	"reflect"
)

func main() {
	var packagesName = "Go Packages"
	var name string

	year := 1998
	fmt.Printf("I was born in %d.", year)
	fmt.Println(reflect.TypeOf(year))

	fmt.Println(" Go programs organized into", packagesName)
	fmt.Println("You should look at the documentation to find out which function is in which", packagesName)

	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("Nice to meet you", name)
}
