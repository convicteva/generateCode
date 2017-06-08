package file

import (
	"os"
	"strings"
)

/**
创建相关目录并返回所有目录
*/
func CreatePackage(packageName string) DirInfo {
	var basePath string
	basePath = "demo" + pathSeparator + "src" + pathSeparator + "main" + pathSeparator

	var codeBasePath string
	codeBasePath = basePath + "java" + pathSeparator

	packageSlice := strings.Split(packageName, ".")
	for _, p := range packageSlice {
		codeBasePath += p + pathSeparator
	}
	modelPath := codeBasePath + "model"
	baseModelPath := modelPath + pathSeparator + "base"
	daoPath := codeBasePath + "dao"
	baseDaoPath := daoPath + pathSeparator + "base"
	managerPath := codeBasePath + "manager"
	baseManagerPath := managerPath + pathSeparator + "base"

	servicePath := codeBasePath + "service"

	os.MkdirAll(modelPath, 777)
	os.MkdirAll(baseModelPath, 777)
	os.MkdirAll(daoPath, 777)
	os.MkdirAll(baseDaoPath, 777)
	os.MkdirAll(baseManagerPath, 777)
	os.MkdirAll(servicePath, 777)

	resourcesPath := basePath + "resources"
	mapperPath := resourcesPath + pathSeparator + "mapper"
	os.MkdirAll(resourcesPath, 777)
	os.MkdirAll(mapperPath, 777)

	return DirInfo{modelPath, baseModelPath, daoPath, baseDaoPath, managerPath, baseManagerPath, mapperPath, servicePath}

}

/**
项目目录
*/
type DirInfo struct {
	ModelPath       string
	BaseModelPath   string
	DaoPath         string
	BaseDaoPath     string
	ManagerPath     string
	BaseManagerPath string
	MapperPath      string
	ServicePath     string
}
