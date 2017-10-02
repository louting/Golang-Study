package main

import (
	"./sort"
	"fmt"
)

type name struct {
	num      int
	fistname string
	lastname string
}

type nameArray struct {
	data []*name
}

func (n *nameArray) Len() int { return len(n.data) }

func (n *nameArray) Less(i, j int) bool { return n.data[i].num < n.data[j].num }

func (n *nameArray) Swap(i, j int) { n.data[i], n.data[j] = n.data[j], n.data[i] }

func names() {
	// 根据lastname 排序
	Jerry := name{1, "Jerry", "Lin"}
	Tom := name{3, "Tom", "Wang"}
	Jone := name{2, "Jone", "Snow"}
	data := []*name{&Jerry, &Tom, &Jone}
	a := nameArray(data)
	sort.Sort(&a)
	for _, d := range data {
		fmt.Printf("%s %s", d.lastname, d.fistname)
	}
	fmt.Printf("\n")

}

func main() {
	names()
}
