package main
/*
函数是基本的代码块，用于执行一个任务
GO语言最少有个main()函数
函数定义格式:
func function_name( [parameter list] ) [return_types] {
    函数体
}
函数定义解析：
func:函数有func开始声明
function_name:函数名称，函数名和参数列表一起构成了函数签名
parameter list:参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的时参数类型、顺序、及参数的个数。参数是可选的，也就是说函数也可以不包含参数
return_types:返回类型，函数返回一列值。return_types是该列值的数据类型。有些功能不需要返回值，这种情况下return_types不是必须的
函数体：函数定义的代码集合
*/
import "fmt"

func main() {
    /*定义局部变量*/
    var a, b int = 100, 200
    var ret int

    /*调用函数并返回最大值*/
    ret = max(a, b)

    fmt.Printf("最大值:%d\n", ret)
}

/*函数返回两个数的最大值*/
func max(num1, num2 int) int {
/*定义局部变量*/
    var result int
    if (num1 > num2) {
        result = num1
    } else {
        result = num2
    }
    return result
}
