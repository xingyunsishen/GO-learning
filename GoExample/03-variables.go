package main

import (
    "fmt"
)

func main() {
    //var声明一个局部变量
    var a = "initial"
    fmt.Println("a = ", a)
    //声明多个局部变量
    var b, c int = 1, 2
    fmt.Println("b =" , b, "c =", c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Printf("变量e是%T类型默认值是%d\n", e, e)

    f := "apple" //等价与var f string = "apple"
    fmt.Println("f =", f)
    fmt.Printf("变量f的类型%T\n", f)
}
