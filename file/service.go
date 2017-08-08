/**
生成server 代码
*/
package file

import (
	"golang/config"
	"golang/util"
)

/**
生成manager
*/
func generateService(filePath, modelName string) {
	packageName := config.Project_package_name
	//文件全路径
	fullPath := filePath + pathSeparator + modelName + "Service.java"

	managerName := util.ToInitialLower(modelName) + "Manager"
	modelNameInitialLower := util.ToInitialLower(modelName)

	//输入文件切片
	inputStr := make([]string, 0, 10)
	inputStr = append(inputStr, "package "+packageName+".service;")
	inputStr = append(inputStr, "")
	inputStr = append(inputStr, "import "+packageName+".manager."+modelName+"Manager;")
	inputStr = append(inputStr, "import "+packageName+".model.Page;")
	inputStr = append(inputStr, "import "+packageName+".model."+modelName+";")
	inputStr = append(inputStr, "import "+packageName+".service.base.BaseService;")
	inputStr = append(inputStr, "import org.springframework.stereotype.Service;")
	inputStr = append(inputStr, "import javax.annotation.Resource;")
	inputStr = append(inputStr, "import java.util.List;")
	inputStr = append(inputStr, "import java.util.Map;")
	inputStr = append(inputStr, "")

	inputStr = append(inputStr, "@Service")
	inputStr = append(inputStr, "public class "+modelName+"Service  implements BaseService<"+modelName+"> {")
	inputStr = append(inputStr, "")
	//manager 注入
	inputStr = append(inputStr, javaCodeRetractSpace_1+"@Resource")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"private "+modelName+"Manager "+managerName+";")
	inputStr = append(inputStr, "")

	//save 方法实现
	inputStr = append(inputStr, javaCodeRetractSpace_1+"@Override")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"public void save("+modelName+" "+modelNameInitialLower+"){")
	inputStr = append(inputStr, javaCodeRetractSpace_2+managerName+".insert("+modelNameInitialLower+");")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"}")
	inputStr = append(inputStr, "")

	//update 方法实现
	inputStr = append(inputStr, javaCodeRetractSpace_1+"@Override")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"public void update("+modelName+" "+modelNameInitialLower+"){")
	inputStr = append(inputStr, javaCodeRetractSpace_2+managerName+".update("+modelNameInitialLower+");")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"}")
	inputStr = append(inputStr, "")

	//update 方法实现
	inputStr = append(inputStr, javaCodeRetractSpace_1+"@Override")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"public void delete(long id){")
	inputStr = append(inputStr, javaCodeRetractSpace_2+managerName+".delete(id);")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"}")
	inputStr = append(inputStr, "")

	//get 方法实现
	inputStr = append(inputStr, javaCodeRetractSpace_1+"@Override")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"public void get(long id){")
	inputStr = append(inputStr, javaCodeRetractSpace_2+managerName+".get(id);")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"}")
	inputStr = append(inputStr, "")

	//find page 方法实现
	inputStr = append(inputStr, javaCodeRetractSpace_1+"@Override")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"public Page<"+modelName+"> findPage(Page<"+modelName+"> page, Map<String, Object> filter) {")
	inputStr = append(inputStr, javaCodeRetractSpace_2+managerName+".findPage(page,filter);")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"}")
	inputStr = append(inputStr, "")

	//find page 方法实现
	inputStr = append(inputStr, javaCodeRetractSpace_1+"@Override")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"public List<"+modelName+"> findList(Map<String, Object> filter) {")
	inputStr = append(inputStr, javaCodeRetractSpace_2+managerName+".findList(page,filter);")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"}")
	inputStr = append(inputStr, "")

	//find page 方法实现
	inputStr = append(inputStr, javaCodeRetractSpace_1+"@Override")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"public long count(Map<String, Object> filter) {")
	inputStr = append(inputStr, javaCodeRetractSpace_2+managerName+".count(page,filter);")
	inputStr = append(inputStr, javaCodeRetractSpace_1+"}")

	inputStr = append(inputStr, "")

	inputStr = append(inputStr, "}")
	writeFile(fullPath, inputStr)
}
