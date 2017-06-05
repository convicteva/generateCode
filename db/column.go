package db

var mysqlTypeToJava = make(map[string]string)

func init() {
	//TODO mysql 的数据类型对java 数据类型的转化
	mysqlTypeToJava["int"] = "java.long.Integer"

}

/**
表字段定义
*/
type Column struct {
	//字段名
	name string
	//注释
	comment string
	//类型
	dataType string
}

/**
把mysql 数据类型，转化成java数据类型
*/
func (column *Column) GetJavaType() string {
	v, exists := mysqlTypeToJava[column.dataType]
	if exists {
		return v
	} else {
		panic(column.dataType + ",不存在java数据类型")
	}

}
