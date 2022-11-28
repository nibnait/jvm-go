package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// 使用 -Xjre 选项解析启动类路径和扩展类路径
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// get jre目录
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		// 优先使用 用户输入的 -Xjre 选项做为jre目录
		return jreOption
	}
	if exists("./jre") {
		// 在当前目录下，寻找jre目录
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		// 尝试使用 JAVA_HOME 环境变量
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 使用 -cp 选项解析用户类路径
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	// 使用 -Xjre 选项解析启动类路径和扩展类路径
	cp.parseBootAndExtClasspath(jreOption)
	// 使用 -cp 选项解析用户类路径
	cp.parseUserClasspath(cpOption)
	return cp
}

// 如果用户没有提供 -cp 选项，则使用当前目录作为用户类路径
// ReadClass() 方法一次从启动类路径、扩展类路径、和用户类路径中搜索 class 文件
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

// 返回用户类路径的字符串
func (self *Classpath) String() string {
	return self.userClasspath.String()
}
