package example

import (
	. "fmt"
)

func Range() {
	Println("一.使用range求slice的和")
	nums := []int{1, 2, 3, 4, 5, 6}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	Println("切片nums的总和为:", sum)
	Println()
	// 在数组上使用range将传入index和值两个变量
	// 使用range循环数组切片时,可通过第一参数获取索引,若不需要我们可以使用空白符"_"省略
	Println("二.使用range获取数组切片的索引值")
	for i, num := range nums {
		if num == 3 {
			Printf("值为%d的索引为%d\n", num, i)
		}
	}
	Println()

	// range也可使用在map的键值对上 k,v := range nums  k值对应键 v值对应键值
	Println("三.range也可使用在map的键值对上")
	kvs := map[string]string{"a": "a的值", "b": "b的值"}
	for k, v := range kvs {
		Printf("%s->%s\n", k, v)
	}
	Println()

	// range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	Println("四.range也可以用来枚举Unicode字符串")
	for i, c := range "go" {
		Printf("'go'字符串索引%d的Unicode值为%d\n", i, c)
	}

}
