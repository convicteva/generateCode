/**
manager 生成
*/
package file

import (
	"genereateCode/config"
	"genereateCode/util"
)

/**
生成manager
*/
func generateManager(filePath, modelName string) {

	packageName := config.Project_package_name

	//文件全路径
	fullPath := filePath + pathSeparator + modelName + "Manager.java"
	//输入文件切片
	inputStr := make([]string, 0, 10)
	inputStr = append(inputStr, "package "+packageName+".manager;")
	inputStr = append(inputStr, "")
	inputStr = append(inputStr, "import "+packageName+".model."+modelName+";")
	inputStr = append(inputStr, "import "+packageName+".manager.base.AbstractBaseManager;")
	inputStr = append(inputStr, "import "+packageName+".model."+modelName+"Dao;")
	inputStr = append(inputStr, "import org.springframework.stereotype.Service;")
	inputStr = append(inputStr, "import javax.annotation.Resource;")
	inputStr = append(inputStr, "")

	inputStr = append(inputStr, "@Service")
	inputStr = append(inputStr, "public class "+modelName+"Manager extends AbstractBaseManager<"+modelName+"> {")
	inputStr = append(inputStr, "")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"@Resource")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"private "+modelName+"Dao "+util.ToInitialLower(modelName)+"Dao;")
	inputStr = append(inputStr, "")
	inputStr = append(inputStr, "}")
	writeFile(fullPath, inputStr)
}
