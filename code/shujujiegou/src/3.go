// 3.go
package main

import (
	"fmt"
	//"strconv"
	"strings"
)

func reverseSlice(arr []string) {

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	//return arr
}

func main() {

	// a := 12345
	// s := strconv.FormatInt(int64(a), 10)
	// arr := strings.Split(s, "")
	// reverseSlice(arr)
	// num := strings.Join(arr, "")
	// fmt.Println(num)
	// rtn, _ := strconv.ParseInt(num, 10, 64)
	// fmt.Println(rtn)

	a := []int{1, 2, 3, 4}
	s := strings.Join(a, "")
	fmt.Println(s)
}
