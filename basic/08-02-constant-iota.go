/*
iota 特殊常量，可以认为是一个可以被编译器修改的常量
iota在const关键字出现时将被重置为0(const内部的第一行之前)const中每新增一行
常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
*/
package main
import "fmt"

func main() {
    const (
        a = iota //0
        b        //1
        c        //2
        d = "ha" //独立值,iota += 1
        e        //“ha”, iota += 1
        f = 100  //iota += 1
        g        //100 iota += 1
        h = iota //7,恢复计数
        i        //8
    )
    fmt.Println(a, b, c, d, e, f, g, h, i)
}
