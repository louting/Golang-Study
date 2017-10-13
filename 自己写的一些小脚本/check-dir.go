package main

import (
	"bytes"
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	pwd, _ := os.Getwd()

	filename := "/test2.go"
	//file := strings.Join(pwd, filename)
	b := bytes.Buffer{}
	b.WriteString(pwd)
	b.WriteString(filename)
	s := b.String()
	fmt.Println(s)
	a, _ := PathExists(s)
	fmt.Println(a)
}
