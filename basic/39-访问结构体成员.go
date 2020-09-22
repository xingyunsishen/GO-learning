package main
import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}   

func main() {
    var Book1 Books //声明Book1为Books类型
    var Book2 Books //声明Book2为Books类型
    //book1描述
    Book1.title = "Go 语言"
    Book1.author = "www.golang.org"
    Book1.subject = "Go 语言编程"
    Book1.book_id = 23914234

    //book2描述
    Book2.title = "Python 教程"
    Book2.author = "www.python.org"
    Book2.subject = "Python 语言编程"
    Book2.book_id = 23923223

    //打印Book1信息
    fmt.Printf("Book 1 title : %s\n", Book1.title)
    fmt.Printf("Book 1 author : %s\n", Book1.author)
    fmt.Printf("Book 1 subject : %s\n", Book1.title)
    fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

    //打印Book2信息
    fmt.Printf("Book 2 title : %s\n", Book2.title)
    fmt.Printf("Book 2 author : %s\n", Book2.author)
    fmt.Printf("Book 2 subject : %s\n", Book2.title)
    fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}
