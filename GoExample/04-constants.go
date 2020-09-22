package main

import (
    "fmt"
    "math"
)
//声明一个常量
const s string = "constant"

func main() {
    fmt.Println("s的值:", s)
    //常量表达式以任意精度执行算术运算
    const n = 500000000
    //在给定数值常量之前，数字常量是没有类型的，例如通过显示转换
    //可以通过在需要数字的上下文中使用数字类为其指定类型，例如变量赋值或函数
    //调用。例如，这里math.Sin期望一个faloat64
    const d = 3e20 / n
    fmt.Println("d = ", d)
    fmt.Println("d = ", int64(d))
    fmt.Println("d = ", math.Sin(n))
}

