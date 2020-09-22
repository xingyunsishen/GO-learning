package main
import "fmt"

func main() {
    var a bool = true
    var b bool = false
    if ( a && b ) {
        fmt.Printf("1. 条件为true\n")
    }
    if ( a || b ) {
        fmt.Printf("2. 条件为true\n")
    }
    /*修改a和b的值*/
    a = false
    b = true
    if ( a && b ) {
        fmt.Printf("3. 条件为true\n")
    } else {
        fmt.Printf("3. 条件为false\n")
    }
    if ( !( a && b )) {
        fmt.Printf("4. 条件为true\n")
    }
}
