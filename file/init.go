package file

import (
	"bufio"
	"fmt"
	"os"
)

const (
	javaCodeRetractSpace_1 = "    "
	javaCodeRetractSpace_2 = "        "
	javaCodeRetractSpace_3 = "            "
)

var pathSeparator string

func init() {
	pathSeparator = string(os.PathSeparator)
}

/**
写入文件
*/
func writeFile(fullPath string, inputStr []string) {
	//删除已经存在的BaseModel
	os.Remove(fullPath)
	//创建mapper 文件
	f, _ := os.OpenFile(fullPath, os.O_CREATE, os.ModePerm)
	defer f.Close()
	//写入文件
	w := bufio.NewWriter(f)
	for _, v := range inputStr {
		fmt.Fprintln(w, v)
	}
	w.Flush()
}
