package main

/*
author:tonylou
QQ:34530529
Date:2017-09-18
Function instruction:发送企业微信脚本
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type AuthenStruct struct {
	DepartmentAu1 DepartmentAu
	DepartmentAu2 DepartmentAu
}

type DepartmentAu struct {
	AgentId string
	Secret  string
}

func showUsage(output io.Writer) {
	fmt.Fprintf(output, "USAGE: Exec QYwechatMessageSent OPTIONS\n")
	fmt.Fprintf(output, "\n")
	fmt.Fprintf(output, "OPTIONS:\n")
	fmt.Fprintf(output, "  -d     --department          Must specify department,like Operations or Develop....\n")
	fmt.Fprintf(output, "  -C     --contant             Must specify what you want sent.\n")
	fmt.Fprintf(output, "  -h     --help                Show this usage.\n")
}

func die(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
	os.Exit(1)
}

type MessageSent struct {
	corpid, agentsecret, conn1, conn2 string
	contant                           string
	department                        string
}

// 发送数据内容
type Text struct {
	Content string `json:"content"`
}

type Content struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Agentid string `json:"agentid"`
	Text    Text   `json:"text"`
	Safe    string `json:"safe"`
}

// 返回Access_token 这个每7200秒更新
func (m *MessageSent) Access_token() string {
	var buffer bytes.Buffer
	buffer.WriteString(m.conn1)
	buffer.WriteString(m.corpid)
	buffer.WriteString(m.conn2)
	buffer.WriteString(m.agentsecret)
	token := buffer.String()
	resp, err := http.Get(token)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	js, err := simplejson.NewJson(body)
	access_token := js.Get("access_token").MustString()
	return access_token
}

func (m *MessageSent) Sentdata(access_token, Url, messagebody string) {
	var buffer1 bytes.Buffer
	str1 := strings.NewReader(messagebody)
	body_type := "application/x-www-form-urlencoded"
	buffer1.WriteString(Url)
	buffer1.WriteString(access_token)
	sendurl := buffer1.String()
	resp, _ := http.Post(sendurl, body_type, str1)
	return_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(return_body))

}

func main() {
	params := os.Args[1:]
	var (
		Department  string
		Contentbody string
		Agentid     string
		Agentsecret string
	)

	if len(params) == 0 {
		showUsage(os.Stdout)
		os.Exit(0)
	}
	//设置Department 默认值
	Department = "Operations"

	//参数设定
	for i, arg := range params {

		switch arg {
		case "-d", "--department":
			Department = params[i+1]
		case "-h", "--help":
			showUsage(os.Stdout)
			return
		case "-C", "--content":
			if i+1 >= len(params) {
				die("-C no contant specified")
			}
			Contentbody = params[i+1]
		}
	}
	//创建部门代码和密码的map hash
	Authen := make(map[string]map[string]string)

	DepartmentAu1 := make(map[string]string)
	DepartmentAu1["AgentId"] = "1000003"
	DepartmentAu1["Secret"] = "x6c1p3suYcK5hTqXPyk07yeIbunDQOltp66mCCVSLm0"
	Authen["Operations"] = DepartmentAu1

	DepartmentAu2 := make(map[string]string)
	DepartmentAu2["AgentId"] = "1000002"
	DepartmentAu2["Secret"] = "qxC612OYEsKn-N1V5FuXDvQ0uxfd-DEPhavgtouxVls"
	Authen["Develop"] = DepartmentAu2

	for k, v := range Authen[Department] {
		if k == "AgentId" {
			Agentid = v

		}
		if k == "Secret" {
			Agentsecret = v

		}
	}

	M := MessageSent{
		corpid:      "wx23a411e326ed7008",
		agentsecret: Agentsecret,
		conn1:       "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=",
		conn2:       "&corpsecret=",
		contant:     "test",
		department:  "test1",
	}

	T := M.Access_token()

	Url := `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`

	Te := Text{Content: Contentbody}

	C := Content{Touser: "@all", Msgtype: "text", Agentid: Agentid, Text: Te, Safe: "0"}

	messagebody, err := json.Marshal(C)
	if err != nil {
		panic(err.Error())
	}

	M.Sentdata(T, Url, string(messagebody))
}
