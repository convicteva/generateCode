/**
model 生成
*/
package file

import (
	"bufio"
	"fmt"
	"golang/db"
	"os"
	"strings"
)

func GenerateBaseModel(filepath, packageName string) {
	fullPath := filepath + pathSeparator + "BaseModel.java"
	//删除已经存在的BaseModel
	os.Remove(fullPath)
	//创建BaseModel
	f, _ := os.OpenFile(fullPath, os.O_CREATE, os.ModePerm)
	defer f.Close()

	//列定义
	columnSlice := make([]db.Column, 0, 4)
	columnSlice = append(columnSlice, db.Column{"id", "", "BIGINT"})
	columnSlice = append(columnSlice, db.Column{"remark", "", "VARCHAR"})
	columnSlice = append(columnSlice, db.Column{"create_time", "", "TIMESTAMP"})
	columnSlice = append(columnSlice, db.Column{"status", "", "INT"})

	//列对应的属性
	propertySlice := db.GetJavaProertyByColumn(columnSlice)

	javaImportSlice := generateImportByProperties(propertySlice)

	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "package "+packageName+".model.base;")

	//导入
	fmt.Fprintln(w, "import java.io.Serializable;")
	if len(javaImportSlice) > 0 {
		for _, v := range javaImportSlice {
			fmt.Fprintln(w, v)
		}
	}

	fmt.Fprintln(w, "public class BaseModel implements Serializable {")

	//属性
	propertiesSlice := generateJavaPropertiesByColumn(propertySlice)
	if len(propertiesSlice) > 0 {
		for _, v := range propertiesSlice {
			fmt.Fprintln(w, v)
		}
	}

	fmt.Fprintln(w, "}")

	w.Flush()
}

/**
根据列的切片，生成导入信息
*/
func generateImportByProperties(properties []db.JavaProperty) []string {
	javaImportSlice := make([]string, 0, 10)
	for _, v := range properties {
		if strings.EqualFold(v.DataType, "Date") {
			javaImportSlice = append(javaImportSlice, "import java.util.Date;")
		}
	}
	return javaImportSlice
}

/**
根据列生成java属性代码
*/
func generateJavaPropertiesByColumn(properties []db.JavaProperty) []string {
	propertiesSlice := make([]string, 0, 10)
	for _, p := range properties {
		propertiesSlice = append(propertiesSlice, javaCodeRetractSpace_1+"private "+p.DataType+" "+p.Name)
	}
	return propertiesSlice
}

type javaProperty struct {
	//字段名
	Name string
	//注释
	Comment string
	//类型
	DataType string
}

/**
根据列信息生成get set 方法
*/
func generateGeterSeterFuncByColumn(propertiesSlice []string) []string {
	return nil
}

func GenerateMode(packageName, modelName string) {

}
