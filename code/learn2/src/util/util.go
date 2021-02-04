// go 语言中每个模块必须先声明包
package uitl

// 要想被外面的包引用，则首字母必须大写，
func Add(a, b int) int {
	return a + b
}

func del(a, b int) int {
	return a - b
}
