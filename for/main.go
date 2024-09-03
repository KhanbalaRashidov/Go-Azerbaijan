package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	i := 0
	for {
		if i == 100 {
			break
		}
		fmt.Println(i)

		i++
	}
}
