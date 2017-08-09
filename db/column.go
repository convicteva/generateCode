/**
数据库字段相关信息
*/
package db

import (
	"genereateCode/config"
	"genereateCode/util"
	"strings"
)

/**
表字段定义
*/
type column struct {
	//字段名
	Name string
	//注释
	Comment string
	//类型
	DataType string
}

/**
数据库列信息，包含：列名，注释，列数据类型，列对应的java属性名称，java类型，mybatis 数据类型
*/
type SqlColumnAndJavaPropertiesInfo struct {
	//字段
	column

	//java 属性名称
	JavaPropertyName string

	//java 类型
	JavaType string

	//mybatis jdbc 数据类型
	JdbcType string
}

/**
表对应的字段信息(包含字段对应的java信息)
*/
type TableColumnAndJavaInfo struct {
	TableName  string
	ColumnInfo []SqlColumnAndJavaPropertiesInfo
}

/**
根据数据库数据获取 JdbcJavaTypeMap
*/
func getJdbcJavaTypeMapBySqlType(sqlDataType string) JdbcJavaTypeMap {
	v, exists := MysqlTypeToJava[strings.ToUpper(sqlDataType)]
	if exists {
		return v
	} else {
		panic(sqlDataType + ",mysql 数据类型映射未找到")
	}
}

/**
表的列，转化成列信息，和java属性信息
返回 SqlColumnAndJavaPropertiesInfo 切片
*/
func ColumnInfo2JavaInfo(columns []column) []SqlColumnAndJavaPropertiesInfo {
	columnAndJavaInfoSlice := make([]SqlColumnAndJavaPropertiesInfo, 0, len(columns))
	for _, c := range columns {
		jdbcJavaTypeMap := getJdbcJavaTypeMapBySqlType(c.DataType)

		info := SqlColumnAndJavaPropertiesInfo{JavaPropertyName: util.FormatColumnNameToProperty(c.Name),
			JavaType: jdbcJavaTypeMap.JavaType,
			JdbcType: jdbcJavaTypeMap.JdbcType}

		info.Name = c.Name
		info.DataType = c.DataType
		info.Comment = c.Comment
		columnAndJavaInfoSlice = append(columnAndJavaInfoSlice, info)
	}
	return columnAndJavaInfoSlice
}

/**
表的列，转化成列信息，和java属性信息
返回 SqlColumnAndJavaPropertiesInfo 切片
*/
func ConfigColumn2JavaInfo(configColumns []config.ConfigColumn) []SqlColumnAndJavaPropertiesInfo {
	columnSlice := make([]column, 0, len(configColumns))
	for _, c := range configColumns {
		columnSlice = append(columnSlice, column{Name: c.Name, Comment: c.Comment, DataType: c.DataType})
	}
	return ColumnInfo2JavaInfo(columnSlice)
}
