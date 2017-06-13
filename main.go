/**
mybatis mapper 生成器
*/
package main

import (
	"golang/db"
	"golang/file"
	"golang/util"
)

var packageName string = "com.masz.demo"

func main() {
	//生成相关目录
	dirInfo := file.CreatePackage(packageName)

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
	file.GenerateBaseModel(dirInfo.BaseModelPath, packageName)
	//生成model
	for tabelName, columnAndJavaInfo := range tableColumnAndJavaInfoMap {
		file.GenerateMode(dirInfo.ModelPath, packageName, stringutil.FormatTableNameToModelName(tabelName), columnAndJavaInfo)
	}
}

func createMapper(dirInfo file.DirInfo, tableColumnAndJavaInfoMap map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	//生成mapper
	for tabelName, columnAndJavaInfo := range tableColumnAndJavaInfoMap {
		file.GenerateMapper(dirInfo.MapperPath, packageName, stringutil.FormatTableNameToModelName(tabelName), tabelName, columnAndJavaInfo)
	}
}

func generateDao(dirInfo file.DirInfo, tableColumnAndJavaInfoMap map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	//生成mapper
	for tabelName, _ := range tableColumnAndJavaInfoMap {
		file.GenerateDao(dirInfo.DaoPath, packageName, stringutil.FormatTableNameToModelName(tabelName))
	}
}

func generateManager(dirInfo file.DirInfo, tableColumnAndJavaInfoMap map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	//生成mapper
	for tabelName, _ := range tableColumnAndJavaInfoMap {
		file.GenerateManager(dirInfo.ManagerPath, packageName, stringutil.FormatTableNameToModelName(tabelName))
	}
}

func generateService(dirInfo file.DirInfo, tableColumnAndJavaInfoMap map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	//生成mapper
	for tabelName, _ := range tableColumnAndJavaInfoMap {
		file.GenerateService(dirInfo.ServicePath, packageName, stringutil.FormatTableNameToModelName(tabelName))
	}
}
