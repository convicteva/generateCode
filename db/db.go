/**
数据库配置及访问
*/
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ip           string = "192.168.157.128:3306"
	databaseName string = "cs"
	username     string = "root"
	passwd              = "Abc123!@#"
	maxOpenConns int    = 5
	maxIdleConns int    = 2
)

var db *sql.DB = nil

func init() {
	sqlUrl := username + ":" + passwd + "@tcp(" + ip + ")/" + databaseName + "?charset=utf8"
	fmt.Println(sqlUrl)
	db, _ = sql.Open("mysql", sqlUrl)
	if db != nil {
		db.SetMaxOpenConns(maxOpenConns)
		db.SetMaxIdleConns(maxIdleConns)
		db.Ping()
	} else {
		panic("db open fail")
	}
}

/**
查询所有的表
*/
func GetTableName() []string {
	if db == nil {
		panic("db is nil")
	}
	sqlStr := "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '" + databaseName + "'"
	rows, err := db.Query(sqlStr)
	defer rows.Close()
	
	if err == nil {
		var nameSlice = make([]string,2,5)
		var tableName string = ""
		for rows.Next(){
			rows.Scan(&tableName)
			nameSlice = append(nameSlice,tableName)
		}
		return nameSlice
	} else {
		fmt.Println("GetAllTable execute fail,", err.Error())
		panic(err)
	}
	return nil
}
