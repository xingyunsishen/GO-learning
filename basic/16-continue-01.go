package main

import "fmt"

func main() {
    //不使用标记
    fmt.Println("-----cintinue-----")
    for i := 1; i <= 3; i++ {
        fmt.Printf("i:%d\n", i)
            for i2 := 11; i2 <= 13; i2++ {
                fmt.Printf("i2:%d\n", i2)
                continue
            }
    }
    //使用标记
    fmt.Println("-----continue label-----")
    re:
        for i := 1; i <= 3; i++ {
            fmt.Printf("i:%d\n", i)
                for i2 := 11; i2 <= 13; i2++ {
                    fmt.Printf("i2:%d\n", i2)
                    continue re
                }
        }
}
