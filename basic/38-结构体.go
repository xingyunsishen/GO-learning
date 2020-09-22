/*GO语言数组中可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性:
·Title：标题
·Author：作者
·Subject：学科
·ID：书籍id
定义结构体需要使用type和struct语句。struct语句定义一个新的数据类型，结构体中有一个或多个成员。type语句设定了结构体的名称。结构体格式如下：
type struct_variable_type struct {
    member definition
    member definition
    ...
    member definition
}
一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
variable_name := structure_variable_type {value1, value2,...,valuem}
或者
variable_name := structure_variable_type {key1: value1, key2:value2,..., keyn: valuen}
*/
package main

import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main() {
    //创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.golang.org", "Go 语言教程", 6871233})
    //也可以使用key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.golang.org", subject: "Go 语言编程", book_id: 68712333})
    //忽略的字段为0或空
    fmt.Println(Books{title: "Go 语言", author: "www.golang.org"})
}

