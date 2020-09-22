package main
import "fmt"

func main() {
    //定义局部变量
    var a int = 100
    var b int = 200

    fmt.Printf("交换前a的值: %d\n", a)
    fmt.Printf("交换前b的值: %d\n", b)

    /*
     调用函数用于交换值
    * &a 指向a变量的地址
    * &b 指向b变量的地址
    */
   //swap(&a, &b); //这里也可以不用调用swap函数
   //直接写为
    a, b = b, a

    fmt.Printf("交换后a的值: %d\n", a)
    fmt.Printf("交换后b的值: %d\n", b)
}
/*
func swap(x *int, y *int) {
    var temp int
    temp = *x   //保存x的值
    *x = *y     //将y的值赋给x
    *y = temp   //将temp赋给y
}
*/
//swap交换函数可以简化为
/*
func swap(x *int, y *int) {
    *x, *y = *y, *x
}
*/
