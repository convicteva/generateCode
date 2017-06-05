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
)

func main() {
	tableNameSlice := db.GetTableName()
	fmt.Println(len(tableNameSlice))
	for _, v := range tableNameSlice {
		fmt.Println("table name " + v)
		columnSlice := db.GetTableColumn(v)
		fmt.Println(columnSlice)
	}
}
