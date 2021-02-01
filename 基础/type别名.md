1. 在go 语言中type 有两种写法

+ 作为别名

        type bieming = Type

        type myInt =int 

   myInt 还是int 类型，只是使用myInt 别名代替，通常别名用于复杂类型的简写，比如

        type Map =map[string]string

        a:=Map{"name":"张三"}

        // 对比原有的

        a:=map[string]string{"name":"张三"}

        // 查看类型

        fmt.Println(reflect.TypeOf(a))

        返回map[string]string 类型

+ 生成新的类型

        type newType Type 

        比如： 结构

        type student struct{

        }

        比如： 

        type myInt int // 本质上还是int

        func main() {
            type myInt int

            var a myInt = 3

            fmt.Println(reflect.TypeOf(a))
        }

        输出：main.myInt

   这种用法通常用来扩展方法，包括go 语言的基础类型，

        type myInt int

        // 两数比较
        func (a myInt) compare(b myInt) bool {
            return a > b
        }

        func main() {

            var a myInt = 3

            // 调用扩展方法
            b := a.compare(4)

            fmt.Println(b)
        }
