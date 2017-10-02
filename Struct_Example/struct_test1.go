package main

import (
	"fmt"
)

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day // 取值 day struct 类型的切片的值
}

func main() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday} //struct 切片
	a := dayArray{data}

	fmt.Println(a)
}
