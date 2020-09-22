package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {
    //int
    var num_01 int = 1
    fmt.Println("num_01 = %d", num_01)  
    fmt.Println("这里输出时会把%d原样输出")
    fmt.Printf("num_01 = %d \n", num_01) //Printf只可以打印出格式化的字符串，可以输出字符串类型的变量
    fmt.Printf("num_01 = %v \n", num_01) //输出1
    fmt.Printf("num_01 = %05d \n", num_01) //输出00001

    //bool
    var result bool = false
    fmt.Printf("result = %t \n", result) //result = false %t格式化布尔类型
    fmt.Printf("result = %s \n", result) //输出%!s(bool=false)

    //char
    var ch int = '\u0041'
    var ch2 int = '\u03B2'
    //var ch3 int = '\u00101234'
   // fmt.Printf("%d - %d - %d", ch, ch2, ch3) //interger，输出65 - 946 - 105326
   // fmt.Printf("%c - %c - %c", ch, ch2, ch3) //character,输出A - β - r
   // fmt.Printf("%X - %X - %X", ch, ch2, ch3) //UTF-8 bytes，输出41 - 3B2 - 101234
   // fmt.Printf("%U - %U - %U", ch, ch2, ch3) //UTF-8 code point，输出U+0041 - U+03B2 - U+101234
    fmt.Printf("%d - %d\n", ch, ch2) //interger，输出65 - 946 - 105326
    fmt.Printf("%c - %c\n", ch, ch2) //character,输出A - β - r
    fmt.Printf("%X - %X\n", ch, ch2) //UTF-8 bytes，输出41 - 3B2 - 101234
    g := func(i int) { fmt.Printf("%d", i)}
    fmt.Printf("g is of type %T and has value %v\n", g, g) //%T打印类型的完整说明
    
    //打印point结构体
    p := point{1, 2}
    fmt.Printf("%v\n", p) //输出 {1， 2}
    
    //如果是一个结构体，%+v的格式化输出内容将包括结构体的字段名
    fmt.Printf("%+v\n", p) // 输出 {x:1, y:2}

    //%#v形式输出这个值的Go语法表示
    fmt.Printf("%#v\n", p) //main.point{x:1, y:2}
    
    //需要打印值的类型，使用%T
    fmt.Printf("%T\n", p) //main.point

    //格式化布尔值
    fmt.Printf("%t\n", true)

    //格式化整型输出 %d标准十进制格式化
    fmt.Printf("%d\n", 123) // 123

    //14的二进制形式
    fmt.Printf("%b\n", 14) // 1110

    //整数转字符,65对应的字母是A
    fmt.Printf("%c\n", 65) //A
    
    //1024转换16进制
    fmt.Printf("%x\n", 1024) //400

    //%f基本的十进制浮点型
    fmt.Printf("%f\n", 3.1415926) //默认只能输出小数点后6位 3.141593
    fmt.Printf("%7f\n", 3.1415926) //结果显示这种形式与上面一样
    fmt.Printf("%0.2f\n", 3.1415926) //3.14
    fmt.Printf("%.2f\n", 3.1415926) 
    fmt.Printf("%.4f\n", 3.1415926) 
    
    //科学记数法
    fmt.Printf("%e\n", 1024000000.0)
    fmt.Printf("%E\n", 1024000000.0)

    //基本的字符串输出
    fmt.Printf("%s\n", "welcome come Golang word")

    //带有双引号的输出 %q
    fmt.Printf("%q\n", "\"string\"")

    //和整型数一样，%x输出使用base-16编码的字符串，每个字节使用2个字符表示
    fmt.Printf("%x\n", "1o8")
    
    //输出一个指针的值 %p
    fmt.Printf("%p\n", &p)

    //控制输出精度和宽度
    fmt.Printf("|%6d|%6d|\n", 123, 456)
    fmt.Printf("|%6.2f|%6.2f|\n", 3.14, 3.1415926)
    
    //使用-对齐
    fmt.Printf("|%-6.2f|%-6.2f|\n", 3.14, 3.1415926)
    
    //控制字符输出时的宽度，特别是要确保他们在类表格输出时的对齐，这是基本的右对齐宽度表示
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
    //左对齐
    fmt.Printf("|%-6s|%-6s|\n", "foo", "bobo")

    //Sprintf格式化并返回一个字符串而不带任何输出
    s := fmt.Sprintf("a % s", "string")
    fmt.Println(s)

    //Fprintf格式化并输出到io.Writers而不是os.Stdout
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
