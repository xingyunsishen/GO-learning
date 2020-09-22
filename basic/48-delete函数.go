//delete()函数用于删除集合的元素，参数为map和其对应的key
package main

import "fmt"
func main() {
    //创建map
    countryCapitalMap := map[string]string{
        "China": "Beijing",
        "Russian": "Moscow",
        "Pakistan": "Islamabad",
        "Cuba": "La Habana",
        "Sweden": "Stockholm" }

    //打印地图
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [ country ])
    }
    
    //删除元素delete(countryCapitalMap, "Russian")
    fmt.Println("俄罗斯删除成功")
    fmt.Println("删除元素后地图")

    //打印地图
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [ country ])
    }
}
