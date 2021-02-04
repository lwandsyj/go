// concurrery.go
package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	//并行运行函数使用 go 关键字
	// main 和count 同时执行
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World!")
	time.Sleep(time.Millisecond * 5)
}
