/*
Go语言切片是对数组的抽象
Go数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强大的内置类型切片（“动态数组”），与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大

定义切片
var identifier []type；切片不需要说明长度
或使用make()函数来创建切片:
var slicel []type = make([]type, len)
也可以简写为
slicel := make([]type, len)
也可以指定容量，其中capacity为可选参数
make([]T, length, capacity)
这里len是数组的长度并且也是切片的初始长度

切片初始化
s := [] int {1, 2, 3}
直接初始化切片，[]表示切片类型，{1, 2, 3}初始化值依次是1, 2, 3，其中cap=len=3
s := arr[:]
初始化切片s,是数组arr的引用
s := arr[startIndex:endIndex]
将arr中从下标startIndex到endIndex-1下的元素创建为一个新的切片
s := arr[startIndex:]
默认endIndex时将表示一直到arr的最后一个元素
s := arr[:endIndex]
默认startIndex时将表示从arr的第一个元素开始
s1 := s[startIndex:endIndex]
通过切片s初始化切片s1
s := make([]int, len, cap)
通过内置函数make()初始化切片s,[]int标识为其元素类型为int的切片

len()和cap()函数
切片是可索引的，并且可以由len()方法获取长度。
切片提供了计算容量的方法cap()可以测量切片最长可以达到多少
*/

package main
import "fmt"

func main() {
    var numbers = make([]int, 3, 5)
    printSlice(numbers)
}

func printSlice(x []int) {
    fmt.Printf("len = %d cap=%d slice=%v\n", len(x), cap(x), x)
}
