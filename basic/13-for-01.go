package main
import "fmt"

func main() {
    sum := 1
    for; sum <= 10; {
        sum += sum
    }
    fmt.Println(sum)
    //这样写也可以，更像while语句
    for sum <= 10{
        sum += sum
    }
    fmt.Println(sum)
}
