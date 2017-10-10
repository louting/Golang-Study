／*
go 动态类型的小白表述

定义：实现了某个接口的类型可以被传给任何以此接口为参数的函数。

小白理解何为动态：如果理解这个动态，就要先说静态，静态就是规定了某个接口类型“只能”被某个函数调用。而动态则是借口可以被任意定义了其为参数的函数调用。

go的动态类型使用interface实现。



实现方法：
    1 创建一个接口“x”，接口包含方法一，方法二。 例如：接口 x {方法1，方法2}

    2 创建一个函数，或者N个函数，以接口为参数。例如： 函数1（x）{方法1，方法2}

    3 创建一个数据。例如：data。 <--- 这个数据类型，就是实现了接口的动态类型。 问：是不是只要实现了某个接口的类型，就是动态类型？

    4 创建数据方法，要满足x接口。例如： func(d data) 方法1（）{}.    func(d data) 方法2（）{}.

*／


package main

import "fmt"

type IDuck interface {
    Quack()
    Walk()
}

func DuckDance(duck IDuck) {
    for i := 1; i <= 3; i++ {
        duck.Quack()
        duck.Walk()
    }
}

type Bird struct {
    // ...
}

func (b *Bird) Quack() {
    fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
    fmt.Println("I am walking!")
}

func main() {
    b := new(Bird)
    DuckDance(b)
}
