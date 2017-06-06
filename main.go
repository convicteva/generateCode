/**
mybatis mapper 生成器
*/
package main

import (
//"golang/db"
)
import (
	"fmt"
	"golang/db"
	"golang/file"
	"golang/util"
)

var packageName string = "com.masz.demo"

func main() {
	createPackageTest()
}

func createPackageTest() {
	//生成相关目录
	dirInfo := file.CreatePackage(packageName)
	//生成BaseModel
	file.GenerateBaseModel(dirInfo.BaseModelPath, packageName)

	//TODO 生成BaseDao

	//表对应的字段map
	tableColumnMap := make(map[string][]db.Column)
	//所有的表
	tableNameSlice := db.GetTableName()
	for _, v := range tableNameSlice {
		tableColumnMap[v] = db.GetTableColumn(v)
	}

	//生成model
	for tabelName, columns := range tableColumnMap {
		file.GenerateMode(dirInfo.ModelPath, packageName, stringutil.FormatTableNameToModelName(tabelName), db.GetJavaProertyByColumn(columns))
	}

}

/**
测试获取表的信息
*/
func testTableInfo() {
	tableNameSlice := db.GetTableName()

	for _, v := range tableNameSlice {
		fmt.Println(stringutil.FormatTableNameToModelName(v))
		columnSlice := db.GetTableColumn(v)
		fmt.Println(columnSlice)
		for _, v := range columnSlice {
			fmt.Println(stringutil.FormatColumnNameToProperty(v.Name))
		}
	}
}
