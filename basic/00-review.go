package main
import "fmt"
/*
变量声明：
第一种：指定变量类型，如果没有初始化，则默认为零值
*//*
func main() {
    // 声明一个变量并初始化
    var a = "GoLang perferect"
    fmt.Println(a)

    //没有初始化就为零值
    var b int
    fmt.Println(b)

    //bool零值为false
    var c bool
    fmt.Println(c)
}
*//*
func main() {
   var i int
   var j float64
   var k bool
   var l string
   fmt.Printf("%v %v %v %q\n", i, j, k, l)
}
*/

//第二种：根据值自行判定变量类型
/*
func main() {
    var d = true
    fmt.Println(d)
}*/

//第三种：省略var，注意:=左侧如果没有声明新的变量，就产生编译错误
/*
func main() {
    f := "GO Language" //<=> var f string = "GO Language"
    fmt.Println(f)
}
*/

//多变量声明
var x, y, z int //类型相同多个变量，非全局变量
var (           //这种因式分解关键字的写法一般用于声明全局变量
    res string
    bu bool
)

var c, d int = 1, 2
var e, f = 123, "hello"
//g, h := 111, "hahaha"  //这种不带声明格式的只能在函数体中出现
func main() {
    g, h := 111, "hahaha"
    //fmt.Println(x, y, z, res, bu, c, d, e, f, g, h)
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
    fmt.Println("%v",res)
    fmt.Println(bu)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)
    fmt.Println(f)
    fmt.Println(g)
    fmt.Println(h)
}
