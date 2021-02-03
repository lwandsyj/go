// 数组
package main

import (
	"fmt"
)

func main() {
	// 定义时长度就固定了，长度不能修改，因此append 不能用于数组
	a := [4]bool{true, false, false}
	// 使用索引下标访问值
	// 长度固定不能修改，即不可以添加数组，但是可以修改元素值
	a[0] = false
	fmt.Println(a)
}
