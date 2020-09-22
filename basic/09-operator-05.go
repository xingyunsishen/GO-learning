package main
import "fmt"

func main() {
    var a  int = 4
    var b int32 
    var c float32
    var ptr *int

    /*运算符实例*/
    fmt.Printf("1. a变量类型为:%T\n", a)
    fmt.Printf("2. b变量类型为:%T\n", b)
    fmt.Printf("3. c变量类型为:%T\n", c)
    
    /* & 和 *运算符实例*/
    ptr = &a  /* 'ptr' 包含了'a'变量的地址*/
    fmt.Printf("a的值:%d\n", a)
    fmt.Printf("*ptr 为%d\n", *ptr)
}
