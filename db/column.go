/**
数据库字段相关信息
*/
package db

import (
	"golang/util"
	"strings"
)

//所有表中都会有的列，把放入到BaseModel中
var BaseColumn = make([]Column, 0, 3)
var baseColumnMap = make(map[string]string, 3)

func init() {

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
		return v.JavaType
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
