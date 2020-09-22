/*
GO语言提供了数组类型的数据结构
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型如整型、字符串或自定义类型。
相对于去声明number0,number1,...,number99的变量，使用数组形式numbers[0],numbers[1],...,numbers[99]更加方便且易于扩展。
数组元素可以通过索引(位置)来读取(或者修改)，索引0开始，第一个元素索引为0，第二个元素索引为1，以此类推
*/
package main
import "fmt"

func main() {
    var n [10]int //n是一个长度为10的数组
    var i, j int

    //初始化数组n中的元素
    for i = 0; i < 10; i++ {
        n[i] = i + 100 //设置元素为i + 100
    }
    //输出每个数组元素的值
    for j = 0; j < 10; j++ {
        fmt.Printf("Element[%d] = %d\n", j, n[j])
    }
}
