package main

import "fmt"
func main() {
    //定义局部变量
    var a int = 10
    //for循环
    for a < 20 {
        fmt.Printf("a的值为:%d\n", a);
        a++;
        if a > 15 {
            //使用break语句跳出循环
            break;
        }
    }
}
