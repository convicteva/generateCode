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

	//启动goroutine 数量
	goroutineNum := runtime.NumCPU()

	//任务通道
	jobChannel := make(chan map[string][]db.SqlColumnAndJavaPropertiesInfo, goroutineNum)
	done := make(chan bool, goroutineNum)

	//添加job
	addJob(jobChannel, tableColumnAndJavaInfoMap)
	//执行job
	doJob(dirInfo, goroutineNum, jobChannel, done)
	//等待执行完毕
	awaitCompletion(goroutineNum, done)

}

/**
添加任务
*/
func addJob(jobChannel chan map[string][]db.SqlColumnAndJavaPropertiesInfo, jobs map[string][]db.SqlColumnAndJavaPropertiesInfo) {
	go func() {
		for k, v := range jobs {
			jobChannel <- map[string][]db.SqlColumnAndJavaPropertiesInfo{k: v}
		}
		close(jobChannel)
	}()
}

/**
执行任务
创建 goroutineNum 数量的goroutine
*/
func doJob(dirInfo file.DirInfo, goroutineNum int, jobs chan map[string][]db.SqlColumnAndJavaPropertiesInfo, done chan bool) {
	for i := 0; i < goroutineNum; i++ {
		go func() {
			for job := range jobs {
				for k, v := range job {
					ganerateFile(dirInfo, k, v)
				}
			}
			//执行完后，往任务通道中发送一个完成标识
			done <- true
		}()
	}
}

/**
等待执行完毕
*/
func awaitCompletion(gorouniteNum int, done chan bool) {
	for i := 0; i < gorouniteNum; i++ {
		<-done
	}
	close(done)
}

/**
生成文件
*/
func ganerateFile(dirInfo file.DirInfo, tableName string, columnAndJavaInfo []db.SqlColumnAndJavaPropertiesInfo) {
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
}
