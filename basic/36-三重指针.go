package main
import "fmt"

func main() {
    var a int = 5
    //把ptr指针指向a所在地址
    var ptr *int = &a
    //开辟一个新的指针，指向ptr指针指向的地方
    var pts *int = ptr
    //二级指针，指向一个地址，这个地址存储的是一级指针的地址
    var pto **int = &ptr
    //三级指针，指向一个地址，这个地址存储的是二级指针的地址，二级指针同上
    var pt3 ***int = &pto

    fmt.Println("a的地址:", &a, 
                "\n a = ", a, "\n\n",
                
                "ptr指针所在地址:", &ptr,
                "\n ptr指向的地址:", ptr,
                "\n ptr指针指向地址对应的值", *ptr, "\n\n",

                "pts指针所在地址:", &pts,
                "\n pts指向的地址:", pts,
                "\n pts指针指向的地址对应的值:", *pts, "\n\n",
                
                "pto指针所在地址:", &pto,
                "\n pto指向的指针(ptr)的存储地址:", pto,
                "\n pto指向的指针(ptr)所指向的地址:", *pto,
                "\n pto最终指向的地址对应的值(a)", **pto, "\n\n",
                
                "pt3指针所在的地址:", &pt3,
                "\n pt3指向的指针(pto)的地址:", pt3,
                "\n pt3指向的指针(pto)所指向的指针(ptr)地址:", *pt3,
                "\n pt3指向的指针(pto)所指向的指针(ptr)所指向的地址(a):", **pt3,
                "\n pt3最终指向的地址对应的值(a)", ***pt3)
}
