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

func main() {
	createPackageTest()
}

func createPackageTest() {
	packageName := "com.masz.demo"
	dirInfo := file.CreatePackage(packageName)
	file.GenerateBaseModel(dirInfo.BaseModelPath, packageName)
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
