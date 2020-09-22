package main
import "fmt"
func main() {
    //使用var声明一个变量
    var a string = "Runoob"
    fmt.Println(a) 
    
    var b, c int = 1, 2
    fmt.Println(b, c) 
    //指定变量类型，如果没有初始化，则变量默认为零值
    //声明一个变量并初始化
    var a1 = "RUNOOB"
    fmt.Println(a1)
    
    //没有初始化就为零值
    var b1 int 
    fmt.Println(b1)
    
    //bool 零值为false
    var c1 bool 
    fmt.Println(c1)
    
    var i int
    var f float64
    var b2 bool
    var s string
    fmt.Printf('%v %v %v %q\n', i, f, b2, s)
}   

