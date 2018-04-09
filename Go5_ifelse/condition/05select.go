package condition

import (
	. "fmt"
)

func init() {
	Println("五、select语句用法")
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			Printf("received ", i3, " from c3\n")
		} else {
			Printf("c3 is closed\n")
		}
	default:
		Printf("no communication\n")
	}
	Println()
}
