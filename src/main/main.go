package main

import (
	"fmt"
	"jvm-go/src/main/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	// 搜索 class 文件
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// 先打印命令行参数
	fmt.Printf("classpath: %v, class: %v, args: %v \n",
		cp, cmd.class, cmd.args)

	// 读取出类数据
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s \n", cmd.class)
		return
	}

	// 打印出来
	fmt.Printf("class data: %v \n", classData)
}
