/**
mybatis mapper 生成器
*/
package main

import (
	"golang/config"
	"golang/db"
	"golang/file"
	"golang/util"
	"os"
	"runtime"
	"strings"
)

//输出代码目录
var root_path string = ""

func init() {
	//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	var system = strings.ToUpper(runtime.GOOS)
	//根据操作系统使用不同的默认目录。留以后导出使用
	if strings.EqualFold(system, "WINDOWS") {
		root_path = "C:\\temp\\"
	} else if strings.EqualFold(system, "LINUX") {
		root_path = "/tmp/"
	}
	//删除之前存在的内容
	os.Remove(root_path)
}

func main() {
	//生成相关目录
	dirInfo := file.CreatePackage(root_path, config.Project_package_name)

	//表对应的字段  map
	tableColumnAndJavaInfoMap := db.GetTableInfo(nil)

	//生成BaseModel
	file.GenerateBaseModel(dirInfo.BaseModelPath, config.Project_package_name)

	//任务通道,通道长度为表的个数。
	jobs := make(chan bool, len(tableColumnAndJavaInfoMap))

	for tableName, columnAndJavaInfo := range tableColumnAndJavaInfoMap {
		//第一个表，产生一个协程
		go func() {
			//表名对应的modelName
			modelName := stringutil.FormatTableNameToModelName(tableName)

			//生成model
			file.GenerateMode(dirInfo.ModelPath, modelName, columnAndJavaInfo)

			//生成dao
			file.GenerateDao(dirInfo.DaoPath, modelName)

			//生成mapper
			file.GenerateMapper(dirInfo.MapperPath, modelName, tableName, columnAndJavaInfo)

			//生成manager
			file.GenerateManager(dirInfo.ManagerPath, modelName)

			//生成service
			file.GenerateService(dirInfo.ServicePath, modelName)

			//执行完后，往任务通道中发送一个完成标识
			jobs <- true
		}()
	}
	for i := 0; i < len(tableColumnAndJavaInfoMap); i++ {
		//主的 goroutine,等待任务goroutine 执行完成
		<-jobs
	}
	close(jobs)
}
