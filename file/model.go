/**
model 生成
*/
package file

import (
	"bufio"
	"fmt"
	"os"
)

func GenerateBaseModel(filepath, packageName string) {
	f, _ := os.OpenFile(filepath+pathSeparator+"BaseModel.java", os.O_CREATE, os.ModePerm)
	defer f.Close()
	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "package "+packageName+".model.base;")
	fmt.Fprintln(w, "import java.io.Serializable;")
	fmt.Fprintln(w, "import java.util.Date;")

	w.Flush()
}

func GenerateMode(packageName, modelName string) {

}
