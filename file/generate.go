/**
生成文件
*/
package file

import (
	"golang/config"
	"golang/db"
	"golang/util"
	"runtime"
)

/**
生成文件
*/
func Generate(dirInfo DirInfo, tableNameSlice []string) {

	//表对应的字段  map
	tableColumnAndJavaInfo := db.GetTableInfo(tableNameSlice)

	//生成BaseModel
	generateBaseModel(dirInfo.BaseModelPath, config.Project_package_name)

	//启动goroutine 数量
	goroutineNum := runtime.NumCPU()

	//任务通道
	jobChannel := make(chan db.TableColumnAndJavaInfo, goroutineNum)
	done := make(chan bool, goroutineNum)

	//添加job
	addJob(jobChannel, tableColumnAndJavaInfo)
	//执行job
	doJob(dirInfo, goroutineNum, jobChannel, done)
	//等待执行完毕
	awaitCompletion(goroutineNum, done)
}

/**
添加任务
*/
func addJob(jobChannel chan db.TableColumnAndJavaInfo, jobs []db.TableColumnAndJavaInfo) {
	go func() {
		for _, job := range jobs {
			jobChannel <- job
		}
		close(jobChannel)
	}()
}

/**
执行任务
创建 goroutineNum 数量的goroutine
*/
func doJob(dirInfo DirInfo, goroutineNum int, jobs chan db.TableColumnAndJavaInfo, done chan bool) {
	for i := 0; i < goroutineNum; i++ {
		go func() {
			for job := range jobs {
				ganerateFile(dirInfo, job)
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
func ganerateFile(dirInfo DirInfo, tableColumnAndJavaInfo db.TableColumnAndJavaInfo) {
	columnAndJavaInfo := tableColumnAndJavaInfo.ColumnInfo
	tableName := tableColumnAndJavaInfo.TableName
	//表名对应的modelName
	modelName := stringutil.FormatTableNameToModelName(tableName)

	//生成model
	generateMode(dirInfo.ModelPath, modelName, columnAndJavaInfo)

	//生成dao
	generateDao(dirInfo.DaoPath, modelName)

	//生成mapper
	generateMapper(dirInfo.MapperPath, modelName, tableName, columnAndJavaInfo)

	//生成manager
	generateManager(dirInfo.ManagerPath, modelName)

	//生成service
	generateService(dirInfo.ServicePath, modelName)
}
