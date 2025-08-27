// kr_sqr.go |  SIMPLE RECURSIVE POWER FUNCTION
package main 
import "fmt"

func pwr(x int, n int) int {
    // edge case
    if n == 0 {
        return 1
    }
    
    return x * pwr(x, n - 1)
}

func main() {
    var x, y int
    var result int
    
    fmt.Print("Enter x: ")
    fmt.Scan(&x)
    fmt.Print("Enter n: ")
    fmt.Scan(&y)
    
    result = pwr(x, y)
    fmt.Println("The result is:", result)
}
