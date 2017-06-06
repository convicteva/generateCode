package file

import "os"

var pathSeparator string

func init() {
	pathSeparator = string(os.PathSeparator)
}
