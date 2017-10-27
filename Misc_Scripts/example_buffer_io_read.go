package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"reflect"
)

func istype(i interface{}) {
	xtype := reflect.TypeOf(i)
	fmt.Println(xtype) // result: *os.file
}

func main() {
	//读取这个脚本自身
	cmd := exec.Command("cat", "get_os_env.go") //func Command(name string, arg ...string) *Cmd

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe() // cmd.StdoutPipe() 返回 io.ReadCloser 接口
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	istype(stdout)

	//执行命令
	if err := cmd.Start(); err != nil { //***如果是cmd.Run() 程序运行会报错。
		fmt.Println("Error:The command is err,", err)
		return
	}

	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout) // func NewReader(rd io.Reader) *Reader

	for {

		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine() // func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		if err != nil {

			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err) // *** errmessage: Error :read |0: file already closed
			}
			return
		}
		fmt.Printf("%s\n", string(output))
	}

	//wait 方法会一直阻塞到其所属的命令完全运行结束为
	if err := cmd.Wait(); err != nil { //Wait会阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的。
		fmt.Println("wait:", err.Error())
		return
	}
}

