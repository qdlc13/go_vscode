package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//读取文件的内容并显示在终端(使用ioutil一次将整个文件读入到内存中)，
	//这种方式适用于文件不大的情况。相关方法和函数(ioutil.ReadFile)

	//备注：在下面的程序中不需要进行 Open\Close操作，
	//因为文件的打开和关闭操作被封装在ReadFile函数内部了
	content, err := ioutil.ReadFile("E:/git_master/go_vscode/File/a.txt")
	if err != nil {

		fmt.Println("err is =", err)
	}
	//如果读取成功，将内容显示在终端即可：
	// fmt.Printf("%v", content)
	fmt.Printf("%v", string(content))

}
