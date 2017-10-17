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
	"github.com/Unknwon/goconfig"
	"github.com/bitly/go-simplejson"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func showUsage(output io.Writer) {
	fmt.Fprintf(output, "USAGE: Exec QYwechatMessageSent OPTIONS\n")
	fmt.Fprintf(output, "\n")
	fmt.Fprintf(output, "OPTIONS:\n")
	fmt.Fprintf(output, "  -d     --department          Must specify department,like Operations or Develop....\n")
	fmt.Fprintf(output, "  -C     --content             Must specify what you want sent.\n")
	fmt.Fprintf(output, "  -h     --help                Show this usage.\n")
}

func die(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
	os.Exit(1)
}

type MessageSent struct {
	corpid, agentsecret, conn1, conn2 string
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

//取得配置文件中的账号和密码
func Getconfig() ([]string, map[string]map[string]string) {
	var departmentInfo = make(map[string]map[string]string)
	var confile = "config.ini"
	config, _ := goconfig.LoadConfigFile(confile)
	departmentlist := config.GetSectionList()

	for _, v := range departmentlist {
		configcontent, _ := config.GetSection(v)
		var section = make(map[string]string)
		for sectname, sectvalue := range configcontent {

			section[sectname] = sectvalue
			departmentInfo[v] = section
		}
	}
	return departmentlist, departmentInfo
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

	// ####参数检查设定
	params := os.Args[1:]
	var (
		Contentbody string
	)

	if len(params) == 0 {
		showUsage(os.Stdout)
		os.Exit(0)
	}

	for i, arg := range params {

		switch arg {
		case "-h", "--help":
			showUsage(os.Stdout)
			return
		case "-C", "--content":
			if i+1 >= len(params) {
				die("-C no content specified")
			}
			Contentbody = params[i+1]
		}
	}
	// ####参数检查设定完成

	department, info := Getconfig()
	fmt.Println(department)
	fmt.Println(info)
	for _, value := range department {
		fmt.Println(value)
		for k, v := range info {
			if k == value {
				M := MessageSent{
					corpid:      "wx23a411e326ed7008", // 我这里写死了，这个每个企业是唯一的。
					agentsecret: v["Agentsecret"],
					conn1:       "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=",
					conn2:       "&corpsecret=",
					department:  value,
				}

				Token := M.Access_token()

				Url := `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=`

				Te := Text{Content: Contentbody} // 因为json转换的原因

				C := Content{Touser: "@all", Msgtype: "text", Agentid: v["AgentId"], Text: Te, Safe: "0"}

				messagebody, err := json.Marshal(C)
				if err != nil {
					panic(err.Error())
				}

				M.Sentdata(Token, Url, string(messagebody))

			}

		}

	}

}
