// 接口的定义和使用方法

package main

import (
	"fmt"
)

type Simplyer interface {
	Get() int
	Set(int)
}

type Simple struct {
	aa int
}

func (s *Simple) Get() int {
	return s.aa
}

func (s *Simple) Set(b int) {
	s.aa = b
}

func main() {
	//S := Simple{}  这种表达式是错误的，如果是赋值例如：Simple{1,2} 类似这样，这个表达式就没问题。
	S := new(Simple)
	//Sr := Simplyer{S} 这种定义的方式也是错误的。
	var Sr Simplyer
	Sr = S
	Sr.Set(100)
	fmt.Println(Sr.Get())
}
