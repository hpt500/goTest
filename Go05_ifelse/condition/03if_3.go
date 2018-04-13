package condition

import (
	. "fmt"
)

func init() {
	Println("三、寻找到100以内的所有的素数")
	var count int
	var flag bool
	count = 1

	for count < 100 {
		count++
		flag = true
		for temp := 2; temp < count; temp++ {
			if count%temp == 0 {
				flag = false
			}
		}
		if flag == true {
			Println(count, "是素数")
		} else {
			continue
		}
	}

	Println()
}
