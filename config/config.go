package config

import (
	"strings"
)

const (
	//数据库连接配置
	MAXOPENCONNS int = 5
	MAXIDLECONNS int = 2

	//项目包名
	Project_package_name string = "com.masz.demo"
)

/**
不能使用db.Column ，出现循环依赖
*/
type ConfigColumn struct {
	Name     string
	Comment  string
	DataType string
}

//所有表拥有的共同字段
var BaseColumn = [3]ConfigColumn{ConfigColumn{"id", "", "BIGINT"}, ConfigColumn{"remark", "", "VARCHAR"}, ConfigColumn{"create_time", "", "TIMESTAMP"}}

//基础列对应的map，生成一般model 时，过滤使用
var BaseColumnMap = make(map[string]string, len(BaseColumn))

func init() {
	for _, c := range BaseColumn {
		BaseColumnMap[strings.ToUpper(c.Name)] = c.Name
	}
}
