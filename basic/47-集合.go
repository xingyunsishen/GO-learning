/*
Map是一种无需的键值对的集合。Map最重要的一点是通过key快速加检索数据，key
类似与索引，只想数据的值。
Map是一种集合，所以我们可以向迭代数组和切片那样迭代它。不过，Map是无序的，我们无法决定它的返回顺序，这是因为Map是使用hash表来实现的。
定义Map
可以使用内奸函数make也可以使用map关键字来定义Map：
//声明变量，默认map是nil
var map_variable map[key_data_type]value_data_type

//使用make函数
map_variable := make(map[key_data_type]value_data_type)
*/
package main
import "fmt"

func main() {
    var countryCapitalMap map[string]string //创建集合
    countryCapitalMap = make(map[string]string)

    //map插入key - value对，各个国家对应的首都
    countryCapitalMap ["China"] = "北京"
    countryCapitalMap ["Russian"] = "莫斯科"
    countryCapitalMap ["Pakistan"] = "伊斯兰堡"
    countryCapitalMap ["Sweden"] = "斯德哥尔摩"

    //使用键输出地图值
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

    //查看元素在集合中是否存在
    capital, ok := countryCapitalMap ["American"] //如果确定是真实的，则存在，否则不存在
    //fmt.Println(capital)
    //fmt.Println(ok)
    if (ok) {
        fmt.Println("American 的首都是", capital)
    } else {
        fmt.Println("American 的首都被炸飞了，不在了")
    }
}
