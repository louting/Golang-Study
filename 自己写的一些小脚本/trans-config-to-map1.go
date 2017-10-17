package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func main() {
	var departmentInfo = make(map[string]map[string]string)
	var confile = "config.ini"
	config, _ := goconfig.LoadConfigFile(confile)
	departmentlist := config.GetSectionList()
	for _, v := range departmentlist {
		configcontent, _ := config.GetSection(v)
		var section = make(map[string]string) // 注意这个map的定义位置，放到for以外或者最内层的for 都是有问题的。
		for sectname, sectvalue := range configcontent {
			section[sectname] = sectvalue
			departmentInfo[v] = section
		}

	}
	fmt.Println(departmentInfo)

}
