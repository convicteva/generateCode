/**
数据库字段相关信息
*/
package db

import (
	"golang/util"
	"strings"
)

/**
mysql 数据类型对应的jdbc  和 java 的数据类型
*/
var MysqlTypeToJava = make(map[string]JdbcJavaTypeMap)

func init() {
	//mysql 的数据类型对java 数据类型的转化

	//数值
	MysqlTypeToJava["int"] = JdbcJavaTypeMap{"INTEGER", "Integer"}
	MysqlTypeToJava["TINYINT"] = JdbcJavaTypeMap{"INTEGER", "Integer"}
	MysqlTypeToJava["SMALLINT"] = JdbcJavaTypeMap{"INTEGER", "Integer"}
	MysqlTypeToJava["MEDIUMINT"] = JdbcJavaTypeMap{"INTEGER", "Integer"}
	MysqlTypeToJava["INT"] = JdbcJavaTypeMap{"INTEGER", "Integer"}
	MysqlTypeToJava["INTEGER"] = JdbcJavaTypeMap{"INTEGER", "Integer"}
	MysqlTypeToJava["BIGINT"] = JdbcJavaTypeMap{"BIGINT", "Long"}
	MysqlTypeToJava["FLOAT"] = JdbcJavaTypeMap{"DECIMAL", "Float"}
	MysqlTypeToJava["DOUBLE"] = JdbcJavaTypeMap{"DECIMAL", "Double"}
	MysqlTypeToJava["DECIMAL"] = JdbcJavaTypeMap{"DECIMAL", "Double"}

	//日期 DATE   TIME  YEAR  不支持
	MysqlTypeToJava["DATETIME"] = JdbcJavaTypeMap{"DATE", "Date"}
	MysqlTypeToJava["TIMESTAMP"] = JdbcJavaTypeMap{"TIMESTAMP", "Date"}

	//字符串  二进制的不支持
	MysqlTypeToJava["CHAR"] = JdbcJavaTypeMap{"CHAR", "char"}
	MysqlTypeToJava["VARCHAR"] = JdbcJavaTypeMap{"VARCHAR", "String"}
	MysqlTypeToJava["TINYTEXT"] = JdbcJavaTypeMap{"VARCHAR", "String"}
	MysqlTypeToJava["TEXT"] = JdbcJavaTypeMap{"VARCHAR", "String"}
	MysqlTypeToJava["MEDIUMTEXT"] = JdbcJavaTypeMap{"VARCHAR", "String"}
	MysqlTypeToJava["LONGTEXT"] = JdbcJavaTypeMap{"LONGVARCHAR", "String"}

	//mysqlTypeToJava["MEDIUMBLOB"] = "java.util.Date"
	//mysqlTypeToJava["LONGBLOB"] = "java.util.Date"
	//mysqlTypeToJava["BLOB"] = "java.util.Date"
	//mysqlTypeToJava["TINYBLOB"] = "java.util.Date"

}

/**
jdbc  java 关系
*/
type JdbcJavaTypeMap struct {
	jdbcType string
	javaType string
}

/**
表字段定义
*/
type Column struct {
	//字段名
	Name string
	//注释
	Comment string
	//类型
	DataType string
}

/**
java 属性
*/
type JavaProperty Column

/**
把mysql 数据类型，转化成jdbc 和 java数据类型
*/
func (column *Column) getJavaType() string {
	v, exists := MysqlTypeToJava[strings.ToUpper(column.DataType)]
	if exists {
		return v.javaType
	} else {
		panic(column.DataType + ",mysql 数据类型映射未找到")
	}
}

func GetJavaProertyByColumn(columns []Column) []JavaProperty {
	propertySlice := make([]JavaProperty, 0, 10)
	for _, c := range columns {
		propertySlice = append(propertySlice, JavaProperty{stringutil.FormatColumnNameToProperty(c.Name), c.Comment, c.getJavaType()})
	}
	return propertySlice
}
