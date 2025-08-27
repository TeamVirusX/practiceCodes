package main 
import "fmt"

func divide (numerator int, denominator int) (int, int) {
    var quotient, remainder int;
    quotient = numerator / denominator
    remainder = numerator % denominator
    return quotient, remainder
}

func main() {
    var x, y int
    var quotient, remainder int
    
    fmt.Print("Enter first number: ")
    fmt.Scan(&x)
    fmt.Print("Enter second number: ")
    fmt.Scan(&y)
    
    // unbox and print values
    quotient, remainder = divide(x, y)
    fmt.Println("Quotient =", quotient)
    fmt.Println("Remainder =", remainder)
}
