package main
import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}

func main() {
    var Book1 Books  //声明Book1 为Books类型
    var Book2 Books  //声明Book2 为Books类型

    //book1 描述
    Book1.title = "Go 语言"
    Book1.author = "www.golang.com"
    Book1.subject = "Go 语言教程"
    Book1.book_id = 9822312

    //book2 描述
    Book2.title = "Python 语言"
    Book2.author = "www.python.org"
    Book2.subject = "Python 语言教程"
    Book2.book_id = 9822322

    //打印Book1信息
    printBook(Book1)

    //打印Book2信息
    printBook(Book2)
}

func printBook( book Books ) {
    fmt.Printf("Book title: %s\n", book.title)
    fmt.Printf("Book author: %s\n", book.author)
    fmt.Printf("Book subject: %s\n", book.subject)
    fmt.Printf("Book book_id: %d\n", book.book_id)
}
