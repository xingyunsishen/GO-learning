package main
import "fmt"

func main() {
    var a int = 21
    var c int

    c = a
    fmt.Printf("1. c = %d\n", c)

    c += a
    fmt.Printf("2. %d += %d = %d\n", c, a, c)

    c -= a
    fmt.Printf("3. %d -= %d = %d\n", a, c, c)
    
    c *= a
    fmt.Printf("4. %d *= %d = %d\n", a, c, c)
    
    c /= a
    fmt.Printf("5. %d /= %d = %d\n", a, c, c)

    c = 200
    c <<= 2
    fmt.Printf("6. %d << 2 = %d\n", c, c)

    c >>= 2
    fmt.Printf("7. %d >>= 2 = %d\n", c, c)

    c &= 2
    fmt.Printf("8. %d &= 2 = %d\n", c, c)

    c ^= 2
    fmt.Printf("9. %d ^= 2 = %d\n", c, c)

    c |= 2
    fmt.Printf("10 %d |= 2 = %d\n", c, c)
}
