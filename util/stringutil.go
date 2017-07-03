/**
字符串工具
*/
package stringutil

import (
	"strings"
)

/**
格式化表名为model 名称
示例：
        tb_user  -->  User
        tb_admin_role --> AdminRole
tableName 不能为nil  或者为 空白字符串
*/
func FormatTableNameToModelName(tableName string) string {
	if isBlank(tableName) {
		panic("字符串不能是空白的")
	}
	nameSlice := strings.Split(tableName, "_")
	var modelName string
	for _, v := range nameSlice {
		if v != "tb" {
			modelName += strings.Title(strings.ToLower(v))
		}
	}
	return modelName
}

func FormatColumnNameToProperty(columnName string) string {
	if isBlank(columnName) {
		panic("字符串不能是空白的")
	}

	if !strings.Contains(columnName, "_") {
		return strings.ToLower(columnName)
	} else {
		nameSlice := strings.Split(columnName, "_")
		var propertyName string
		for i, v := range nameSlice {
			v = strings.ToLower(v)
			if i > 0 {
				propertyName += strings.Title(v)
			} else {
				propertyName += v
			}
		}
		return propertyName
	}

}

/**
首字母小写
*/
func ToInitialLower(str string) string {
	if isBlank(str) {
		return ""
	} else {
		return strings.ToLower(string(str[0])) + string(str[1:])
	}
}

/**
str 是否是nil 或者是空白字符串
*/
func isBlank(str string) bool {
	if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}
