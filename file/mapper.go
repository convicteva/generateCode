/**
mapper xml 生成
*/
package file

import (
	"golang/db"
	"golang/util"
	"strings"
)

const (
	SQL_BASE_COLUMN_NAME = "Base_Column_List"
)

/**
生成mapper
*/
func GenerateMapper(filepath, packageName, modelName, tableName string, columnSlice []db.Column) {

	//文件全路径
	fullPath := filepath + pathSeparator + modelName + "Mapper.xml"

	//mapper 文件中使用的model 全路径，如：com.masz.demo.model.User
	modelFullPath := packageName + ".model." + modelName

	//resultMapName,如UserMap
	resultMapName := modelName + "Map"

	//输入文件的切片
	inputStr := make([]string, 0, len(columnSlice)*5)

	inputStr = append(inputStr, `<?xml version="1.0" encoding="UTF-8" ?>`)
	inputStr = append(inputStr, `<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >`)
	inputStr = append(inputStr, `<mapper namespace="`+modelFullPath+`">`)
	inputStr = append(inputStr, ``)

	//生成resultMap
	resultMapSlice := generateResultMap(modelFullPath, resultMapName, columnSlice)
	inputStr = append(inputStr, resultMapSlice...)

	//生成所有字段sql
	columnSql := generateColumnSql(columnSlice)
	inputStr = append(inputStr, columnSql...)

	//生成insert
	insertSql := generateInsertSql(modelFullPath, tableName, columnSlice)
	inputStr = append(inputStr, insertSql...)

	//生成delete sql
	delSql := generateDelSql(tableName)
	inputStr = append(inputStr, delSql...)

	//生成update sql
	updateSql := generateUpdateSql(modelFullPath, tableName, columnSlice)
	inputStr = append(inputStr, updateSql...)

	//生成getsql
	getSql := generateGetSql(tableName, resultMapName)
	inputStr = append(inputStr, getSql...)

	inputStr = append(inputStr, `</mapper>`)

	//写入文件
	writeFile(fullPath, inputStr)
}

/**
生成resultMap
*/
func generateResultMap(modelFullPath, resultMapName string, columnSlice []db.Column) []string {
	resultMapSlice := make([]string, 0, len(columnSlice))

	resultMapSlice = append(resultMapSlice, javaCodeRetractSpace_1+`<resultMap type="`+modelFullPath+`" id="`+resultMapName+`">`)
	resultSlice := generateResult(columnSlice)
	resultMapSlice = append(resultMapSlice, resultSlice...)
	resultMapSlice = append(resultMapSlice, javaCodeRetractSpace_1+`</resultMap>`)
	resultMapSlice = append(resultMapSlice, ``)
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
		resultSlice = append(resultSlice, javaCodeRetractSpace_2+`<`+resultTag+` column="`+strings.ToUpper(c.Name)+`" jdbcType="`+jdbcJavaTypeMap.JdbcType+`" property="`+property+`"/>`)
	}
	resultSlice = append(resultSlice, ``)
	return resultSlice
}

/**
生成所有字段的sql 片段
*/
func generateColumnSql(columnSlice []db.Column) []string {
	length := len(columnSlice)
	sqlSegmentSlice := make([]string, 0, length)
	sqlSegmentSlice = append(sqlSegmentSlice, javaCodeRetractSpace_1+`<sql id="`+SQL_BASE_COLUMN_NAME+`">`)
	c := ""
	for i, v := range columnSlice {
		c = javaCodeRetractSpace_2 + strings.ToUpper(v.Name) + ","
		if i == length-1 {
			c = strings.Replace(c, ",", "", 1)
		}
		sqlSegmentSlice = append(sqlSegmentSlice, c)
	}
	sqlSegmentSlice = append(sqlSegmentSlice, javaCodeRetractSpace_1+`</sql>`)

	sqlSegmentSlice = append(sqlSegmentSlice, ``)
	return sqlSegmentSlice
}

/**
生成insert 语句切片
*/
func generateInsertSql(modelFullPath, tableName string, columnSlice []db.Column) []string {
	length := len(columnSlice)
	sqlSlice := make([]string, 0, length*2+6)
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

	sqlSlice = append(sqlSlice, ``)
	return sqlSlice
}

/**
delete sql
*/
func generateDelSql(tableName string) []string {
	sqlSlice := make([]string, 0, 4)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`<delete id="delete" parameterType="long" >`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`delete from `+tableName)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`where id = #{id}`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`</delete>`)
	sqlSlice = append(sqlSlice, ``)
	return sqlSlice
}

/**
生成update 语句
*/
func generateUpdateSql(modelFullPath, tableName string, columnSlice []db.Column) []string {
	length := len(columnSlice)
	sqlSlice := make([]string, 0, length*3+5)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`<update id="update" parameterType="`+modelFullPath+`">`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"UPDATE "+strings.ToUpper(tableName))
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"<set>")
	for _, v := range columnSlice {
		if !strings.EqualFold(strings.ToUpper(v.Name), "ID") {
			sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`<if test="`+stringutil.FormatColumnNameToProperty(v.Name)+`!=null">`)
			sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+strings.ToUpper(v.Name)+`=#{`+stringutil.FormatTableNameToModelName(v.Name)+`,jdbcType=`+db.GetJdbcTypeByMysqlType(v.DataType)+`},`)
			sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`</if>`)
		}
	}
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"</set> ")
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"where id=#{id}")
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`</update>`)
	sqlSlice = append(sqlSlice, ``)
	return sqlSlice
}

/**
生成get 语句
*/
func generateGetSql(tableName, resultMapName string) []string {
	sqlSlice := make([]string, 0, 4)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`<select id="get" parameterType="long" resultMap="`+resultMapName+`">`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`select <include refid="`+SQL_BASE_COLUMN_NAME+`" /> from `+tableName)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`where id = #{id}`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`</select>`)
	sqlSlice = append(sqlSlice, ``)
	return sqlSlice
}
