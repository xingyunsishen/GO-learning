/*
Go语言提供了另外一种数据类型即接口，它把搜友的具有共性的定义在一起，任何其他
类型只要实现了这些方法就是实现了这个接口
*/
/*
//定义接口
type interface_name interface {
    method_name1 [return_type]
    method_name2 [return_type]
    method_name3 [return_type]
    ... 
    method_namen [return_type]
}
//定义结构体
type (struct_name_variable struct_name) method_name1() [return_type] {
    //方法实现
}
...
func  (struct_name_variable struct_name) method_namen() [return_type] {
    //方法实现
}
*/

package main
import (
    "fmt"
)

type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iphone IPhone) call() {
    fmt.Println("I am iphone, I can call you~")
}

func main() {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()

}
//在上面的例子中，我们定义了一个接口Phone，接口里面有一个方法call()。然后我
//们在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和
//IPhone。然后调用call()方法
