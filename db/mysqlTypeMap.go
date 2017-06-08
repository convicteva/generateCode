package db

import "strings"

/**
mysql 数据类型对应的jdbc  和 java 的数据类型
*/
var MysqlTypeToJava = make(map[string]JdbcJavaTypeMap)

/**
jdbc  java 关系
*/
type JdbcJavaTypeMap struct {
	JdbcType string
	JavaType string
}

func init() {
	//mysql 的数据类型对java 数据类型的转化

	//数值
	MysqlTypeToJava["INT"] = JdbcJavaTypeMap{"INTEGER", "Integer"}
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
	MysqlTypeToJava["DATE"] = JdbcJavaTypeMap{"DATE", "Date"}
	MysqlTypeToJava["TIME"] = JdbcJavaTypeMap{"DATE", "Date"}
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

func GetJdbcTypeByMysqlType(mysqlType string) string {
	v, exists := MysqlTypeToJava[strings.ToUpper(mysqlType)]
	if exists {
		return v.JdbcType
	} else {
		panic("mysql 数据类型" + mysqlType + ", 不支持")
	}
}
