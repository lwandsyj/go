package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 数字类型转字符串，字符串转数字类型
func convertNumToStr() {
	a := 10
	// 数字转字符串
	// 定义： func Itoa(i int) string
	b := strconv.Itoa(a)
	fmt.Println(reflect.TypeOf(b)) // string

	// 字符串转int
	// 定义： func Atoi(s string) (int,error)
	// 因为有可能转换失败，因此string 转基本类型都返回两个值，一个是转换的类型值，一个是error
	// 转换成功，返回正常数字，error 为nil
	// 转换失败，数字返回默认值为0，error 包含错误信息，不为nil
	c, error := strconv.Atoi(b)
	if error != nil {
		fmt.Println("error:", error)
		return
	}
	fmt.Println(c)

	// FormatInt 64位数字转换字符串
	// 这里使用的int64, 在go 语言中不存在隐式转换，int 是根据操作系统来定的，即使是64位的，和int64 也是类型不同的,因此需要隐式转换
	//./convert.go:31:6: cannot use a (type int) as type int64 in assignment
	// var g int64 = a
	//func FormatInt(i int64, base int) string
	//func FormatUint(i uint64, base int) string
	// base 是进制数，2~64
	d := strconv.FormatInt(int64(a), 10)
	fmt.Println(d)

	// ParseInt
	//func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// @param s : 要转成数字的字符串
	// @param base: 进制数
	// @param bitSize: int位数 bitSize 必须设置无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64。 如果 bitSize 小于 0 或大于 64 将返回错误。
	// 转换是注意int 类型的最大范围
	h, err := strconv.ParseInt(d, 10, 64) // 那么h 就是int64
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(h)
}

func main() {

}
