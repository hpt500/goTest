package example

import (
	. "fmt"
)

func DeleFunc() {
	Println("二.delete函数的使用")
	// delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key
	countryMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	Println("-1.delete函数删除前的map集合:")
	for key, val := range countryMap {
		Println("该国家为", key, "其首都为", val)
	}

	// delete函数删除一项"France"
	delete(countryMap, "France")

	Println("-2.delete函数删除后的map集合:")
	for key, val := range countryMap {
		Println("该国家为", key, "其首都为", val)
	}

	Println()

}
