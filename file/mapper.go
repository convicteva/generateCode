/**
mapper xml 生成
*/
package file

import (
	"golang/config"
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
func GenerateMapper(filepath, modelName, tableName string, columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) {

	packageName := config.Project_package_name

	//文件全路径
	fullPath := filepath + pathSeparator + modelName + "Mapper.xml"

	//mapper 文件中使用的model 全路径，如：com.masz.demo.model.User
	modelFullPath := packageName + ".model." + modelName

	//resultMapName,如UserMap
	resultMapName := modelName + "Map"

	//输入文件的切片
	inputStr := make([]string, 0, len(columnAndJavaInfo)*5)

	inputStr = append(inputStr, `<?xml version="1.0" encoding="UTF-8" ?>`)
	inputStr = append(inputStr, `<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >`)
	inputStr = append(inputStr, `<mapper namespace="`+modelFullPath+`">`)
	inputStr = append(inputStr, ``)

	//生成resultMap
	resultMapSlice := generateResultMap(modelFullPath, resultMapName, columnAndJavaInfo)
	inputStr = append(inputStr, resultMapSlice...)

	//生成所有字段sql
	columnSql := generateColumnSql(columnAndJavaInfo)
	inputStr = append(inputStr, columnSql...)

	//生成insert
	insertSql := generateInsertSql(modelFullPath, tableName, columnAndJavaInfo)
	inputStr = append(inputStr, insertSql...)

	//生成delete sql
	delSql := generateDelSql(tableName)
	inputStr = append(inputStr, delSql...)

	//生成update sql
	updateSql := generateUpdateSql(modelFullPath, tableName, columnAndJavaInfo)
	inputStr = append(inputStr, updateSql...)

	//生成count sql
	consql := generateCoun(tableName, columnAndJavaInfo)
	inputStr = append(inputStr, consql...)

	//生成getsql
	getSql := generateGetSql(tableName, resultMapName)
	inputStr = append(inputStr, getSql...)

	//生成findList
	findListSql := generateFindListSql(tableName, resultMapName, columnAndJavaInfo)
	inputStr = append(inputStr, findListSql...)

	//生成分页查询sql
	findPageSql := generateFindPageSql(tableName, resultMapName, columnAndJavaInfo)
	inputStr = append(inputStr, findPageSql...)

	inputStr = append(inputStr, `</mapper>`)

	//写入文件
	writeFile(fullPath, inputStr)
}

/**
生成resultMap
*/
func generateResultMap(modelFullPath, resultMapName string, columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) []string {
	resultMapSlice := make([]string, 0, len(columnAndJavaInfo))
	resultMapSlice = append(resultMapSlice, javaCodeRetractSpace_1+`<resultMap type="`+modelFullPath+`" id="`+resultMapName+`">`)
	resultSlice := generateResult(columnAndJavaInfo)
	resultMapSlice = append(resultMapSlice, resultSlice...)
	resultMapSlice = append(resultMapSlice, javaCodeRetractSpace_1+`</resultMap>`)
	resultMapSlice = append(resultMapSlice, ``)
	return resultMapSlice
}

/**
生成resultMap  result的子元素
*/
func generateResult(columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) []string {
	resultSlice := make([]string, 0, len(columnAndJavaInfo))
	resultTag := "result"
	for _, c := range columnAndJavaInfo {
		if strings.EqualFold(c.ColumnName, "ID") {
			resultTag = "id"
		} else {
			resultTag = "result"
		}
		resultSlice = append(resultSlice, javaCodeRetractSpace_2+`<`+resultTag+` column="`+strings.ToUpper(c.ColumnName)+`" jdbcType="`+c.JdbcType+`" property="`+c.JavaPropertyName+`"/>`)
	}
	return resultSlice
}

/**
生成所有字段的sql 片段
*/
func generateColumnSql(columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) []string {
	length := len(columnAndJavaInfo)
	sqlSegmentSlice := make([]string, 0, length)
	sqlSegmentSlice = append(sqlSegmentSlice, javaCodeRetractSpace_1+`<sql id="`+SQL_BASE_COLUMN_NAME+`">`)
	c := ""
	for i, v := range columnAndJavaInfo {
		c = javaCodeRetractSpace_2 + strings.ToUpper(v.ColumnName) + ","
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
func generateInsertSql(modelFullPath, tableName string, columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) []string {
	length := len(columnAndJavaInfo)
	sqlSlice := make([]string, 0, length*2+6)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`<insert id="insert" parameterType="`+modelFullPath+`">`)

	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"INSERT INTO "+strings.ToUpper(tableName)+"(")

	c := ""
	for i, v := range columnAndJavaInfo {
		c = javaCodeRetractSpace_2 + strings.ToUpper(v.ColumnName) + ","
		if i == length-1 {
			c = strings.Replace(c, ",", "", 1)
		}
		sqlSlice = append(sqlSlice, c)
	}
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+")")
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"values(")
	for i, v := range columnAndJavaInfo {
		c = javaCodeRetractSpace_2 + `#{` + v.JavaPropertyName + `,jdbcType=` + v.JdbcType + `},`
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
func generateUpdateSql(modelFullPath, tableName string, columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) []string {
	sqlSlice := make([]string, 0, len(columnAndJavaInfo)*3+5)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`<update id="update" parameterType="`+modelFullPath+`">`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"UPDATE "+strings.ToUpper(tableName))
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+"<set>")
	for _, v := range columnAndJavaInfo {
		if !strings.EqualFold(strings.ToUpper(v.ColumnName), "ID") {
			sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`<if test="`+v.JavaPropertyName+`!=null">`)
			sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+strings.ToUpper(v.ColumnName)+`=#{`+v.JavaPropertyName+`,jdbcType=`+v.JdbcType+`},`)
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

