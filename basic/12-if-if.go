package main
import "fmt"

func main() {
    //定义局部变量
    var a int = 100
    var b int = 200

    //判断条件
    if a == 100 {
        //if条件为true执行
        if b == 200 {
            //if条件为true执行
            fmt.Printf("a的值:100, b的值:200\n");
        }
    }
    fmt.Printf("a 的值：%d\n", a);
    fmt.Printf("b 的值：%d\n", b);
}
