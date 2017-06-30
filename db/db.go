/**
数据库配置及访问
*/
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang/config"
	"strings"
)

var db *sql.DB = nil

func init() {
	sqlUrl := config.USERNAME + ":" + config.PASSWD + "@tcp(" + config.IP + ")/" + config.DATABASENAME + "?charset=utf8"
	db, _ = sql.Open("mysql", sqlUrl)
	if db != nil {
		db.SetMaxOpenConns(config.MAXOPENCONNS)
		db.SetMaxIdleConns(config.MAXIDLECONNS)
		db.Ping()
	} else {
		panic("db open fail")
	}
}

/**
查询所有的表
*/
func getTableName() []string {
	if db == nil {
		panic("db is nil")
	}
	sqlStr := "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '" + config.DATABASENAME + "'"
	rows, err := db.Query(sqlStr)
	defer rows.Close()

	if err == nil {
		var nameSlice = make([]string, 0, 5)
		var tableName string = ""
		for rows.Next() {
			rows.Scan(&tableName)
			nameSlice = append(nameSlice, tableName)
		}
		return nameSlice
	} else {
		fmt.Println("GetAllTable execute fail,", err.Error())
		panic(err)
	}
	return nil
}

/**
查询表的所有字段
*/
func getTableColumn(tableName string) []Column {
	sqlStr := "SELECT column_name,column_comment,data_type FROM information_schema.COLUMNS WHERE table_name='" + tableName + "' AND table_schema = '" + config.DATABASENAME + "'"
	rows, err := db.Query(sqlStr)
	columnSlice := make([]Column, 0, 10)
	if err == nil {
		var name string
		var comment string
		var dataType string
		for rows.Next() {
			rows.Scan(&name, &comment, &dataType)
			columnSlice = append(columnSlice, Column{strings.ToUpper(name), comment, dataType})
		}
		return columnSlice
	}
	return columnSlice
}

/**
生成表对应的字段，返回以表名为key，字段人间信息SqlColumnAndJavaPropertiesInfo 切片的map
*/
func GetTableInfo(tableNameSlice []string) map[string][]SqlColumnAndJavaPropertiesInfo {

	//返回值，以table name 为key
	columnAndJavaInfo := make(map[string][]SqlColumnAndJavaPropertiesInfo)

	if tableNameSlice == nil {
		tableNameSlice = getTableName()
	}

	tableColumnMap := make(map[string][]Column)
	for _, v := range tableNameSlice {
		tableColumnMap[v] = getTableColumn(v)
	}
	for key, columnSlice := range tableColumnMap {
		columnAndJavaInfo[key] = ColumnInfo2JavaInfo(columnSlice)
	}
	return columnAndJavaInfo
}
