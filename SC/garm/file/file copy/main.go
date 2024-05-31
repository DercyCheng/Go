package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	srcFile := "E:/code/Go/src/SC/garm/file/files/1/1.txt"
	destFile := "E:/code/Go/src/SC/garm/file/files/2/2.txt"
	total, err := rwCopy(srcFile, destFile)

	//performance test

	//total, err := RWioutil(srcFile, destFile)
	fmt.Println(total, err)
}

// RWioutil ioutil包中的 `ioutil.WriteFile()`和 `ioutil.ReadFile()`
// 但由于使用一次性读取文件
// 再一次性写入文件的方式,所以该方法不适用于大文件,容易内存溢出
func RWioutil(srcFile, destFile string) (int, error) {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		fmt.Println("操作失败：", destFile)
		fmt.Println(err)
		return 0, err
	}

	return len(input), nil
}

// 直接使用io包下的Copy()方法
func ioCopy(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	fmt.Println("done")
	return io.Copy(file1, file2)

}

// io包下的Read()和Write()方法实现
func rwCopy(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	//读写
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕。。")
			break
		} else if err != nil {
			fmt.Println("报错了。。")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}
