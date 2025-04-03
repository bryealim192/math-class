package main

import (
    "fmt"
)

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)
    
    if a > 0 && b > 0 {
        sum := a + b
        multiply := a * b
        print("The sum of two positive integers is:", sum)
        print("The product of the two positive integers is:", multiply)
    } else {
        fmt.Println("Invalid input. At least one number must be positive.")
    }
}
