/**
数据库字段相关信息
*/
package db

import "strings"

/**
mysql 数据类型对应的jdbc  和 java 的数据类型
*/
var mysqlTypeToJava = make(map[string]JdbcJavaTypeMap)

func init() {
	//mysql 的数据类型对java 数据类型的转化

	//数值
	mysqlTypeToJava["int"] = JdbcJavaTypeMap{"INTEGER", "java.lang.Integer"}
	mysqlTypeToJava["TINYINT"] = JdbcJavaTypeMap{"INTEGER", "java.lang.Integer"}
	mysqlTypeToJava["SMALLINT"] = JdbcJavaTypeMap{"INTEGER", "java.lang.Integer"}
	mysqlTypeToJava["MEDIUMINT"] = JdbcJavaTypeMap{"INTEGER", "java.lang.Integer"}
	mysqlTypeToJava["INT"] = JdbcJavaTypeMap{"INTEGER", "java.lang.Integer"}
	mysqlTypeToJava["INTEGER"] = JdbcJavaTypeMap{"INTEGER", "java.lang.Integer"}
	mysqlTypeToJava["BIGINT"] = JdbcJavaTypeMap{"BIGINT", "java.lang.Long"}
	mysqlTypeToJava["FLOAT"] = JdbcJavaTypeMap{"DECIMAL", "java.lang.Float"}
	mysqlTypeToJava["DOUBLE"] = JdbcJavaTypeMap{"DECIMAL", "java.lang.Double"}
	mysqlTypeToJava["DECIMAL"] = JdbcJavaTypeMap{"DECIMAL", "java.math.BigDecimal"}

	//日期 DATE   TIME  YEAR  不支持
	mysqlTypeToJava["DATETIME"] = JdbcJavaTypeMap{"DATE", "java.util.Date"}
	mysqlTypeToJava["TIMESTAMP"] = JdbcJavaTypeMap{"TIMESTAMP", "java.util.Date"}

	//字符串  二进制的不支持
	mysqlTypeToJava["CHAR"] = JdbcJavaTypeMap{"CHAR", "java.lang.Character"}
	mysqlTypeToJava["VARCHAR"] = JdbcJavaTypeMap{"VARCHAR", "java.lang.String"}
	mysqlTypeToJava["TINYTEXT"] = JdbcJavaTypeMap{"VARCHAR", "java.lang.String"}
	mysqlTypeToJava["TEXT"] = JdbcJavaTypeMap{"VARCHAR", "java.lang.String"}
	mysqlTypeToJava["MEDIUMTEXT"] = JdbcJavaTypeMap{"VARCHAR", "java.lang.String"}
	mysqlTypeToJava["LONGTEXT"] = JdbcJavaTypeMap{"LONGVARCHAR", "java.lang.String"}

	//mysqlTypeToJava["MEDIUMBLOB"] = "java.util.Date"
	//mysqlTypeToJava["LONGBLOB"] = "java.util.Date"
	//mysqlTypeToJava["BLOB"] = "java.util.Date"
	//mysqlTypeToJava["TINYBLOB"] = "java.util.Date"

}

/**
jdbc  java 三者关系
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
把mysql 数据类型，转化成jdbc 和 java数据类型
*/
func (column *Column) GetJavaType() JdbcJavaTypeMap {
	v, exists := mysqlTypeToJava[strings.ToUpper(column.DataType)]
	if exists {
		return v
	} else {
		panic(column.DataType + ",mysql 数据类型映射未找到")
	}

}
