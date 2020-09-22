package main
import "fmt"
/*
初始化局部变量和全局变量
数据类型            初始化默认值
int                 0
float32             0
pointer             nil
*/
//声明全局变量
var a int = 20;

func main() {
    //main函数中声明局部变量
    var a int = 10
    var b int = 20
    var c int = 0

    fmt.Printf("main()函数中 a = %d\n", a);
    c = sum(a, b);
    fmt.Printf("main()函数中 c = %d\n", c);
}

//函数定义-两数相加
func sum(a, b int) int {
    fmt.Printf("sum() 函数中a = %d\n", a);
    fmt.Printf("sum() 函数中b = %d\n", b);

    return a + b;
}
