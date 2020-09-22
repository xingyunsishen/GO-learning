package main

import(
    "fmt"
    "runtime"
)

func main() {
    println("Hello", "Golang")
    fmt.Printf("Go 版本 :%s\n", runtime.Version())
}

