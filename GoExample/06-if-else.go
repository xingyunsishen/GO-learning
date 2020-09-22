package main

import (
    "fmt"
)

func main() {
    if 7 % 2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8 % 4 == 0 {
        fmt.Println("8 is divisible by 4")
    }
    //这里的num变量在以下的分支中均可以使用
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
    //Go中没有三元组，因此if即使在基本条件下也需要使用完整的语句
}
