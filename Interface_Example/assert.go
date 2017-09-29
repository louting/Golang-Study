// 断言判断接口类型

package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf ?
	if t, ok := areaIntf.(*Square); ok { // 如果这个*号 和下面56行的* 都去掉，那么结果就是 areaIntf does not contain a variable of type Circle
		fmt.Printf("The type of areaIntf is: %T\n", t) // %T 是打印类型
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
	// testing with switch:
	switch t := areaIntf.(type) { //type 关键字可以这么用。。。。 不好理解啊
	case *Square: // 注意判断条件的写法。%T 打印出来是 *main.Square
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	/*
				case bool:
		   			fmt.Printf("Type boolean %t\n", t)
				case int:
		   			fmt.Printf("Type int %d\n", t)
				case *bool:
		   			fmt.Printf("Type pointer to boolean %t\n", *t)
				case *int:
		   			fmt.Printf("Type pointer to int %d\n", *t)
	*/
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

/*  result

The type of areaIntf is: *main.Square
areaIntf does not contain a variable of type Circle
Type Square *main.Square with value &{5}
*/
