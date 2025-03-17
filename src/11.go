package main

import "fmt"

func main() {
	result := 0
	for i := 1; i <= 10; i++ {
		result += i
	}
	fmt.Println(result)
}
