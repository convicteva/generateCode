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

	//生成BaseModel
	file.GenerateBaseModel(dirInfo.BaseModelPath, config.Project_package_name)

	for tableName, columnAndJavaInfo := range tableColumnAndJavaInfoMap {
		modelName := stringutil.FormatTableNameToModelName(tableName)
		//生成model
		file.GenerateMode(dirInfo.ModelPath, config.Project_package_name, modelName, columnAndJavaInfo)

		//生成dao
		file.GenerateDao(dirInfo.DaoPath, config.Project_package_name, modelName)

		//生成mapper
		file.GenerateMapper(dirInfo.MapperPath, config.Project_package_name, modelName, tableName, columnAndJavaInfo)

		//生成manager
		file.GenerateManager(dirInfo.ManagerPath, config.Project_package_name, modelName)

		//生成service
		file.GenerateService(dirInfo.ServicePath, config.Project_package_name, modelName)
	}
}
