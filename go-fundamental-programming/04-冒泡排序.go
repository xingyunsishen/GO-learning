package main

import (
    "fmt"
)

func main() {
    //为排序数组
    sort := [...]int{3, 2, 4, 5, 1, 9, 0}
    fmt.Println(sort)

    //冒泡排序，由大到小
    num := len(sort)
    for i := 0; i < num; i ++ {
        for j := i + 1; j < num; j++ {
            //比较大小
            if sort[i] > sort[j] {
                temp := sort[i]
                sort[i] = sort[j]
                sort[j] = temp
            }
        }
    }
    fmt.Println(sort)
}
