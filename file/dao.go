/**
dao 生成
*/
package file

import "fmt"

//生成dao
func GenerateDao(filePath, packageName, modelName string) {

	//文件全路径
	fullPath := filePath + pathSeparator + modelName + "Dao.java"
	fmt.Println(fullPath)
	//输入文件切片
	inputStr := make([]string, 0, 10)
	inputStr = append(inputStr, "package "+packageName+".dao;")
	inputStr = append(inputStr, "")
	inputStr = append(inputStr, "import "+packageName+".model."+modelName+";")
	inputStr = append(inputStr, "import "+packageName+".dao.base.MyBatisDao;")
	inputStr = append(inputStr, "import org.springframework.stereotype.Repository;")
	inputStr = append(inputStr, "")
	inputStr = append(inputStr, "@Repository")
	inputStr = append(inputStr, "public class "+modelName+"Dao extends MyBatisDao<"+modelName+"> {")
	inputStr = append(inputStr, "")
	inputStr = append(inputStr, "}")
	writeFile(fullPath, inputStr)

}
