// 字符串操作
package main

import (
	"fmt"
	"strings"
)

//1 func 定义
func defineStr() {
	var a string // 默认值 空
	b := "hello" // 必须使用双引号包裹
	// 使用``可以换行
	c := `
		hello 
		world
	`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// 2.字符串字符
func charStr() {
	a := "hello中国world"
	// 中文转成字节一个中文占三个字节
	b := []byte(a) //[104 101 108 108 111 228 184 173 229 155 189 119 111 114 108 100]
	// rune 可以包含中文，日文，韩文等Unicode 字符
	c := []rune(a) //[104 101 108 108 111 20013 22269 119 111 114 108 100]
	fmt.Println(b)
	fmt.Println(c)
	d := string(c) //helloworld, 字符数组转成字符串
	fmt.Println(d)

	fmt.Println(string(a[5])) // ä
	fmt.Println(string(c[5])) // 中

}

//3. len 和索引
func lenStr() {
	a := "hello中国world"
	b := len(a) // 16，一个中文占三个字节
	fmt.Println(b)

	// 因此要获取字符串字符的个数，要先把字符串转成rune(使用rune 而不是用byte是因为rune 能包含中文)
	c := []rune(a)
	// 求取字符串的字符个数
	fmt.Println(len(c)) //12

	fmt.Println(a[0])         //104
	fmt.Println(string(a[0])) //英文因为占一个字节，因此能正常返回 h
	fmt.Println(string(a[5])) // ä,不能正常返回
	// 索引获取的是字节，因此要用string类型转换
	fmt.Println(c[0])         //104
	fmt.Println(string(c[0])) // h
	fmt.Println(string(c[5])) // 中
}

//4. 字符串循环
func forStr() {
	a := "hello中国world"
	for index, value := range a {
		fmt.Printf("index=%d,value=%s\n", index, string(value))
	}
}

// 字符串连接
func concatStr() {
	a := "hello"
	b := "world"
	fmt.Println(a + b)

	s := fmt.Sprintf("%[2]s %[1]s", a, b)
	fmt.Println(s)
}

// strings 类库操作
func stringsStr() {
	//1. compare(a,b string) int 两个字符串大小比较 a==b 返回0，a<b 返回-1 a>b 返回1
	a := strings.Compare("a", "b") // a<b 根据转换的assic 码值计算
	fmt.Println(a)
	//2. Contains 检查一个字符串中是否包含子字符串，包含返回true, 不包含返回false
	//定义：Contains(s,sub string) bool
	b := strings.Contains("hello", "wo")
	fmt.Println(b) // false
	//3.是否已pref 开头，类似其他语言的startsWith
	// func hasPrefix(s,pref string) bool
	c := strings.HasPrefix("hello", "he")
	fmt.Println(c)
	// 4. 是否已suffix 结尾，类似其他语言的endsWith
	// func hasSuffix(s,suffix string) bool
	d := strings.HasSuffix("hello", "lo")
	fmt.Println(d)
	//5. index 返回子串首次出现在字符串中的下标，没有返回-1
	// LastIndex: 子串最后出现的位置，没有返回-1
	e := strings.Index("hello", "e")
	fmt.Println("5-index:", e)
	f := strings.Index("hello", "g")
	fmt.Println(f)
	// 6.join
	g := []string{"h", "b", "c"}
	h := strings.Join(g, "-")
	fmt.Println("6-join:", h)
	//7. split
	i := strings.Split(h, "-")
	fmt.Println(i)
	//8. replace
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//9. Trim
	fmt.Println(strings.Trim("¡¡¡Helloi, Gophers!!!", "!¡"))
}
func main() {
	stringsStr()
}
