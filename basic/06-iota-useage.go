package main

import "fmt"

func main() {
    const (
        a1 = iota
        a2 
        a3
    )

    fmt.Println(a1, a2, a3)

    const (
        a = iota // 0
        b        // 1
        c        // 2
        d = "ha" // 独立值,iota += 1
        e        // "ha" iota += 1
        f = 100  // iota += 1
        g        // 100  iota += 1
        h = iota // 7,恢复计数
        i        // 8
    )

    fmt.Println(a, b, c, d, e, f, g, h, i)
    
    const (
        m = 1 << iota
        j = 3 << iota
        k
        l
    )

    fmt.Println("m = ", m)
    fmt.Println("j = ", j)
    fmt.Println("k = ", k)
    fmt.Println("l = ", l)
}
/*
iota 表示从0开始自动加1，所以m=1<<0,j=3<<1(<<表示左移的意思)，即m=1,j=6,这没问题，关键是k和l，从输出结果看k=3<<2,l=3<<3
m=1:左移0位，不变仍为1：
（1）j=3:左移1位，变为二进制110，即6
（2）k=3:左移2位，变为二进制1100，即12
（3）l=3:左移3位，变为二进制11000，即24
注意：<<n==*(2^n)
iota，特殊常量，可以被认为是一个可以被编译器修改的常量
iota在const关键字出现时，将被重置位0（const内部的第一行之前），const中每新增一行常量声明将使iota计数一次(iota可以理解位const语句块中的行索引)
iota可以被用作枚举值
const (
    a = iota
    b = iota
    c = iota
)
*/
