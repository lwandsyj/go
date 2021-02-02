package main

import (
	"fmt"
)

// 给定一个数组和一个target,选出两个数的下标
// 使用切片，因为数组长度是在编辑期就固定的，
func sum(arr []int, target int) []int {
	// 数组长度，如果数组长度为0 则返回空切片
	l := len(arr)
	// 初始化返回切片
	rtn := []int{}
	// 判断参数数组长度是否为0
	if l == 0 {
		return rtn
	}
	// 定义map
	m := map[int]int{}

	for index, value := range arr {
		// tmp 是目标数-当前值获得另一个参数
		tmp := target - value
		// 查看map 中是否存在另一个相加的参数，时间复杂度o(1)
		d, exist := m[tmp]
		// true 表示存在
		if exist {
			// 在数组末尾添加元素
			rtn = append(rtn, d)
			rtn = append(rtn, index)
		}
		m[value] = index
	}

	return rtn

}

func main() {
	//nums = [2,7,11,15], target = 9
	arr := []int{2, 7, 11, 15}
	target := 9
	rtn := sum(arr, target)
	fmt.Println(rtn)
}
