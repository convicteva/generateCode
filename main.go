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
	"golang/util"
)

func main() {
	testTableInfo()
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
