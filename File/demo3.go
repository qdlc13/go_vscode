//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取文件的内容并显示在终端(带缓冲区的方式-4096字节)，
// 适合读取比较大的文件，使用os.Open,file.Close,bufio.NewReader(),reader.ReadString函数和方法
func main() {
	file, err := os.Open("E:/git_master/go_vscode/File/a.txt")

	if err != nil {
		fmt.Println("open file err", err)
	}
	defer file.Close()

	//创建一个流
	reader := bufio.NewReader(file)

	for {
		//读到一个换行就结束
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读到文件结尾
			break
		}
		//没有读到结尾就正常输出内容
		fmt.Println(str)
	}
	//结束
	fmt.Println("file is over")

}
