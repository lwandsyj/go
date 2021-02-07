// concurrery.go
package main

import (
	//"errors"
	"fmt"
	"strings"
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
	//通道为两个goroutine提供了一种相互通信的方式
	fmt.Println(strings.ContainsAny("team", "i")) // false
	// true ,任何一个字符在里面就是true ，ui 拆成u 和i,这两个字符任何一个就可以
	fmt.Println(strings.ContainsAny("fail", "ui")) // 包含i true
	//包含u
	fmt.Println(strings.ContainsAny("ure", "ui"))
	// u 和 i 都包含
	fmt.Println(strings.ContainsAny("failure", "ui"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World!")
	time.Sleep(time.Millisecond * 5)
}
