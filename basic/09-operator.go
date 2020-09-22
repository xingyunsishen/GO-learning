package main
import "fmt"

func main() {
    var a int = 21
    var b int = 10
    var c int

    c = a + b
    fmt.Printf("第一行 a + b = %d\n", c)
    c = a - b
    fmt.Printf("第二行 a - b = %d\n", c)
    c = a * b
    fmt.Printf("第三行 a * b =  %d\n", c)
    c = a / b
    fmt.Printf("第四行 a / b =  %d\n", c)
    c = a % b
    fmt.Printf("第五行 a % b =  %d\n", c)
    a++
    fmt.Printf("第六行 a++ =  %d\n", a)
    a = 21
    a --
    fmt.Printf("第七行 a--  %d\n", a)
}
