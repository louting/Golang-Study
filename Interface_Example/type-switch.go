package main

import (
	"fmt"
)

func classifier(items ...interface{}) {
	for i, x := range items { //i 索引， x 值
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}

func main() {

	classifier(100, "int")

}

/* the result is
Param #0 is a int
Param #1 is a string
*/

//类型判断是比类型断言更全面的， 类型‘断言’用于简单可预计情况判断是否符合接口要求的值
