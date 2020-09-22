/*
类型转换用于将数据类型的变量转换为另一种类型的变量。Go语言类型转换的格式如
下：
type_name(expression)
type_name为类型，expresssion为表达式
*/
package main

import "fmt"

func main() {
    var sum int = 17
    var count int = 5
    var mean float32

    mean = float32(sum)/float32(count)
    fmt.Printf("mean 的值为: %f\n", mean)
} 
