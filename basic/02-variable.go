package main
import "fmt"
func main() {
    //声明一个变量
    var a string = "Runoob"
    fmt.Println(a)
    //一次声明多个变量
    var b, c int = 1, 2
    fmt.Println(b, c)
    //bool零值为false   
    var d bool
    fmt.Println(d)
    //以下几种类型为nil
    var e *int
    var f []int
    var g map[string] int
    var h chan int
    var j func(string) int
    var k error //error是接口
    fmt.Println(e, f, g, h, j, k)
    
    var l int
    var m float64
    var n bool
    var o string
    fmt.Println("%v %v %v %q\n", l, m, n, o)

    //根据值自行判定值类型
    var p = true
    fmt.Println(p)
    //省略var与否
    var q int
    //q := 1
    //fmt.Println(q)
    /*
    报错：33:7: no new variables on left side of :=
    */
    q, q1 := 1, 2
    fmt.Println(q, q1)
    res := "Runoob" //等价于var res string = "Runoob"
    fmt.Println(res)
}

