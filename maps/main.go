package main

import "fmt"

func main() {
	var colors map[string]string
	colors = make(map[string]string)
	colors["red"] = "#FF0000"
	colors["green"] = "#00FF00"
	colors["blue"] = "#0000FF"
	fmt.Println(colors)

	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}
	fmt.Println(ages)
}
