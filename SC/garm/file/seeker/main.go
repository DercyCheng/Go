package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//断点续传
	//了解
	/*
			seek(offset,whence),设置指针光标的位置
			第一个参数：偏移量
			第二个参数：如何设置
				0：seekStart表示相对于文件开始，
				1：seekCurrent表示相对于当前偏移量，
				2：seek end表示相对于结束。


			const (
			SeekStart   = 0 // seek relative to the origin of the file
			SeekCurrent = 1 // seek relative to the current offset
			SeekEnd     = 2 // seek relative to the end
		)

			随机读取文件：
				可以设置指针光标的位置
	*/

	file, _ := os.OpenFile("E:/code/Go/src/SC/garm/file/files/1/1.txt", os.O_RDWR, 0)
	defer file.Close()
	bs := []byte{0}

	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4, io.SeekStart) //距离头4个字节
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(2, 0) //也是SeekStart
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(3, io.SeekCurrent)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(0, io.SeekEnd) //距离末尾0个字节
	file.WriteString("ABC")  //末尾写
}
