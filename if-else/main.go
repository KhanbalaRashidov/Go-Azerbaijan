package main

import "fmt"

func main() {

	x := 3

	if x > 0 {
		fmt.Println("Positive number")
	} else if x < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Zero")
	}

	if x > 10 && x < 20 {
		fmt.Println("x is between 10 and 20")
	} else if x > 20 && x < 30 {
		fmt.Println("x is between 20 and 30")
	} else {
		fmt.Println("x is not between 10 and 30")
	}
}
