package main
import "fmt"

func main() {
    //定义局部变量
    var a int = 100
    var b int = 200

    fmt.Printf("交换前a的值:%d\n", a)
    fmt.Printf("交换前b的值:%d\n", b)

    //通过调用函数来交换值
    swap(a, b)

    fmt.Printf("交换后a的值:%d\n", a)
    fmt.Printf("交换后b的值:%d\n", b)
}
//定义相互交换值的函数
func swap(x, y int) int {
    var temp int

    temp = x //保存x的值
    x = y //将y值赋给x
    y = temp //将temp值赋给y

    return temp;
}
