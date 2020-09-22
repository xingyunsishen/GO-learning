package main
import "fmt"

func main() {
    //Go语言中range关键字用于for循环中迭代数组(array)、切片(slice)、
    //通道(channe)或者集合(map)的元素，在数组和切片中它返回元素的索引和
    //索引对应的值，在集合中返回key-value对
    //这是我们使用range去求一个slice的和，使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    //在数组上使用range将传入index和值两个变量，上面那个例子我们不需要使用该
    //素的序号，所以我们使用了空白符"_" 省略了。有时候我们确实需要知道的索引

    for i, num := range nums {
        if num == 3 {
            fmt.Println("idnex:", i)
        }
    }
    //range也可以用在map的键值对上，
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    //range也可以用来枚举Unicode字符串，第一个参数是字符的索引，第二个字符Unicode的值本身
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
