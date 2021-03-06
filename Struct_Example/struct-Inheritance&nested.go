// 结构的继承

package main

import (
	"fmt"
)

type Content1 struct {
	safe int
}
type Content2 struct {
	Name string
}

type test1 struct {
	Content1
	Content2
}

func (t *Content1) Get(a int) {
	t.safe = a
	fmt.Println(t.safe)
}

func (t *Content2) Get2(a string) {
	t.Name = a
	fmt.Println(t.Name)
}

func main() {

	// 只要定义一个结构，其继承了被嵌套对象的所有方法。
	type mystruct struct {
		test1
	}

	a := mystruct{}
	a.Get(100)
	a.Get2("test")
}
