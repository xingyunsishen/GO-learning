package main
import "fmt"

func main() {
    //一个切片在未开始初始化前默认未nil，长度为0
    var numbers []int
    printSlice(numbers)

    if(numbers == nil){
        fmt.Printf("切片是空的")
    }
}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d slice=%d\n", len(x), cap(x), x)
}
