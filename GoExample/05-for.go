package main

import(
    "fmt"
)

func main() {
    i := 1
    //最基本的单条件循环
    for i <= 3 {
        fmt.Println("i =", i)
        i++
    }
    //经典的初始-条件-后，for循环
    for j := 7; j <= 9; j++ {
        fmt.Println("j =", j)
    }
    //for在没有条件的情况将反复循环，直到break退出循环或return从封闭的函数中
    //退出
    for {
        fmt.Println("loop")
        break
    }
    for n := 0; n <= 5; n++ {
        if n % 2 == 0 {
            //也可以使用continue转到下一个迭代
            continue
        }
        fmt.Println("n =", n)
    }
}
