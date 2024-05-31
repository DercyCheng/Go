package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

// bufio是在file的基础上封装了一层API，支持更多的功能
// bufio通过缓存提高效率
func TestBufio(t *testing.T) {
	/*
		bufio:高效io读写
			buffer缓存
			io：input/output

		将io包下的Reader，Write对象进行包装，带缓存的包装，提高读写的效率

			ReadBytes()
			ReadString()
			ReadLine()

	*/

	fileName := "E:/code/Go/src/SC/garm/file/files/1/1.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//创建Reader对象
	//b1 := bufio.NewReader(file)
	//1.Read()，高效读取
	//p := make([]byte,1024)
	//n1,err := b1.Read(p)
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	//2.ReadLine()
	//data,flag,err := b1.ReadLine()
	//fmt.Println(flag)
	//fmt.Println(err)
	//fmt.Println(data)
	//fmt.Println(string(data))

	//3.ReadString()
	//s1, err := b1.ReadString('\n') //读到分割符
	//fmt.Println(err)
	//fmt.Println(s1)
	//s1, err = b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)
	//s1, err = b1.ReadString('\n')
	//fmt.Println(err)
	//fmt.Println(s1)

	//简化
	//for {
	//	s1, err := b1.ReadString('\n')
	//	if err == io.EOF {
	//		fmt.Println("done")
	//		break
	//	}
	//	fmt.Println(s1)
	//}

	//4.ReadBytes()
	//ReadBytes()返回一个包含已读取的数据和delim字节的切片
	//data, err := b1.ReadBytes('\n')
	//fmt.Println(err)
	//fmt.Println(string(data))

	//Scanner bufio读取到delim字段return,无法得到想要的结果
	//s2 := ""
	//fmt.Scanln(&s2)
	//fmt.Println(s2)

	//os.Stdin标准输入
	b2 := bufio.NewReader(os.Stdin)
	s2, _ := b2.ReadString('\n')
	fmt.Println(s2)
}
func TestWirte(t *testing.T) {
	filename := `E:/code/Go/src/SC/garm/file/files/1/1.txt`
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	w := bufio.NewWriter(file)
	file.Seek(0, io.SeekEnd) //尾起

	for i := 0; i < 10; i++ {
		w.WriteString(fmt.Sprintf("%d", i))
	}
	w.Flush()
}
