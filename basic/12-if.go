package main
import "fmt"

func main() {
    /*定义局部变量*/
    var a int = 10

    /*使用if玉菊判断布尔表达式*/
    if a < 20 {
        /*如果条件为true，则执行以下语句*/
        fmt.Printf("a 小于 20\n")
    }
    fmt.Printf("a的值:%d\n", a)
}
