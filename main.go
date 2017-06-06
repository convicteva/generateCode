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
	tableColumnMap := make(map[string][]db.Column)
	//所有的表
	tableNameSlice := db.GetTableName()
	for _, v := range tableNameSlice {
		tableColumnMap[v] = db.GetTableColumn(v)
	}

	//生成model
	createModel(dirInfo, tableColumnMap)

	//生成mapper
	createMapper(dirInfo, tableColumnMap)
}

func createModel(dirInfo file.DirInfo, tableColumnMap map[string][]db.Column) {

	//生成BaseModel
	file.GenerateBaseModel(dirInfo.BaseModelPath, packageName)

	//生成model
	for tabelName, columns := range tableColumnMap {
		file.GenerateMode(dirInfo.ModelPath, packageName, stringutil.FormatTableNameToModelName(tabelName), db.GetJavaProertyByColumn(columns))
	}

}

func createMapper(dirInfo file.DirInfo, tableColumnMap map[string][]db.Column) {

}
