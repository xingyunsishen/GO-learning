package main
//import "fmt"
import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main() {
    println(a, b, c)
}
/*
func main() {
    const LENGTH int = 10
    const WIDTH int = 5
    var area int
    const a, b, c = 1, false, "str" //多重赋值

    area = LENGTH * WIDTH
    fmt.Printf("面积为: %d", area)
    println()
    println(a, b, c)
}
*/
