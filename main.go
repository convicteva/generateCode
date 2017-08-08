/**
mybatis mapper 生成器
*/
package main

import (
	"golang/config"
	"golang/db"
	"golang/file"
	"os"
	"runtime"
	"strings"
)

//输出代码目录
var root_path string = ""

var dirInfo file.DirInfo

func init() {
	//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	var system = strings.ToUpper(runtime.GOOS)
	//根据操作系统使用不同的默认目录。留以后导出使用
	if strings.EqualFold(system, "WINDOWS") {
		root_path = "C:\\temp\\"
	} else if strings.EqualFold(system, "LINUX") {
		root_path = "/tmp/"
	}
	//删除之前存在的内容
	os.Remove(root_path)

	//生成相关目录
	dirInfo = file.CreatePackage(root_path, config.Project_package_name)
}

func main() {

	db.InitDB("test")
	file.Generate(dirInfo, nil)

}
