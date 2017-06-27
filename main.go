/**
mybatis mapper 生成器
*/
package main

import (
	"golang/config"
	"golang/db"
	"golang/file"
	"golang/util"
	"os"
	"runtime"
	"strings"
)

//输出代码目录
var root_path string = ""

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
}

func main() {
	//生成相关目录
	dirInfo := file.CreatePackage(root_path, config.Project_package_name)

	//表对应的字段  map
	tableColumnAndJavaInfoMap := db.GetTableInfo(nil)

	//生成model
	createModel(dirInfo, tableColumnAndJavaInfoMap)

	//生成mapper
	createMapper(dirInfo, tableColumnAndJavaInfoMap)

	//生成dao
	generateDao(dirInfo, tableColumnAndJavaInfoMap)

	//生成manager
	generateManager(dirInfo, tableColumnAndJavaInfoMap)

	//生成service
	generateService(dirInfo, tableColumnAndJavaInfoMap)
}

func createModel(dirInfo file.DirInfo, tableColumnAndJavaInfoMap map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	//生成BaseModel
	file.GenerateBaseModel(dirInfo.BaseModelPath, config.Project_package_name)
	//生成model
	for tabelName, columnAndJavaInfo := range tableColumnAndJavaInfoMap {
		file.GenerateMode(dirInfo.ModelPath, config.Project_package_name, stringutil.FormatTableNameToModelName(tabelName), columnAndJavaInfo)
	}
}

func createMapper(dirInfo file.DirInfo, tableColumnAndJavaInfoMap map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	//生成mapper
	for tabelName, columnAndJavaInfo := range tableColumnAndJavaInfoMap {
		file.GenerateMapper(dirInfo.MapperPath, config.Project_package_name, stringutil.FormatTableNameToModelName(tabelName), tabelName, columnAndJavaInfo)
	}
}

func generateDao(dirInfo file.DirInfo, tableColumnAndJavaInfoMap map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	//生成mapper
	for tabelName, _ := range tableColumnAndJavaInfoMap {
		file.GenerateDao(dirInfo.DaoPath, config.Project_package_name, stringutil.FormatTableNameToModelName(tabelName))
	}
}

func generateManager(dirInfo file.DirInfo, tableColumnAndJavaInfoMap map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	//生成mapper
	for tabelName, _ := range tableColumnAndJavaInfoMap {
		file.GenerateManager(dirInfo.ManagerPath, config.Project_package_name, stringutil.FormatTableNameToModelName(tabelName))
	}
}

func generateService(dirInfo file.DirInfo, tableColumnAndJavaInfoMap map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	//生成mapper
	for tabelName, _ := range tableColumnAndJavaInfoMap {
		file.GenerateService(dirInfo.ServicePath, config.Project_package_name, stringutil.FormatTableNameToModelName(tabelName))
	}
}
