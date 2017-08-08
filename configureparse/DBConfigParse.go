/**
数据库配置文件dbconfig解析
*/
package configureparse

import (
	"bufio"
	"fmt"
	"golang/err"
	"io"
	"os"
	"reflect"
	"strings"
)

func GetDBConfig(node string) (DBConfig, error) {
	//配置文件地址
	var file_path = "./dbconfig"

	//配置文件转成的map
	confMap := findNodeMap(file_path, node)
	if len(confMap) < 1 {
		msg := fmt.Sprintf("%s configure no exists", node)
		return DBConfig{}, &err.Comerr{msg}
	}
	return mapTODBConfig(confMap), nil
}

func findNodeMap(file_path, node string) map[string]string {
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
	fmt.Println(dbConfigValue.String())
	num := dbConfigType.NumField()
	for i := 0; i < num; i++ {
		field := dbConfigType.Field(i)
		if v, e := m[string(field.Tag)]; e {
			dbConfigValue.Field(i).SetString(v)
		}
	}
	return dbConfig
}
