/*
goroutine是轻量级线程，goroutine的调度室有Golang运行时进行管理的。
格式：
go function_name(参数列表)
*/
//Go允许使用go语句开启一个新的运行期线程，即goroutine,以一个不同的、新创建的
//goroutine来执行一个函数。同一个程序中的所有的goroutine共享一个地址空间

package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("2. world")
    say("1. hello")
}
