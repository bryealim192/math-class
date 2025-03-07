
package main
import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) float64 {
	if b == 0 {
		panic("Cannot divide by zero")
	}
	return float64(a) / float64(b)
}

func main() {
	fmt.Println(Add(3, 5))
	fmt.Println(Subtract(7, 2))
	fmt.Println(Multiply(4, 9))
	fmt.Println(Divide(10, 2))
}