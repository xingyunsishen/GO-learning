/*
Go通过range关键字来实现遍历读取到数据，类似于数组或切片。格式如下：
v, ok := <-ch
如果通道接收不到数据后ok就为false,这时通道就可以使用close()函数来关闭。
*/
package main
import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    //range函数遍历每个从通道接收到的数据，因为c在发送完10个数据之后就关闭了
    //通道，所以这里我们range函数在接收到10个数据之后就结束了。如果上面的c通
    //道不关闭，那么range函数就不会结束，从而在接收第11个数据的时候就阻塞了
    for i := range c {
        fmt.Println(i)
    }
}
