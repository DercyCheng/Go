package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	filename = "E:/code/Go/src/SC/garm/file/files/1/1.txt"
	dirname  = "E:/code/Go"
)

func TestReadFile(t *testing.T) {
	// 也许大部分基本的文件读取任务是将文件内容读取到内存中
	data, _ := ioutil.ReadFile(filename)
	fmt.Println(string(data))
}
func TestWriteFile(t *testing.T) {
	//会抹去原来的数据
	s := "ioutil.WriteFile:hello world"
	err := ioutil.WriteFile(filename, []byte(s), os.ModePerm)
	HandleErr(err)
}
func TestReadALl(t *testing.T) {
	s := "ioutil.ReadAll:hello world"
	b := strings.NewReader(s)
	data, err := ioutil.ReadAll(b)
	fmt.Println(string(data))
	HandleErr(err)
}

// ReadDir()，读取一个目录下的子内容：子文件和子目录，但是只能读取一层
func TestReadDir(t *testing.T) {
	fileinfo, err := ioutil.ReadDir(dirname)
	HandleErr(err)
	fmt.Println(len(fileinfo))
	for i := 0; i < len(fileinfo); i++ {
		//fmt.Printf("%T\n",fileInfos[i])
		fmt.Printf("第 %d 个：名称：%s，是否是目录：%t\n", i, fileinfo[i].Name(), fileinfo[i].IsDir())
	}
}

// 创建临时目录
func TestTemp(t *testing.T) {
	dir, err := ioutil.TempDir("E:/code/Go/src/SC/garm/file/files", "test")
	HandleErr(err)
	defer os.Remove(dir) // 用完删除
	fmt.Printf("%s\n", dir)

	// 创建临时文件
	f, err := ioutil.TempFile(dir, "Test")
	HandleErr(err)
	defer os.Remove(f.Name()) // 用完删除
	fmt.Printf("%s\n", f.Name())
}
