package file

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)

/*
FileInfo：文件信息

	interface
		Name()，文件名
		Size()，文件大小，字节为单位
		IsDir()，是否是目录
		ModTime()，修改时间
		Mode()，权限
*/
func TestFile(t *testing.T) {
	fileInfo, err := os.Stat(`C:\Users\cheng\Desktop\1.txt`)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Mode())
}
func TestOperation(t *testing.T) {
	/*
		文件操作：
		1.路径：
			相对路径：relative
				ab.txt
				相对于当前工程
			绝对路径：absolute
				/Users/ruby/Documents/pro/a/aa.txt

			.当前目录
			..上一层
		2.创建文件夹，如果文件夹存在，创建失败
			os.MkDir()，创建一层
			os.MkDirAll()，可以创建多层

		3.创建文件，Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
			os.Create()，创建文件

		4.打开文件：让当前的程序，和指定的文件之间建立一个连接
			os.Open(filename)
			os.OpenFile(filename,mode,perm)

		5.关闭文件：程序和文件之间的链接断开。
			file.Close()

		5.删除文件或目录：慎用，慎用，再慎用
			os.Remove()，删除文件和空目录
			os.RemoveAll()，删除所有
	*/
	filename1 := `E:\code\Go\src\SC\garm\file\files\1\1.txt`
	filename2 := `files/2/2.txt`
	fmt.Println(filepath.IsAbs(filename1))
	fmt.Println(filepath.Abs(filename2))
	//创建目录
	//创建单层文件夹
	err_single := os.Mkdir("files", os.ModePerm)
	if err_single != nil {
		fmt.Println("err:", err_single)
		return
	}
	//创建多层文件夹
	err_muti := os.MkdirAll("files/1", os.ModePerm)
	if err_muti != nil {
		fmt.Println("err:", err_muti)
		return
	}
	//创建文件:Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
	file1, err := os.Create("files/1/1.txt")
	if err != nil {
		fmt.Println("err：", err)
		return
	}
	fmt.Println(file1)

	os.MkdirAll("files/2", os.ModePerm)
	file2, err := os.Create(filename2) //创建相对路径的文件，是以当前工程为参照的
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	fmt.Println(file2)

	file3, err := os.Open(filename1) //只读
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file3)

	file4, err := os.OpenFile(filename1, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file4)
	//file4.Close()
}
func TestDelete(t *testing.T) {
	os.Remove(`files\1\1.txt`)
	os.RemoveAll("files")
}
