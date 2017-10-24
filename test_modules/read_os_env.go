package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Getenv("PATH")
	//fmt.Println(path)
	for key, v := range path {
		fmt.Println(key, v)
	}
}
