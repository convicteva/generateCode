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
	columnSlice = append(columnSlice, db.Column{"order_id", "订单id", "BIGINT"})
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
	fmt.Fprintln(w, "")

	//属性
	propertiesSlice := generateJavaPropertiesByProperties(propertySlice)
	if len(propertiesSlice) > 0 {
		for _, v := range propertiesSlice {
			fmt.Fprintln(w, v)
		}
	}
	fmt.Fprintln(w, "")

	//get set 方法
	getSetFuncSlice := generateGeterSeterFuncByProperties(propertySlice)
	if len(getSetFuncSlice) > 0 {
		for _, s := range getSetFuncSlice {
			fmt.Fprintln(w, s)
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
func generateJavaPropertiesByProperties(properties []db.JavaProperty) []string {
	propertiesSlice := make([]string, 0, len(properties)*2)
	for _, p := range properties {
		if !strings.EqualFold("", p.Comment) {
			propertiesSlice = append(propertiesSlice, javaCodeRetractSpace_1+"/** "+p.Comment+" */")
		}
		propertiesSlice = append(propertiesSlice, javaCodeRetractSpace_1+"private "+p.DataType+" "+p.Name)
		propertiesSlice = append(propertiesSlice, "")
	}
	return propertiesSlice
}

/**
根据列信息生成get set 方法
返回插入信息
*/
func generateGeterSeterFuncByProperties(propertiesSlice []db.JavaProperty) []string {
	getSetFunSlice := make([]string, 0, len(propertiesSlice)*8)
	for _, p := range propertiesSlice {
		//get 方法
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_1+"public "+p.DataType+" "+"get"+strings.Title(p.Name)+"(){")
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_2+"return this."+p.Name)
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_1+"}")
		getSetFunSlice = append(getSetFunSlice, "")

		//set 方法
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_1+"public void "+"set"+strings.Title(p.Name)+"("+p.DataType+" "+p.Name+"){")
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_2+"this."+p.Name+" = "+p.Name)
		getSetFunSlice = append(getSetFunSlice, javaCodeRetractSpace_1+"}")
		getSetFunSlice = append(getSetFunSlice, "")

	}

	return getSetFunSlice
}

func GenerateMode(packageName, modelName string) {

}
