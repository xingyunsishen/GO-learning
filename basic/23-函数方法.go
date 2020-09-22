package main
/*
GO语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或是一个指针。所有给定类型的方法舒予该类型的方法集
func (variable_name variable_data_type) function_name() [return_type] {
    //函数体
}
*/
import (
    "fmt"
)

//定义结构体
type Circle struct {
    radius float64
}

func main() {
    var c1 Circle
    c1.radius = 10.00
    fmt.Println("圆的面积 = ", c1.getArea())
}
//该method舒予Circle类型对象中的方法
func (c Circle) getArea() float64 {
    //c.radius 即为Circle类型对象中的属性
    return 3.14 * c.radius * c.radius
}
