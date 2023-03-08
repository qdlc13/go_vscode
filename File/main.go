package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("E:/git_master/go_vscode/File/a.txt")
	if err != nil {
		fmt.Println("open err")
	}

	fmt.Printf("a.txt=%v", file)

	err2 := file.Close()
	if err2 != nil {
		fmt.Println("close err")
	}

}
