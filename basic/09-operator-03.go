package main
import "fmt"

func main() {
    var a uint = 60 /* 60 = 0011 1100*/
    var b uint = 13 /* 13 = 0000 1101*/
    var c uint = 0 
    
    c = a & b      /* 12 = 0000 1100*/
    fmt.Printf("1. %d & %d = %d\n", a, b, c)
    
    c = a | b     /* 61 = 0011 1101*/
    fmt.Printf("2. %d | %d = %d\n", a, b, c)

    c = a ^ b     /* 49 = 0011 0001*/
    fmt.Printf("3. %d ^ %d = %d\n", a, b, c)

    c = a << 2    /* 240 = 1111 0000*/
    fmt.Printf("4. %d << 2 = %d\n", a, c)

    c = a >> 2    /* 15 = 0000 1111*/
    fmt.Printf("5. %d >> 2 = %d\n", a, c)
}
