package file

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const dirname1 = "E:/code/Go/Src/SC"

func TestRange(t *testing.T) {
	listFiles(dirname1, 0)
}
func listFiles(dirname string, level int) {
	//level用来记录当前递归的层次，生成带有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	HandleErr(err)
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			//递归调用方法
			listFiles(filename, level+1)
		}
	}
}
