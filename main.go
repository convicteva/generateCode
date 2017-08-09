/**
mybatis mapper 生成器
*/
package main

import (
	"genereateCode/config"
	"genereateCode/db"
	"genereateCode/file"
	"os"
	"runtime"
	"strings"
)

//输出代码目录
var root_path string = ""

var dirInfo file.DirInfo

func main() {
	//将从web 传入
	var packageName = "com.masz.demo"
	var node = "test"
	var tableNameSlice []string = nil

	generate(packageName, node, tableNameSlice)
}

func generate(packageName, node string, tableNameSlice []string) {
	initBaseInfo(packageName, node)
	file.Generate(dirInfo, tableNameSlice)
}

func initBaseInfo(packageName, node string) {
	//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	var system = strings.ToUpper(runtime.GOOS)
	//根据操作系统使用不同的默认目录。留以后导出使用
	if strings.EqualFold(system, "WINDOWS") {
		root_path = "C:\\temp\\generateCode\\"
	} else if strings.EqualFold(system, "LINUX") {
		root_path = "/tmp/generateCode/"
	}
	//删除之前存在的内容
	os.RemoveAll(root_path)

	config.Project_package_name = packageName
	//生成相关目录
	dirInfo = file.CreatePackage(root_path, config.Project_package_name)
	db.InitDB(node)

}
