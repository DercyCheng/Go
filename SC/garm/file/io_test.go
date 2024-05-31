package file

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	filename := `files/1/1.txt`
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file)
	defer file.Close()

	bs := make([]byte, 4, 4)

	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到了文件的末尾，结束读取操作。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func TestWrite(t *testing.T) {
	filename := `files/1/1.txt`
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 写入数据
	bs := []byte{65, 66, 67, 68, 69, 70} // A,B,C,D,E,F
	// bs := []byte{97, 98, 99, 100} //a,b,c,d
	n, err := file.Write(bs)
	// n, err := file.Write(bs[:2])
	fmt.Println(n)
	HandleErr(err)
	file.WriteString("\n")

	// 直接写出字符串
	n, err = file.WriteString("HelloWorld")
	fmt.Println(n)
	HandleErr(err)

	file.WriteString("\n")
	n, err = file.Write([]byte("today"))
	fmt.Println(n)
	HandleErr(err)

}
