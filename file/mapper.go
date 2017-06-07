/**
mapper xml 生成
*/
package file

import (
	"bufio"
	"fmt"
	"golang/db"
	"golang/util"
	"os"
	"strings"
)

const (
	SQL_COLUMN_NAME = "Base_Column_List"
)

/**
生成mapper
*/
func GenerateMapper(filepath, packageName, modelName, tableName string, columnSlice []db.Column) {

	//文件全路径
	fullPath := filepath + pathSeparator + modelName + "Mapper.xml"

	//mapper 文件中使用的model 全路径，如：com.masz.demo.model.User
	modelFullPath := packageName + ".model." + modelName

	//删除已经存在的BaseModel
	os.Remove(fullPath)
	//创建BaseModel
	f, _ := os.OpenFile(fullPath, os.O_CREATE, os.ModePerm)
	defer f.Close()

	w := bufio.NewWriter(f)

	fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8" ?>`)
	fmt.Fprintln(w, `<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >`)
	fmt.Fprintln(w, `<mapper namespace="`+modelFullPath+`">`)
	fmt.Fprintln(w, ``)

	//生成resultMap
	resultMapSlice := generateResultMap(modelFullPath, modelName, columnSlice)
	for _, v := range resultMapSlice {
		fmt.Fprintln(w, v)
	}
	fmt.Fprintln(w, ``)

	//生成所有字段sql
	columnSql := generateColumnSql(columnSlice)
	for _, v := range columnSql {
		fmt.Fprintln(w, v)
	}
	fmt.Fprintln(w, ``)

	//生成insert
	insertSql := generateInsertSql(modelFullPath, tableName, columnSlice)
	for _, v := range insertSql {
		fmt.Fprintln(w, v)
	}
	fmt.Fprintln(w, ``)

	fmt.Fprintln(w, `</mapper>`)
	w.Flush()

}

/**
生成resultMap
*/
func generateResultMap(modelFullPath, modelName string, columnSlice []db.Column) []string {
	resultMapSlice := make([]string, 0, len(columnSlice))

	resultMapSlice = append(resultMapSlice, javaCodeRetractSpace_1+`<resultMap type="`+modelFullPath+`" id="`+modelName+`Map">`)
	resultSlice := generateResult(columnSlice)
	resultMapSlice = append(resultMapSlice, resultSlice...)
	resultMapSlice = append(resultMapSlice, javaCodeRetractSpace_1+`</resultMap>`)
	return resultMapSlice
}

/**
生成resultMap  result的子元素
*/
func generateResult(columnSlice []db.Column) []string {
	resultSlice := make([]string, 0, len(columnSlice))
	resultTag := "result"
	for _, c := range columnSlice {
		jdbcJavaTypeMap := db.MysqlTypeToJava[strings.ToUpper(c.DataType)]
		property := stringutil.FormatColumnNameToProperty(c.Name)
		if strings.EqualFold(property, "ID") {
			resultTag = "id"
		} else {
			resultTag = "result"
		}
		resultSlice = append(resultSlice, javaCodeRetractSpace_2+`<`+resultTag+` column="`+c.Name+`" jdbcType="`+jdbcJavaTypeMap.JdbcType+`" property="`+property+`"/>`)
	}
	return resultSlice
}

/**
生成所有字段的sql 片段
*/
func generateColumnSql(columnSlice []db.Column) []string {
	length := len(columnSlice)
	sqlSegmentSlice := make([]string, 0, length)
	sqlSegmentSlice = append(sqlSegmentSlice, javaCodeRetractSpace_1+`<sql id="`+SQL_COLUMN_NAME+`">`)
	c := ""
	for i, v := range columnSlice {
		c = javaCodeRetractSpace_2 + strings.ToUpper(v.Name) + ","
		if i == length-1 {
			c = strings.Replace(c, ",", "", 1)
		}
		sqlSegmentSlice = append(sqlSegmentSlice, c)
	}
	sqlSegmentSlice = append(sqlSegmentSlice, javaCodeRetractSpace_1+`</sql>`)
	return sqlSegmentSlice
}

/**
生成insert 语句切片
*/
func generateInsertSql(modelFullPath, tableName string, columnSlice []db.Column) []string {
	length := len(columnSlice)
	sqlSlice := make([]string, 0, length*2+5)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`<insert id="insert" parameterType="`+modelFullPath+`">`)

	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"INSERT INTO "+strings.ToUpper(tableName)+"(")

	c := ""
	for i, v := range columnSlice {
		c = javaCodeRetractSpace_2 + strings.ToUpper(v.Name) + ","
		if i == length-1 {
			c = strings.Replace(c, ",", "", 1)
		}
		sqlSlice = append(sqlSlice, c)
	}
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+")")
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"values(")
	for i, v := range columnSlice {
		c = javaCodeRetractSpace_2 + `#{` + stringutil.FormatColumnNameToProperty(v.Name) + `,jdbcType=` + db.GetJdbcTypeByMysqlType(v.DataType) + `},`
		if i == length-1 {
			c = strings.Replace(c, "},", "}", -1)
		}
		sqlSlice = append(sqlSlice, c)
	}
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+")")

	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`</insert>`)

	return sqlSlice
}
