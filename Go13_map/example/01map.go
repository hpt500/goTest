package example

import (
	. "fmt"
)

func MapStart() {
	Println("一.创建和使用map")
	// 注意:-1.如果不初始化 map,那么就会创建一个 nil map. nil map 不能用来存放键值对!!!
	// var countryMap map[string]string
	// 所以:-2.我们来给它创建集合!!!
	// countryMap = make(map[string]string)

	// 来个假设...
	countryMap := make(map[string]string)
	// -->假设成立!!!
	// 所以:使用内建make函数定义map的同时ye初始化了map!!!
	// 建议map关键字定义使用在全局,而函数体内强势建议使用内建make定义!!!

	// 强势插入国家键值对!!!
	countryMap["Franch"] = "Paris"
	countryMap["Italy"] = "Rome"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "new Delhi"

	for key, val := range countryMap { // 参考Go12_range用法
		Println("该国家为", key, "其首都为", val)
	}

	// 我们来查看该map集合中是否存在某个元素
	exam := "China"
	capital, ok := countryMap[exam]

	if ok {
		Println("Capital of", exam, "is", capital)
	} else {
		Println("Capital of", exam, "is not present!!!")
	}

	Println()

}
