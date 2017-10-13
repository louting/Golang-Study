package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func main() {
	var configFile = make(map[string]string)
	var configname = "config.ini"
	config, _ := goconfig.LoadConfigFile(configname)
	configcontent, _ := config.GetSection("topicArr")
	for sectname, sectvalue := range configcontent {
		configFile[sectname] = sectvalue
	}
	fmt.Println(configFile)
}

