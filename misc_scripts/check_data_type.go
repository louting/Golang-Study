package main

import (
	"fmt"
	"reflect"
)

func istype(i interface{}) {
	xtype := reflect.TypeOf(i)
	fmt.Println(xtype)
}

func main() {
	a := []int{1, 2, 3}
	istype(a)
}

