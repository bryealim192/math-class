
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	result := x + y
	fmt.Println("The sum of", x, "and", y, "is", result)
}