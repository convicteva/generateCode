package util

import "encoding/json"

func ToSlice(str string) []string {
	var c = make([]string, 0, 10)
	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		return c
	}
	return c
}
