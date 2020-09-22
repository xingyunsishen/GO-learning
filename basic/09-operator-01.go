package main
import "fmt"

func main() {
    var a int = 21
    var b int = 10

    if( a == b ) {
        fmt.Printf("1, %d == %d\n", a, b)
    }else {
        fmt.Printf("1, %d != %d\n", a, b)
    }
    if ( a < b ) {
        fmt.Printf("2, %d < %d\n", a, b)
    }else {
        fmt.Printf("2, %d >= %d\n", a, b)
    }
    if ( a > b ) {
        fmt.Printf("3, %d > %d\n", a, b)
    } else {
        fmt.Printf("3, %d <= %d\n", a, b)
    }
    if ( a <= b ) {
        fmt.Printf("4, %d <= %d\n", a, b)
    } else {
        fmt.Printf("4, %d >= %d\n", a, b)
    }
    if ( a >= b ) {
        fmt.Printf("5, %d >= %d\n", a, b)
    } else {
        fmt.Printf("5, %d <= %d\n", a, b)
    }
}
