1. 所有相关的正则表达式包在regex 标注库里面

        import (
            "regex"
        )

2. 方法

   > Match(pattern string, b []byte) (matched bool, err error) 检查字符串是否满足正则表达式

    参数说明：

   + pattern: 正则表达式

   + b: 字符串转换的字节

   返回值：

    + matched: 满足返回true,不满足返回false

    + error ： 错误，有错误返回，没有错误为nil,一般都是正则表达式错误

            matched, err := regexp.Match(`foo.*`, []byte(`seafood`))

            // 返回true, nil
            matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
            fmt.Println(matched, err)
            // 返回false, nil 

            matched, err = regexp.Match(`a(b`, []byte(`seafood`))
	        fmt.Println(matched, err)

            false,error parsing regexp: missing closing ): `a(b`

            如果错误，则返回默认值false

   > MatchString(pattern string, s string) (matched bool, err error)

    同上面一样，只是不用在转换s

   > Regex 结构体
   
        type Regexp struct {
            // contains filtered or unexported fields
        }

    + func (re *Regexp) FindAllString(s string, n int) []string : 查找满足表达式的所有字符串，返回字符串切片(slice)

        + s: 要查找的string

        + n: 查找的次数，n>0 表示要查找n此，小于0 表示要查找全部

                func findAllString(searchIn, pat string) {
                    re, _ := regexp.Compile(pat)
                    // n=1, 表示只查找一个就结束了  [2578.34]
                    // n=2, 表示查找两次 [2578.34 4567.23]
                    // n<0, 表示查找全部 [2578.34 4567.23 5632.18]
                    c := re.FindAllString(searchIn, -2)
                    fmt.Println(c)
                }

                func main() {
                    searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18 at 4567+23"
                    pat := "[0-9]+\\.[0-9]+" //正则
                    findAllString(searchIn, pat)
                }

    +  (re *Regexp) MatchString(s string) bool ： 是否有匹配正则的字符串

    + ReplaceAllString(src, repl string) string: 替换字符串

            func replaceAllString() {
                    // 编译
                    re, _ := regexp.Compile(`a(x*)b`)
                    // a(x*)b 将匹配 ab axxb
                    // replace 使用新的替换正则表达式匹配的字符串，即ab 或 axxb
                    // 使用T 替换ab 或者axxb,因此返回-T-T-
                    fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
                    // 正则捕获分组分为两种，一种数字分组（使用$1-$n 代替小括号里面的内容），一种命名分组，$name a(?P<1W>x*)b ?p 命名捕获
                    //$1 这种只有数字的是数字分组， $1 在此实例中为“” 或者xx , 匹配ab 时$1是“”，匹配axxb 时，$1 是xx
                    // 因此--xx-
                    fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
                    // $1 后面有字符时，不再是数字分组，由最大长度作为名称，因此为 ${1d} ,当在正则中没有找到时使用”" 代替，因此结果为---
                    fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1d"))
                    // $1==${1} ,而 $1w==${1w}
                    fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))
                    fmt.Println(re.String())
                }
    + Split(s string, n int) []string: 分割字符串成切片

    +  String() string: 返回使用的正则表达式

            re, _ := regexp.Compile(`a(x*)b`)
            // 返回使用的正则表达式
	        fmt.Println(re.String())
            // a(x*)b

   > Compile(pattern string) (*Regexp, error) 返回结构体的指针

         pat := "[0-9]+.[0-9]+" //正则
        re, _ := regexp.Compile(pat)
    

        
