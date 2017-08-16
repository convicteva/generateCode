/**
数据库配置文件dbconfig解析
*/
package configureparse

import (
	"bufio"
	"fmt"
	"genereateCode/err"
	"io"
	"os"
	"reflect"
	"strings"
)

//配置文件地址
const file_path = "./dbconfig"

func GetDBConfig(node string) (DBConfig, error) {

	//配置文件转成的map
	confMap := findNodeMap(node)
	if len(confMap) < 1 {
		msg := fmt.Sprintf("%s configure no exists", node)
		return DBConfig{}, &err.Comerr{msg}
	}
	confMap["node"] = node
	return mapTODBConfig(confMap), nil
}

/**
查询数据库配置的[node] -> node
*/
func GetDBConfigNode() []string {
	f, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nodeSlice := make([]string, 0, 8)

	//读文件
	r := bufio.NewReader(f)
	for {
		//行读取
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		s := strings.TrimSpace(string(b))

		//如果是node。注意：注释和值的部分如果包含 [ ] 没有处理，
		n1 := strings.Index(s, "[")
		n2 := strings.Index(s, "]")
		if n1 == 0 && n2 > 1 {
			s = strings.Replace(s, "[", "", 1)
			s = strings.Replace(s, "]", "", 1)
			nodeSlice = append(nodeSlice, s)
		}
	}

	return nodeSlice
}

func findNodeMap(node string) map[string]string {
	f, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//读文件
	r := bufio.NewReader(f)

	conf := make(map[string]string)

	//要找的node 是否存在
	var exists bool = false
	for {
		//行读取
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		s := strings.TrimSpace(string(b))

		//如果是node。注意：注释和值的部分如果包含 [ ] 没有处理，
		n1 := strings.Index(s, "[")
		n2 := strings.Index(s, "]")
		if n1 == 0 && n2 > 1 {
			exists = false
		}
		//是要找的 node
		if strings.EqualFold(s, fmt.Sprintf("[%s]", node)) {
			exists = true
		} else {
			if exists && len(s) > 0 {
				pos := strings.Index(s, "=")
				key := s[0:pos]
				value := s[pos+1:]
				conf[key] = value
			}
		}
	}
	return conf
}

/**
数据库配置 tag 对应的配置文件的key
*/
type DBConfig struct {
	Node         string "node"
	Ip           string "ip"
	Port         string "port"
	Databasename string "databasename"
	Username     string "username"
	Passwd       string "passwd"
}

func mapTODBConfig(m map[string]string) DBConfig {
	dbConfig := DBConfig{}
	dbConfigValue := reflect.ValueOf(&dbConfig).Elem()
	dbConfigType := reflect.TypeOf(dbConfig)
	num := dbConfigType.NumField()
	for i := 0; i < num; i++ {
		field := dbConfigType.Field(i)
		if v, e := m[string(field.Tag)]; e {
			dbConfigValue.Field(i).SetString(v)
		}
	}
	return dbConfig
}

/**
转成map ， key 为 tag
*/
func dbConfigToMap(dbConfig DBConfig) (map[string]string, error) {
	var result = make(map[string]string)
	dbConfigValue := reflect.ValueOf(&dbConfig).Elem()
	dbConfigType := reflect.TypeOf(dbConfig)
	num := dbConfigValue.NumField()
	for i := 0; i < num; i++ {
		field := dbConfigType.Field(i)
		result[string(field.Tag)] = dbConfigValue.Field(i).String()
	}
	return result, nil
}

/**
追加数据源配置
*/
func AppendDBConfig(dbConfig DBConfig) error {
	f, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//写入的字符串切片
	inputStr := make([]string, 0, 6)

	//dbConfig 转成map ， key:tag v:value
	configMap, err := dbConfigToMap(dbConfig)
	if err != nil {
		return err
	}
	var nodeStr string
	for k, v := range configMap {
		if strings.EqualFold(k, "node") {
			nodeStr = "[" + v + "]"
		} else {
			inputStr = append(inputStr, k+"="+v)
		}
	}
	//写入文件
	if len(inputStr) > 1 {
		w := bufio.NewWriter(f)
		fmt.Fprintln(w, nodeStr)
		for _, v := range inputStr {
			fmt.Fprintln(w, v)
		}
		w.Flush()
	}
	return nil
}
