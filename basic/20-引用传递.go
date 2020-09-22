package main
import "fmt"

func main() {
    //定义局部变量
    var a int = 100
    var b int = 200

    fmt.Printf("交换前,a的值:%d\n", a)
    fmt.Printf("交换前,b的值:%d\n", b)

    /*调用swap()函数
    * &a 指向a指针,a变量的地址
    * &b 指向b指针,b变量的地址
    */
    swap(&a, &b)

    fmt.Printf("交换后,a的值:%d\n", a)
    fmt.Printf("交换后,b的值:%d\n", b)
}

func swap(x *int, y *int) {
    var temp int
    temp = *x //保存x地址上的值
    *x = *y   //将y值赋给x
    *y = temp  //将temp值赋给y
}
