// go 语言中内置函数
package main

import (
	"fmt"
)

// append 内置函数，在slice 末尾追加元素
// 不能用于数组，因为数组长度是在编辑期就固定的，不能修改数组的长度
// append 的定义
/*
	func append(slice []Type,elems...Type) []Type
*/
func learnAppend() {
	// 切片
	a := []int{}
	// 必须是切片,否则会报类型不一致
	//cannot use b (type [2]int) as type []int in append
	b := []int{8, 9}

	a = append(a, 4)
	// append 采用的是变长参数，可以传入多个
	a = append(a, 5, 6, 7)
	// ... 把数组或切片变成上面的形式传入参数
	a = append(a, b...)
	fmt.Println(a)
}

// len 内置函数，返回字符串，数组，slice,map ，channel 长度
// len定义
// func len(v Type) int
func learnLen() {
	// 字符串返回字符所占字节数
	// 一个英文占一个字节，一个中文占三个字节
	a := "hello world"
	fmt.Println(len(a)) //11

	b := "你好，中国!h"
	fmt.Println(len(b))         // 返回 17
	fmt.Println(len([]byte(b))) // 字符串转成字节，17

	fmt.Println(len([]rune(b))) // unicode 字符，7
	// 切片返回slice 格式
	c := []string{"a", "中"}
	fmt.Println(len(c))
	// map 返回key 的个数
	d := map[string]string{"name": "1", "age": "2"}
	fmt.Println(len(d))
}

// 用来分配内存，返回Type本身(只能应用于slice, map, channel)
// 初始化
func learnMake() {
	a := make([]int, 4)
	fmt.Println(a)
	//
	b := make([]int, 2, 4)
	fmt.Println(b)
	// 长度
	fmt.Println(len(b))
	// 容量
	fmt.Println(cap(b))
	// map 和channel 可以省略size
	c := make(map[string]string)
	fmt.Println(c)
}

type good struct {
	name string
	age  int
}

// 初始化，但是返回指针
func learnNew() {
	a := new(int)
	b := new(good)
	fmt.Println(a)      //0xc00010c008
	fmt.Println(*a)     // a 的值
	fmt.Println(b)      //&{ 0}
	b.name = "zhangsan" // 点可以用作解指针
	fmt.Println(b.name)
	fmt.Println(b.age)
}

// delete 从map 中删除key 对象的value
// 定义： func delete(m map[kType]vType,key Type)
func learnDelete() {
	c := make(map[string]string)
	c = map[string]string{"name": "张三", "age": "14", "pwd": "pwd"}
	fmt.Println(c) //map[age:14 name:张三 pwd:pwd]
	// 删除pwd key
	delete(c, "pwd") //map[age:14 name:张三]
	fmt.Println(c)
}
func main() {
	learnDelete()
}