/**
生成count
*/
func generateCoun(tableName string, columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) []string {
	sqlSlice := make([]string, 0, len(columnAndJavaInfo)*3+5)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`<select id="count" parameterType="map" resultType="long">`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`select count(1) from `+tableName)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`<where>`)
	for _, v := range columnAndJavaInfo {
		sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`<if test="`+stringutil.FormatColumnNameToProperty(v.JavaPropertyName)+`!=null">`)
		sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`and `+strings.ToUpper(v.ColumnName)+`=#{`+v.JavaPropertyName+`,jdbcType=`+v.JdbcType+`}`)
		sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`</if>`)
	}
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`</where>`)

	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`</select>`)
	sqlSlice = append(sqlSlice, ``)
	return sqlSlice
}

//findList
func generateFindListSql(tableName, resultMapName string, columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) []string {
	sqlSlice := make([]string, 0, 4)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`<select id="findList" parameterType="map" resultMap="`+resultMapName+`">`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`select <include refid="`+SQL_BASE_COLUMN_NAME+`" /> from `+tableName)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`<where>`)
	for _, v := range columnAndJavaInfo {
		sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`<if test="`+v.JavaPropertyName+`!=null">`)
		sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`and `+strings.ToUpper(v.ColumnName)+`=#{`+v.JavaPropertyName+`,jdbcType=`+v.JdbcType+`}`)
		sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`</if>`)
	}
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`</where>`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`</select>`)
	sqlSlice = append(sqlSlice, ``)
	return sqlSlice
}

//findPage
func generateFindPageSql(tableName, resultMapName string, columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) []string {
	sqlSlice := make([]string, 0, 4)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`<select id="findPage" parameterType="map" resultMap="`+resultMapName+`">`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`select <include refid="`+SQL_BASE_COLUMN_NAME+`" /> from `+tableName)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`<where>`)
	for _, v := range columnAndJavaInfo {
		sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`<if test="`+v.JavaPropertyName+`!=null">`)
		sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`and `+v.ColumnName+`=#{`+v.JavaPropertyName+`,jdbcType=`+v.JdbcType+`}`)
		sqlSlice = append(sqlSlice, javaCodeRetractSpace_3+`</if>`)
	}
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`</where>`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_2+`LIMIT #{offset},#{pageSize}`)
	sqlSlice = append(sqlSlice, javaCodeRetractSpace_1+`</select>`)
	sqlSlice = append(sqlSlice, ``)
	return sqlSlice
}
