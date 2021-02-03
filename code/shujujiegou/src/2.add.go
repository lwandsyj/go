// 两个链表相加和

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// node 结构
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 1. 反转链表中的数据
	rl1 := reverseListNode(l1)
	rl2 := reverseListNode(l2)
	// 2. 两数相加
	result := rl1 + rl2
	// 3. 返回输出nodelist
	rtn := reverseToListNode(result)
	return rtn
}

func reverseSlice(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// // 把整形切片转换成数字
// func converArrayToInt(arr []int) int {
// 	s := strings.Join(arr)
// 	rtn, err := strconv.ParseInt(s, 32, 10)
// 	if err {
// 		fmt.Println("strconv.ParseInt err:", err)
// 		return 0
// 	}
// 	return rtn
// }
func reverseListNode(l *ListNode) int {
	current := l
	rtn := []string{}
	for current.Next != nil {
		rtn = append(rtn, strconv.FormatInt(int64(current.Val), 10))
		current = current.Next
	}
	// 反转
	reverseSlice(rtn)
	result := strings.Join(rtn, "")

	i, _ := strconv.ParseInt(result, 10, 64)
	return int(i)
}
func reverseToListNode(num int) *ListNode {
	s := strconv.FormatInt(int64(num), 10) // 转成字符串
	arr := strings.Split(s, "")            // 转成切片
	reverseSlice(arr)                      // 反转
	return convertSliceToListNode(arr)
}

func convertSliceToListNode(arr []string) *ListNode {
	var tmp *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		m, _ := strconv.ParseInt(arr[i], 10, 64)
		n := int(m)
		if i == len(arr)-1 {
			tmp = &ListNode{Val: n, Next: nil}
		} else {
			tmp = &ListNode{Val: n, Next: tmp}
		}
	}
	return tmp
}

func main() {
	a := []string{"2", "4", "3"}
	b := []string{"5", "6", "4"}
	c := convertSliceToListNode(a)
	fmt.Println(c)
	d := convertSliceToListNode(b)
	fmt.Println(d)
	fmt.Println(addTwoNumbers(c, d))
}
