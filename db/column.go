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

//所有表中都会有的列，把放入到BaseModel中
var BaseColumn = make([]Column, 0, 3)
var baseColumnMap = make(map[string]string, 3)

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

	//基础列
	BaseColumn = append(BaseColumn, Column{"id", "", "BIGINT"})
	BaseColumn = append(BaseColumn, Column{"remark", "", "VARCHAR"})
	BaseColumn = append(BaseColumn, Column{"create_time", "", "TIMESTAMP"})

	//基础列对应的map，生成一般model 时，过滤使用
	baseColumnMap["id"] = "id"
	baseColumnMap["remark"] = "remark"
	baseColumnMap["create_time"] = "create_time"
}

func GetBaseColumn() []Column {
	return BaseColumn
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

/**
生成BaseModel 中的列的 java 属性
*/
func GetJavaBaseProertyByColumn(columns []Column) []JavaProperty {
	propertySlice := make([]JavaProperty, 0, 10)
	for _, c := range columns {
		propertySlice = append(propertySlice, JavaProperty{stringutil.FormatColumnNameToProperty(c.Name), c.Comment, c.getJavaType()})
	}
	return propertySlice
}

/**
根据column 获取java 的类型，去除BaseModel 中的列
*/
func GetJavaProertyByColumn(columns []Column) []JavaProperty {
	propertySlice := make([]JavaProperty, 0, 10)
	for _, c := range columns {
		if _, exists := baseColumnMap[strings.ToLower(c.Name)]; !exists {
			propertySlice = append(propertySlice, JavaProperty{stringutil.FormatColumnNameToProperty(c.Name), c.Comment, c.getJavaType()})
		}
	}
	return propertySlice
}
