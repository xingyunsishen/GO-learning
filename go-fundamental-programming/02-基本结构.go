//当前程序的包名
package main

//导入其他的包
import std "fmt"

//常量的定义
const PI = 3.14

//全局变量的声明与赋值
var name = "global_variable"

//一般类型声明
type newType int

//结构的声明
type global_variable struct{}

//接口的声明
type golang interface{}

//由main函数作为程序入口点启动
func main() {
    std.Println("Hello Golang~,Nice to meet you~")
}
