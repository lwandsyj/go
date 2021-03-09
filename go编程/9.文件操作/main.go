package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 读取文件内容，返回[]byte 切片
	content, _ := ioutil.ReadFile("./test.txt")
	// 转成string类型
	fmt.Println(string(content)) //hello file go 中国
	file := "./test2.txt"
	//type FileMode uint32
	// 写入文件，文件路径，内容为[]byte 切片，filemod 类型实际为为unit32
	// 777 代表全部权限，4代表读，2代表写，1 代表执行
	error := ioutil.WriteFile(file, content, 0777)
	if error != nil {
		fmt.Println(error)
	}
	// 返回当前文件所在的目录
	dir, error := os.Getwd()
	fmt.Println(dir)
}
