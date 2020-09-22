/*
作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围
GO语言中变量可以在三个地方声明：
函数内定义的变量为局部变量
函数外定义的变量为全局变量
函数定义中的变量成为形式参数
*/
//局部变量
package main
import "fmt"

func main() {
    //声明局部变量
    var a, b, c int
    //初始化参数
    a = 10
    b = 20
    c = a + b

    fmt.Printf("结果:a = %d, b = %d and c = %d\n", a, b, c)
}
