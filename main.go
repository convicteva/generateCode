/**
mybatis mapper 生成器
*/
package main

import (
//"golang/db"
)
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
	columnAndJavaInfoMap := db.GetTableInfo(nil)

	//生成model
	createModel(dirInfo, columnAndJavaInfoMap)

	//生成mapper
	createMapper(dirInfo, columnAndJavaInfoMap)
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
