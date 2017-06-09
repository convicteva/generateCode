/**
model 生成
*/
package file

import (
	"golang/db"
	"strings"
)

/**
生成BaseModel
*/
func GenerateBaseModel(filepath, packageName string) {
	//model 全路径
	fullPath := filepath + pathSeparator + "BaseModel.java"

	//基本列
	columnSlice := db.GetBaseColumn()

	//写入文件内容切片
	inputStrSlice := make([]string, 0, len(columnSlice)*4+10)

	//基本列对应的 列和java属性信息
	columnAndJavaInfo := db.ColumnInfo2JavaInfo(columnSlice)

	inputStrSlice = append(inputStrSlice, "package "+packageName+".model.base;")

	//生成导入信息
	javaImportSlice := generateImportByProperties(columnAndJavaInfo)
	inputStrSlice = append(inputStrSlice, javaImportSlice...)

	//导入
	inputStrSlice = append(inputStrSlice, "import java.io.Serializable;")
	inputStrSlice = append(inputStrSlice, "")

	inputStrSlice = append(inputStrSlice, "public class BaseModel implements Serializable {")
	inputStrSlice = append(inputStrSlice, "")

	//属性
	propertiesSlice := generateJavaPropertiesByProperties(columnAndJavaInfo)
	inputStrSlice = append(inputStrSlice, propertiesSlice...)

	//get set 方法
	getSetFuncSlice := generateGeterSeterFuncByProperties(columnAndJavaInfo)
	inputStrSlice = append(inputStrSlice, getSetFuncSlice...)

	inputStrSlice = append(inputStrSlice, "}")

	//写入文件
	writeFile(fullPath, inputStrSlice)

}

/**
根据列的切片，生成导入信息
*/
func generateImportByProperties(columnAndJavaInfoSlice []db.SqlColumnAndJavaPropertiesInfo) []string {
	javaImportSlice := make([]string, 0, 10)
	for _, c := range columnAndJavaInfoSlice {
		if strings.EqualFold(c.JavaType, "Date") && c {
			javaImportSlice = append(javaImportSlice, "import java.util.Date;")
		}
	}
	javaImportSlice = append(javaImportSlice, "")
	return javaImportSlice
}

/**
根据列生成java属性代码
*/
func generateJavaPropertiesByProperties(columnAndJavaInfoSlice []db.SqlColumnAndJavaPropertiesInfo) []string {
	propertiesSlice := make([]string, 0, len(columnAndJavaInfoSlice)*2)
	for _, c := range columnAndJavaInfoSlice {
		if !strings.EqualFold("", c.Comment) {
			propertiesSlice = append(propertiesSlice, javaCodeRetractSpace_1+"/** "+c.Comment+" */")
		}
		propertiesSlice = append(propertiesSlice, javaCodeRetractSpace_1+"private "+c.JavaType+" "+c.JavaPropertyName)
		propertiesSlice = append(propertiesSlice, "")
	}
	propertiesSlice = append(propertiesSlice, "")
	return propertiesSlice
}

/**
根据列信息生成get set 方法
返回插入信息
*/
func generateGeterSeterFuncByProperties(columnAndJavaInfoSlice []db.SqlColumnAndJavaPropertiesInfo) []string {
	getSetFunSlice := make([]string, 0, len(columnAndJavaInfoSlice)*8)

	for _, p := range columnAndJavaInfoSlice {
		//get 方法
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_1+"public "+p.JavaType+" "+"get"+strings.Title(p.JavaPropertyName)+"(){")
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_2+"return this."+p.JavaPropertyName)
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_1+"}")
		getSetFunSlice = append(getSetFunSlice, "")

		//set 方法
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_1+"public void "+"set"+strings.Title(p.JavaPropertyName)+"("+p.JavaType+" "+p.JavaPropertyName+"){")
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_2+"this."+p.JavaPropertyName+" = "+p.JavaPropertyName)
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_1+"}")
		getSetFunSlice = append(getSetFunSlice, "")
	}

	return getSetFunSlice
}

/**
生成model
*/
func GenerateMode(filepath, packageName, modelName string, columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) {

	//过滤掉BaseModel 中的属性
	commomModelColumnAndJavaInfo := make([]db.SqlColumnAndJavaPropertiesInfo, 0, len(columnAndJavaInfo))
	for _, c := range columnAndJavaInfo {
		if _, exists := db.BaseColumnMap[c.ColumnName]; !exists {
			commomModelColumnAndJavaInfo = append(commomModelColumnAndJavaInfo, c)
		}
	}

	fullPath := filepath + pathSeparator + modelName + ".java"
	//写入文件内容切片
	inputStrSlice := make([]string, 0, len(commomModelColumnAndJavaInfo)*4+10)

	inputStrSlice = append(inputStrSlice, "package "+packageName+".model;")
	inputStrSlice = append(inputStrSlice, "")

	inputStrSlice = append(inputStrSlice, "import "+packageName+".base.BaseModel;")
	//生成导入信息
	javaImportSlice := generateImportByProperties(commomModelColumnAndJavaInfo)
	inputStrSlice = append(inputStrSlice, javaImportSlice...)
	inputStrSlice = append(inputStrSlice, "")

	inputStrSlice = append(inputStrSlice, "public class "+modelName+" extends BaseModel {")
	inputStrSlice = append(inputStrSlice, "")

	//属性
	propertiesSlice := generateJavaPropertiesByProperties(commomModelColumnAndJavaInfo)
	inputStrSlice = append(inputStrSlice, propertiesSlice...)

	//get set 方法
	getSetFuncSlice := generateGeterSeterFuncByProperties(commomModelColumnAndJavaInfo)
	inputStrSlice = append(inputStrSlice, getSetFuncSlice...)

	inputStrSlice = append(inputStrSlice, "}")

	//写入文件
	writeFile(fullPath, inputStrSlice)

}
