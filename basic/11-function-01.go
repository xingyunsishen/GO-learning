package main
import "fmt"
/*
函数返回多个值
*/
func swap(x, y string) (string, string) {
    return y, x
}   

func main() {
    a, b := swap("Hello", "GOLANG")
    fmt.Println(a, b)
}
