/**
mybatis mapper 生成器
*/
package main

import (
	"genereateCode/config"
	"genereateCode/configureparse"
	"genereateCode/db"
	"genereateCode/file"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime"
	"strings"
)

//输出代码目录
var root_path string = ""

var dirInfo file.DirInfo

func main() {

	router := gin.Default()

	//静态资源
	router.Static("/static/js", "./webapp/js")
	router.Static("/static/css", "./webapp/css")
	router.Static("/static/img", "./webapp/img")

	//定义模板文件路径
	router.LoadHTMLGlob("./webapp/templates/*")

	//路由
	router.GET("/", func(c *gin.Context) {
		//数据库配置结点
		nodeSlice := configureparse.GetDBConfigNode()
		c.HTML(http.StatusOK, "index.html", gin.H{"nodeSlice": nodeSlice})
	})

	//数据结点所有的table
	router.GET("/getTable/:node", func(c *gin.Context) {
		node := c.Param("node")
		db.InitDB(node)
		c.JSON(http.StatusOK, db.GetTableName())
	})

	router.Run(":8000")

}

/**
生成代码测试
*/
func testGenerate() {
	//将从web 传入
	var packageName = "com.masz.demo"
	var node = "test"
	var tableNameSlice []string = nil
	generate(packageName, node, tableNameSlice)
}

/**
生成代码
*/
func generate(packageName, node string, tableNameSlice []string) {
	initBaseInfo(packageName, node)
	file.Generate(dirInfo, tableNameSlice)
}

func initBaseInfo(packageName, node string) {
	//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	var system = strings.ToUpper(runtime.GOOS)
	//根据操作系统使用不同的默认目录。留以后导出使用
	if strings.EqualFold(system, "WINDOWS") {
		root_path = "C:\\temp\\generateCode\\"
	} else if strings.EqualFold(system, "LINUX") {
		root_path = "/tmp/generateCode/"
	}
	//删除之前存在的内容
	os.RemoveAll(root_path)

	config.Project_package_name = packageName
	//生成相关目录
	dirInfo = file.CreatePackage(root_path, config.Project_package_name)
	db.InitDB(node)
}
