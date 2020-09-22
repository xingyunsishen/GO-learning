package main
import "fmt"

type Books struct{
    title string
    author string
    subject string
    book_id int
}

func main() {
    var Book1 Books     //声明Book1类型为Books
    var Book2 Books     //声明Book1类型为Books

    //book1描述
    Book1.title = "Go 语言"
    Book1.author = "www.golang.org"
    Book1.subject = "Go 语言教程"
    Book1.book_if = 9003020

    //book2描述
    Book2.title = "Python 语言"
    Book2.author = "www.python.org"
    Book2.subject = "Python 语言教程"
    Book2.book_id = "8003203
    
    //打印Book1信息
    printBook(&Book1)

    //打印Book2信息
    printBook(&Book2)
}

func prinBook( book *Books ) {
    fmt.Printf("Book title: %s\n", book.title)
    fmt.Printf("Book author: %s\n", book.author)
    fmt.Printf("Book subject: %s\n", book.subject)
    fmt.Printf("Book book_id: %s\n", book.book_id)
}
